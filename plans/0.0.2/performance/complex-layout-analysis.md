# Complex layout performance analysis

Captured on 2026-10-06, against `feature/perf-observability` at `9e7767a`. The starting worktree was clean. This adds a profiling fixture and analysis; the library and engine implementations are unchanged.

## Findings

The largest measured gain comes from reducing how much document enters layout. Keeping 48 of 480 service rows reduced redraw time from 306.66 ms to 34.74 ms, an 8.83x speedup and 88.7% less elapsed time. Allocation fell from 86.25 MB to 9.27 MB per redraw, a 89.2% reduction. Operation count fell from 12,360 to 1,344. This is evidence for extending the existing row-window strategy to complex collections.

Cached HTML does not make layout cheap. Cached and resized pages retain one parse and one cascade, yet create a new layout and display list for every redraw. A one-label change through `SetData` also reparses the source. Its dirty diff ultimately finds one changed operation and three dirty operations, after paying for the whole layout. Small repaint and small layout are different costs.

The dependency is the next place to investigate. In run 2's cached CPU profile, `measureFlexCrossMax` accounts for about 35% cumulative CPU, `appendSheetRuleHits` about 22%, and `flexMinMainSize` about 16%. These call trees overlap. In the sampled allocation profile, approximately 63% of allocation is attributed directly to the local style override in `buildFlexRowItem`.

## Workload and results

The fixture is an operations dashboard with navigation, six cards, regional bars, an incident timeline, a filter form, and 480 fixed-height service rows. Rows contain nested flex containers, wrapped descriptive text, three metrics, status, and six miniature bars. It produces 9,672 hit boxes and 12,360 operations. The logical canvas is 1440 by 37,032 pixels. The grid is 36,480 pixels high and begins at approximately Y=540.81.

| Scenario | Redraw ms | Allocated MB | Allocations | Boxes | Operations |
|---|---:|---:|---:|---:|---:|
| Fresh page, process caches warm | 313.89 | 91.06 | 275,646 | 9,672 | 12,360 |
| Unchanged cached redraw | 306.66 | 86.25 | 221,047 | 9,672 | 12,360 |
| One revision label changed | 330.63 | 91.05 | 275,489 | 9,672 | 12,360 |
| Alternating width | 312.53 | 86.25 | 221,049 | 9,672 | 12,360 |
| 48 rows and retained-height spacer | 34.74 | 9.27 | 24,696 | 1,032 | 1,344 |

Times are medians of three process medians, each with seven timed samples. Fresh-page sample zero is excluded from the steady fresh-page row. Its three process-first times were 308.62 ms, 340.23 ms, 339.18 ms; fresh-page construction is outside the timed interval. MB means 1,000,000 bytes. Allocation values come from a representative sample in the median-time process. Raw per-process medians and stage samples are in [summary.json](../../../temp/complex-dump/summary.json).

Windowed redraw process medians ranged from 31.47 to 46.42 ms. Full cached redraw medians ranged from 306.31 to 339.68 ms. Run on the same host and use repeated, interleaved A/B measurements before treating a proposed production change as a release claim. The present windowing comparison uses the same executable and fixture in separate sequential processes; it establishes the scale of the gain, rather than a confidence interval.

All samples have `PaintTime=0`, and every redraw retained a replayable display list. The command opened no window, served no web page, and produced no PNG or PDF. Tests and the original dump command were separate validation steps.

## What the profiles mean

`Page.Stats.LayoutTime` currently times `styledDocument`, which prepares CSS or calls `css.Relayout`. Actual box layout is inside `DisplayListTime`. For this fixture, display-list creation accounts for nearly all redraw latency. Calling the bottleneck simply "CSS relayout" would misidentify it.

The cached path skips HTML parsing and sheet collection, but `layout.DisplayListOptions` still resolves styles and places the tree. Engine `appendSheetRuleHits` loops through sheet rules for each candidate node. Cached parse counters therefore do not imply that selector matching vanished.

Engine `measureFlexCrossMax`, in `internal/layout/flex.go:1351`, sets `noEmit` and recursively calls `build` to find the tallest flex item. Final placement builds items again. Nested flex makes these measurement builds expensive. Reusing valid measurement results or a size-only traversal is a plausible engine optimization, subject to stretch, wrapping, percentage sizes, baseline, and container-query behavior.

Engine `buildFlexRowItem`, in `internal/layout/flex.go:1454`, copies a resolved style, adjusts dimensions, then passes its address to `buildWithStyle`. The allocation profile attributes the large direct allocation share to line 1458. Boxes may retain that pointer after the override stack is popped, as the comment on `buildWithStyle` states. Reusing one mutable override object would violate that lifetime. Investigate compact dimension overrides or storage owned by one layout instead.

The dependency is pinned at `v0.2.7-0.20261004151708-1a3918301a68`. These engine changes belong in gowkhtmltopdf, followed by a dependency update here. Source was read from that exact installed module. [Cumulative CPU](../../../temp/complex-dump/cached-cumulative.txt) and [allocation by source line](../../../temp/complex-dump/flex-allocation-lines.txt) preserve the evidence.

## Optimization estimates and order

