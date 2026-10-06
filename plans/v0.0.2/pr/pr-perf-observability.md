# perf(window): opt-in sampling, row windowing, and oversize replay for large pages

Filled copy of [`skills/PR/PR_TEMPLATE.md`](../../../skills/PR/PR_TEMPLATE.md), saved before opening the PR. Local only: nothing is committed, pushed, or published from this file.

- Base: `origin/master` (`7820e4f`)
- Head: `feature/perf-observability` (pushed, working tree clean)
- Commits: 11
- Diff: 154 files, +5349 / -330

---

## Summary

This branch makes the render loop measurable and stops large pages from falling over. Perf sampling is opt-in through `WithPerf`, `Config.Perf`, `WindowOptions.Perf`, and `ServeOptions{Perf: true}`; when it is off, shipped apps pay nothing and the DevTools rows read `-`. Oversized pages no longer crash: content past the Ebiten image limit replays directly or through a bounded viewport cache, and row windowing renders only the visible rows plus overscan. The branch also adds three perf example apps, baseline and soak scripts, and the performance docs.

## Motivation / context

- Large pages were display-list bound. The branch's own dumps show a 1000-row page spending 84.3 of 88.8 ms in display-list build, and the desktop normal (480x16494) and large (800x36579) apps panicked when the persistent buffer tried to hold the whole content.
- There was no way to name the slow stage without attaching a profiler to the whole process, and no committed baseline to compare releases against.
- Plans: `plans/v0.0.2/performance/`, `plans/v0.0.2/0.0.2/performance/complex-layout-analysis.md`.
- Issues: none. This repo has no tracker entries for this work.

---

## Pending changes matrix: what was there, what it is now

Everything below is `origin/master` (`7820e4f`) against this branch HEAD (`fa84ab1`).

| Area | What was there | What it is now |
|---|---|---|
| Perf sampling and `Stats` | No sampling. `host.Stats` had counters only: redraws, parses, cascades, layouts, repaints, ops, last redraw, last draw, reloads, last reload error. | Opt-in sampling. `Config.Perf`, `WithPerf(bool)`, `NewWithOptions`, `Page.SetPerf`/`Page.Perf`, `WindowOptions.Perf`, and `ServeOptions{Perf: true}`. `Stats` adds `LastTemplate`, `LayoutTime`, `DisplayListTime`, `PaintTime`, `DirtyOps`, `DirtyRegions`, `ChangedOps`, `AllocFrame`. Off by default, so untracked redraws leave the new fields at zero. |
| DevTools Frame tab | WINDOW, RENDERING, PIPELINE, RELOAD rows. | Adds Performance rows (frame avg, p95, p99, long frames), Pipeline detail rows (template, layout, display list, paint, dirty ops/regions, changed ops), and Memory rows (alloc/frame, Go heap, RSS, goroutines). Perf rows read `-` until the app opts in. |
| Frame and runtime sampling | None. No frame ring, no runtime read. | 120-frame ring with avg/p95/p99 and a long-frame count over 25 ms. A 500 ms sampler records heap, RSS (Linux `/proc/self/statm`), goroutines, and the `TotalAlloc` delta. Runs only when Perf is on. |
| Oversized pages | Partial repaint kept one persistent buffer sized to the content. Content past `ebiten.MaxImageSize` panicked in the atlas. | `oversized` detects content no image can hold. Oversized and ticking pages replay directly each frame; oversized static pages use a bounded viewport cache with overscan that survives scroll offsets. Content-sized buffers are dropped while the viewport path is active. |
| Scroll and row windowing | Scroll changes blitted or paid a full relayout. No windowing API. | `Page.SetWindowing`, `SetScrollWindow(func(offsetY, viewH int) bool)`, `SetScrollOffset`/`ScrollOffset`, `StepScrollWindow`, and the `host.ScrollObserver` / `host.ScrollWindowStepper` interfaces. The page slices app rows before template execute; an unchanged overscan returns false and the host skips the scroll relayout. |
| Hover | Hover always marked dirty and went through a full redraw. | `Handlers.Hover` and the paint-only `page.HoverHandler`. A handled hover repaints without a layout pass. |
| Dirty tracking and box lookups | Linear box scans per lookup, `contentRect` recomputed on every `TakeDirty`, and diffs that unioned thousands of rects. | `id -> box` index rebuilt per redraw, memoized content rect per canvas size, early full-frame exit past `maxDirtyOps` (8 changed ops) or a dirty area over one third of the frame, and a presized pairing map. Commit 82ca687 measured Invalidate+TakeDirty at about 4.6 us to 7 ns/op, and id lookup at about 15 us to 3.5 us. |
| Web serve and debug surface | `Serve` alone. It mounted page, frame.png, debug/state, click, type, and backspace, with no profiler. | `ServeWithOptions` and `ServeOptions`. `ServeOptions{Perf: true}` adds `GET /debug/pprof/*`; plain `Serve` exposes no profiler. `GET /debug/state` can carry a runtime snapshot, but production never sets `RuntimeSnapshot` yet (see follow-ups). |
| Public API | `New`, `SetData`, `Handle`, `Run`, `Serve`, `BindMobile`, root aliases. | Adds `Option`, `WithPerf`, `NewWithOptions`, `ServeOptions`, `ServeWithOptions`, `HoverHandler`, `SetWindowing`, `SetScrollWindow`, `SetScrollOffset`, `ScrollOffset`, `StepScrollWindow`, `SetPerf`, and `Perf`. Additive, no signature removed. |
| Examples | No perf fixtures. | `examples/perf-benchmarks` (normal, large, flappy apps plus `benchutil` and tests), `examples/perf-complex` (480-row dashboard, profiler, `analyze.py`, `run.sh`), `examples/perf-stress` (280-row dashboard with `SetTick`, headless dumps), and `examples/teams/channelbench`. `examples/readme.md` lists them. |
| Scripts | No perf scripts. | `scripts/perf-baseline.sh` (stage benches), `perf-idle.sh` (idle CPU/RSS), `perf-longrun.sh` (soak CSV), `perf-dump.sh` (headless JSON dumps into `temp/perf-dumps`). |
| Docs | No performance page. `documentation/README.md` had no perf row. | `documentation/performance.md`, `performance-baseline.md`, and `tricks-tips.md` added and indexed. `devtools.md`, `frames.md`, `repaint.md`, and `features.md` updated. `README.md` gains a Demo section with the preview video and the desktop-cat screenshot. |
| Plans | No perf plan. | `plans/v0.0.2/performance/` README plus phase-01 to phase-05, and `plans/v0.0.2/0.0.2/performance/` with the complex-layout analysis and checklist. |
| Instructions | `AGENTS.md` described pre-perf behavior. | `AGENTS.md` documents `NewWithOptions`/`WithPerf`, `WindowOptions.Perf`, and `ServeWithOptions` with gated pprof. It does not mention the Teams benchmark or `perf-complex` yet. |
| Deletions and splits | `internal/window/scrollbar.go` held the scrollbar code; `internal/page/page.go` and `form_change.go` carried code that later moved. | `scrollbar.go` splits into `scrollbar_thumb.go`; the bitmap fallback moves to `page_draw_fallback.go`; dirty helpers move to `page_dirty_geom.go`; all to stay under the 2000-character file limit. |