Except for windowing, the numbers below are planning assumptions derived from profile shares, not measured implementation gains. They apply separately to this workload and must not be added together.

| Priority | Proposed solution and owner | Estimated benefit | Basis and acceptance gate |
|---|---|---|---|
| 1 | Extend app row windowing to heavy lists. Reuse the current scroll observer and spacer pattern. | Measured 88.7% less redraw time and 89.2% less allocation here. Plan around 80-90% less time for comparable fixed-row collections. | First 48 row boxes and ordered visible operation geometry/text match; retained height matches. Verify live scrolling, focus, filtering, selection, resize, and middle/end windows. |
| 2 | Reuse valid flex measurement results or introduce a size-only pass in the engine. | Estimate 9-18% less total CPU on the full fixture if 25-50% of the 35% measurement subtree is removed. About 1.10-1.22x CPU throughput. Latency and allocation gains remain to be measured. | Profile confirms recursive measurement builds. Match full ordered operations, boxes, wrapping, stretch and percentage behavior before claiming a gain. |
| 3 | Replace full escaping flex style copies with compact immutable overrides or layout-owned storage. | Estimate 31-50% less allocation on the full fixture if 50-80% of the 63% direct allocation site is removed. That models roughly 43-59 MB/redraw from an 86 MB baseline. CPU gain is unknown. | Preserve the lifetime of style pointers retained by boxes. Require allocation profiles, heap/RSS checks, output parity and repeat benchmarks. Overlaps priority 2. |
| 4 | Index stylesheet candidates by selector shape in the engine, with a conservative fallback. | Estimate 5-11% less total CPU if 25-50% of the 22% rule-selection subtree is removed. About 1.06-1.12x CPU throughput. | Preserve specificity, source order, state selectors, media/container queries, imports and pseudo elements. Use this fixture and a rule-heavy corpus. Overlaps measurement work. |
| 5 | Update display operations directly for bounded progress/bar animations and avoid redundant redraw scheduling. App/frame path. | A skipped redundant redraw removes its entire measured layout cost. For a supported geometry-only update, estimate avoiding over 95% of this redraw's CPU/allocation work, subject to direct-update measurement. | Existing stress/flappy examples already mutate operations through tick callbacks. Measure update cost and prove dirty-region, font and hit-box validity. Text or layout changes need stronger invalidation. No additional idle saving is established. |
| 6 | Keep unchanged `SetData` output on the parsed cache when safe. Page pipeline. | Only about 3-4 ms of CSS/parse work and roughly 4.8 MB of allocation are at stake here, around 1% time and 5% allocation. A changed source still needs a parse with the present architecture. | `SetData` currently invalidates eagerly. A source-equality cache policy would need theme, form, image, state and reload tests. It will not remove display-list rebuilding. |

Start with windowing because it removes work throughout the pipeline and already has a public callback path. Then reprofile the smaller document. Its remaining costs and optimization ranking can differ from the full 480-row case. Even the windowed redraw is around 35 ms here; this analysis does not establish 60 fps.

## Evidence and limits

[run.sh](complex-dump/run.sh) reproduces the capture. [layout.html](complex-dump/layout.html) is the fixture. [README.md](complex-dump/README.md) explains profiles, scenarios and timing boundaries. [manifest.json](../../../temp/complex-dump/manifest.json) and [build.txt](../../../temp/complex-dump/build.txt) record source, toolchain and binary identity. The host is an Intel i7-13700HX under WSL2, with Go 1.26.4 and the default runtime processor setting.

CPU profiles cover the seven samples and at least two more seconds of repeated work. They include scenario preparation and Go GC activity. CPU samples can exceed wall duration because GC runs concurrently. Heap profiles are sampled process profiles after GC, including package initialization, and are captured before JSON/compression writes. Allocation traffic is not retained heap or RSS. No leak conclusion follows from these short runs.

The dump contains boxes and projected operations, rather than pixels or a complete replay payload. All three runs prove identical initial viewport operation geometry/text, first-48-row geometry, and total height. This is structural equivalence only. Offscreen rows are intentionally absent from the windowed display. Variable row heights, accessibility, focus beyond the window, deep scroll positions and pixels need later validation.

The previous branch already shipped windowing and oversized replay handling in `9e7767a`, plus dirty-diff and box-index changes in `82ca687`. Those are baseline behavior. This report neither attributes their old gains to a new library change nor predicts GPU replay, frame percentiles, desktop CPU, or RSS from a headless layout profile.

See [checklist-complex-layout.md](checklist-complex-layout.md) for completed analysis phases and pending implementation gates.

## Validation

`make test` and `make build` both exited 0 after the final Go edits. All new Go files are formatted and below 2000 characters. Logs are in [make-test.txt](../../../temp/complex-dump/make-test.txt) and [make-build.txt](../../../temp/complex-dump/make-build.txt).

The existing headless dump script also completed. Its single snapshots measured stress-240 at 476 operations and 17.27 ms, large-1000 at 4,003 operations and 92.82 ms, normal at 14 operations and 0.39 ms, and flappy at 11 operations and 0.42 ms. These are context only; the new dashboard has roughly three times the operation count of large-1000 and uses nested flex. The dump is archived under `temp/complex-dump/existing-apps`.