Measured deltas reported by the branch commits:

| Change | Before | After | Source |
|---|---|---|---|
| Stress dashboard ops per redraw | 1980 | 476 | commit 9e7767a |
| Stress dashboard cold redraw | 95.2 ms | 19.4 ms | commit 9e7767a |
| Stress dashboard alloc per redraw | 21 MB | 5.7 MB | commit 9e7767a |
| Complex layout redraw | 306.66 ms | 34.74 ms | commit b391d8e |
| Complex layout alloc per redraw | 86.25 MB | 9.27 MB | commit b391d8e |
| Desktop normal/large under Xvfb | panicked in Ebiten atlas | held 30 s+ | commit 9e7767a |

---

## Changes

### Perf opt-in and stats

- `options.go` adds `Option`, `WithPerf`, and `NewWithOptions`.
- `internal/page` adds `Config.Perf`, `SetPerf`/`Perf`, `page_perf.go` (alloc delta), and `page_perf_dirty.go` (dirty counters, `maxDirtyScan` 4096).
- `internal/window` adds `perf_opt.go` (page opt-in detection and hook wiring), `perf_ring.go` (120-frame ring), `perf_runtime.go` (500 ms sampler), and `perf_stage.go` (timed wrappers).
- `internal/host/inspect.go` grows `Stats` with stage times, dirty counts, and `AllocFrame`.

### DevTools panel

- `devtools_frame_perf.go` holds the Performance and Pipeline detail sections plus three package-level hooks.
- `devtools_frame_mem.go` holds the Memory section and byte formatting.
- `devtools_sync.go` extracts the per-frame devtools flag follow.
- Both hooks return `ok=false` when the shell has no Perf, so the panel shows `-` and the sampling cost is skipped.

### Row windowing and hover

- `internal/host/scrollwindow.go` defines `ScrollObserver` and `ScrollWindowStepper`.
- `internal/page/page_windowing.go` stores the offset and callback; `applyScrollWindow` runs before template execute; `StepScrollWindow` reuses an unchanged overscan.
- `internal/page/hover_handler.go` and root `hover.go` add the paint-only hover path.

### Dirty tracking and caches

- `page_boxcache.go` keeps the `id -> box` index and the memoized content rect.
- `page_dirty_diff.go` exits to a full frame past 8 changed ops; `page_dirty_match.go` presizes the pairing map; `page_dirty_geom.go` holds the shared geometry helpers.
- `content_size.go` in the window caches the content size per page generation.

### Oversize replay and viewport cache

- `replay_oversize.go` detects oversized content and draws ticks and oversized pages directly with `replay.DrawVisible`.
- `replay_viewport.go` and `replay_viewport_paint.go` keep a bounded window-sized buffer with overscan for oversized static pages, rebuild it when it stops covering the wanted area or the generation changes, and repaint just the dirty rect when it still covers the view.
- `internal/replay/visible.go` culls offscreen operations before replay.

### Web, scripts, and measurement

- `serve.go` adds `ServeOptions` and `ServeWithOptions`; `internal/web/serve_options.go` builds the mux and mounts pprof only with Perf; `debug_http.go` adds the `runtime` field and the `RuntimeSnapshot` hook.
- `scripts/perf-baseline.sh`, `perf-idle.sh`, `perf-longrun.sh`, and `perf-dump.sh` cover stage benches, idle capture, soak sampling, and headless JSON dumps.

### Examples

- `examples/perf-benchmarks`: normal (480x640 form), large (1000 rows), flappy (tick-driven op edits), each with a headless benchmark test and a `-web` mode.
- `examples/perf-complex`: 480-row dashboard with windowing on by default, seven timed redraw modes, CPU/heap profiles, compressed layout dumps, and `analyze.py`.
- `examples/perf-stress`: 280-row dashboard using `SetTick` and bound-op edits, with `TestDump` writing JSON for the soak scripts.
- `examples/teams/channelbench`: synthetic post benchmark that produced the tricks-tips numbers.

### Docs, plans, and instructions

- `documentation/performance.md` is the procedure; `performance-baseline.md` is the per-release table; `tricks-tips.md` records the row-window pattern.
- `plans/v0.0.2/performance/` holds the phase plan; `plans/v0.0.2/0.0.2/performance/` holds the complex-layout baseline.
- `AGENTS.md` and `documentation/README.md` point at the new entry points.

---

## Impact

| Area | Impact |
|------|--------|
| Performance | Small pages keep the same path when Perf is off. Windowing cuts the stress dashboard from 1980 to 476 ops and the complex fixture from 306.66 to 34.74 ms per redraw. Oversized pages replay instead of allocating a buffer that cannot exist. |
| Memory | Windowing drops the stress alloc from 21 to 5.7 MB per redraw; the complex fixture from 86.25 to 9.27 MB. The viewport path drops content-sized buffers. |
| Behavior / correctness | Oversized and ticking pages now render instead of panicking. Dirty repaints fall back to the full frame past 8 changed ops or a third of the frame area. Partial repaints stay covered by existing equality tests. |
| API / CLI | Additive public API only: `WithPerf`, `NewWithOptions`, `ServeOptions`, `ServeWithOptions`, `HoverHandler`, and the windowing methods. No released signature changed. |
| Dependencies | None. `go.mod`, `go.work`, and `examples/go.mod` are untouched by this branch. |
| Binary size / build time | Not measured. `make build` compiles every package, including the examples. |

---

## Breaking changes / migration

| Item | Migration |
|------|-----------|
| None against `master` | All new surface is additive. The `SetScrollWindow` signature changed once inside this branch (9e7767a then 073a169), before any release. |

---

## Test plan

- [x] `make build` (runs `go vet -p 1 ./... ./examples/...`) - passed locally on 2026-10-06
- [x] `make test` (runs `go test -p 1 ./... ./examples/...`) - passed locally on 2026-10-06
- [x] Branch commits report `make build` and `make test` green on the final sources (82ca687, b391d8e).
- [x] Desktop normal, large, and stress held 30 s+ after the oversize fallback, where normal and large panicked before (9e7767a).
- [ ] `sh scripts/perf-baseline.sh` and `sh scripts/perf-dump.sh` rerun to refresh `performance-baseline.md`.
- [ ] Heap and goroutine columns stay empty until `RuntimeSnapshot` is wired in production.

### Commands

```sh
make build
make test
sh scripts/perf-baseline.sh
```

Local verification for this PR document: run on 2026-10-06, branch HEAD `fa84ab1`. Results filled in under "Local run" below.

### Local run

```
make build: passed (go vet -p 1 ./... ./examples/...)
make test:  passed (go test -p 1 ./... ./examples/...)
```

---

## Screenshots / sample output

- The Demo section in `README.md` links the preview video and the desktop-cat screenshot.
- `documentation/performance-baseline.md` carries the v0.0.2 rows: idle RSS 28.6 MB (normal web), 52.0 MB (large web), 28.0 MB web / 164 MB desktop (flappy), cold redraws 0.63 / 88.76 / 0.40 ms, and stage benches DisplayList 169 us, Lay 7.7 ms, DrawRectScan 150-166 ns at 0 allocs.

---

## Related issues

- None. This repo has no issues or PRs for this work; `plans/v0.0.2/performance/` is the reference.

---

## PR metadata checklist (author)

- [ ] Self-assigned (`--assignee @me`)
- [ ] Labels applied (`enhancement`; add `documentation` for the docs and plans)
- [x] Filled body saved under `plans/v0.0.2/pr/pr-perf-observability.md`
- [ ] Related issues: none to link

Suggested command once the author decides to open the PR:

```sh
gh pr create \
  --base master \
  --head feature/perf-observability \
  --title "perf(window): opt-in sampling, row windowing, and oversize replay for large pages" \
  --body-file plans/v0.0.2/pr/pr-perf-observability.md \
  --assignee "@me" \
  --label enhancement
```

Do not run this until the author asks. The branch is already pushed but no PR exists yet.

---

## Follow-ups (out of scope)

The highest-value follow-ups:

1. Per-shell perf hooks instead of package globals (`internal/window/perf_opt.go:37` clobbers on a second window).
2. Compute frame percentiles in the 500 ms sampler, not three sorts per frame on the draw goroutine.
3. Wire `RuntimeSnapshot` to the window sampler, or delete the empty heap/goroutine columns.
4. Fix the perf scripts: `perf-baseline.sh` hides failed benches behind `tee`, `perf-dump.sh` never `cd`s to the repo root, `perf-idle.sh` prints `?` for the timestamp, and the documented positional commands do not parse.
5. Reconcile the docs with the code: devtools docs say DIRTY and a GC-runs row; the panel says Pipeline detail and has no GC row.

---

## Reviewer checklist

- [ ] Behavior matches the summary and the matrix.
- [ ] No unrelated changes in the diff.
- [ ] Public API additions documented (`WithPerf`, `NewWithOptions`, `ServeOptions`, windowing).
- [ ] New fallback paths have tests (oversize, viewport clip, perf wiring). Currently partial.
- [ ] PR has assignee and labels.
- [ ] No secrets or generated artifacts committed.
- [ ] Diff-stat-by-extension table below matches `bash scripts/pr-diff-stat.sh origin/master`.

---

## Diff stat by extension

Generated with `bash scripts/pr-diff-stat.sh origin/master` on 2026-10-06.

| Extension | Files | Insertions | Deletions |
| --- | ---: | ---: | ---: |
| `.go` | 114 | 4048 | 316 |
| `.html` | 9 | 16 | 0 |
| `.md` | 25 | 1060 | 14 |
| `.py` | 1 | 62 | 0 |
| `.sh` | 5 | 163 | 0 |
| **Total** | **154** | **5349** | **330** |

---

## Commit list

| Commit | Subject | Files | Insertions | Deletions |
|---|---|---|---|---|
| `6441769` | Add opt-in perf sampling, DevTools panel, benchmarks, and runbooks | 67 | +2288 | -152 |
| `82ca687` | Add perf validation loop: stress app, JSON dumps, hardening, phase plan | 32 | +954 | -142 |
| `558a619` | Fill perf baseline from live runs; fix soak script keys and parens | 3 | +41 | -15 |
| `9e7767a` | Add row windowing and oversize replay fallback for large pages | 17 | +395 | -45 |
| `b391d8e` | Add headless complex layout baseline and performance analysis | 11 | +534 | -0 |
| `073a169` | Improve oversized page scrolling at 1080p | 44 | +856 | -45 |
| `da91bd8` | Add tricks-tips doc and Teams channel benchmark | 4 | +81 | -0 |
| `fa84ab1` | Add Demo section with preview video and desktop-cat screenshot | 1 | +15 | -0 |
| `733de0f` | chore(plans): move version folders under v0.0.1 and v0.0.2 | 15 | +9 | -9 |
| `e02e2d8` | docs(pr): add the perf-observability PR body | 1 | +251 | -0 |
| tip | docs(pr): refresh the PR body stats and commit list | 1 | +8 | -5 |

Counts come from `git show <sha> --shortstat` on 2026-10-06.
