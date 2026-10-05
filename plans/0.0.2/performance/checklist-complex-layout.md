# Complex layout profile and optimization checklist

Scope: build a display-free fixture, capture repeatable dumps and profiles, analyze costs, and estimate optimization gains. Production optimization implementation is a follow-up.

Requested location is `plans/0.0.2/performance`. Earlier branch ledgers live under `plans/v0.0.2/performance`; this analysis uses the requested location.

## Phase 1: establish the baseline

- [x] Inspect current branch, commits, existing examples, and previous performance changes. Clean starting tree at `9e7767a`, branch `feature/perf-observability`.
- [x] Read existing performance docs, dump tooling, and the checklist skill.
- [x] Define the measurement boundary. HTML execution, CSS, layout, display operations, box indexing, and dirty diff run. No window, GPU replay, PNG, PDF, browser, or server runs.

## Phase 2: create the fixture

- [x] Add an operations dashboard with six navigation entries, six metric cards, two summary panels, form controls, and 480 service rows. Each row has nested flex metrics, wrapped text, status, and six spark bars.
- [x] Add cached, fresh layout, one-label data mutation, resize, and 48-row window scenarios.
- [x] Verify operation counts, fixed row heights, retained content height, and visible row geometry between full and windowed layouts.

## Phase 3: capture evidence

- [x] Run seven samples per scenario in three fresh processes. Keep initial layouts separate from cached redraws.
- [x] Save raw timing, allocation, stage counters, hit boxes, display-operation geometry, CPU profiles, and heap profiles.
- [x] Record source hash, revision, toolchain, host, commands, and profile boundaries.

## Phase 4: analyze and estimate

- [x] Identify costs from profiles and source, including costs inside the pinned layout engine.
- [x] Compute windowing gains from measured evidence.
- [x] Estimate remaining proposals with explicit assumptions, CPU ceilings, and validation gates. Avoid adding overlapping gains.
- [x] Write the report with an ordered optimization backlog.

## Phase 5: close the analysis

- [x] Format edited Go files and confirm every file is at most 2000 characters.
- [x] Run `make test` and `make build`.
- [x] Record unresolved runtime and correctness gates.

## Follow-up implementation gates

- [ ] Windowing: validate scroll positions, resize, focus, keyboard traversal, search, selection, and variable-height rows before using this fixture strategy in an app.
- [ ] Engine changes: compare boxes and ordered operations, text wrapping and font handling; rerun the engine's own correctness gates and paired benchmarks before updating the dependency.
- [ ] Redraw avoidance: prove every visible state change still reaches the display and dirty-region path.
- [x] Run desktop replay profiles before making scroll FPS and process resource claims. Idle CPU is not measured.

## Completion evidence

- Analysis: [complex-layout-analysis.md](complex-layout-analysis.md).
- Repeated capture and geometry checks: [summary.json](../../../temp/complex-dump/summary.json) and [analysis.txt](../../../temp/complex-dump/analysis.txt). Fifteen fresh processes, seven timed samples each. Initial sample zero is separate from warm process samples.
- Dumps and profiles: `temp/complex-dump/run-1`, `run-2`, and `run-3`. Every measured redraw has zero bitmap paint time and a retained display list.
- Windowing result: 306.66 to 34.74 ms, 86.25 to 9.27 MB allocated, 12,360 to 1,344 operations. All three structural geometry checks passed.
- Existing app capture: [existing-apps.txt](../../../temp/complex-dump/existing-apps.txt), with the JSON dump under `temp/complex-dump/existing-apps`. Stress 240 rows is already windowed and measures 476 operations; large 1000 measures 4,003 operations. These one-shot captures are context, not repeated performance claims.
- `make test` exited 0 on the final Go source. [Log](../../../temp/complex-dump/make-test.txt).
- `make build` exited 0 on the final Go source. [Log](../../../temp/complex-dump/make-build.txt).
- `gofmt` is clean. Five added Go files range from 342 to 1,714 characters. `git diff --check` exited 0 and all added text files passed a trailing-whitespace check.
- Production engine optimizations, live scrolling, pixel comparison, desktop replay, frame percentiles, RSS and soak checks remain follow-up gates. They are not required to complete this dump analysis.

Generated dumps, profiles, summaries and validation logs were moved to `temp/complex-dump`. File hashes matched before and after the move. The capture and analysis scripts now use that directory.

## Runnable example

- [x] Move the shared fixture and profiling tools to [examples/perf-complex](../../../examples/perf-complex/README.md).
- [x] Add desktop viewing with `go run ./examples/perf-complex`, live row windowing with `-windowed`, and browser viewing with `-web`.
- [x] Keep headless profiling available with `-dump` and `sh examples/perf-complex/run.sh`. Generated output stays in `temp/complex-dump`.
- [x] Validate the moved example and its scroll window; run `make test` and `make build`.

The baseline HTML is [layout.html](../../../examples/perf-complex/layout.html). The full and windowed desktop modes share it with the dump scenarios.

Example validation: `make test` and `make build` passed, with logs in `temp/complex-dump/make-test-example.txt` and `make-build-example.txt`. The new test compares row geometry/text and retained grid height at the top, middle and bottom. Full and windowed desktop modes both stayed open under Xvfb; screenshots and logs are in `temp/complex-dump/desktop-full.*` and `desktop-windowed.*`. The windowed screenshot was inspected. A moved-command `-dump` smoke run completed under `temp/complex-dump/example-smoke`. The baseline HTML matches the original fixture byte for byte.

## Desktop scroll follow-up

- [x] Reproduce the reported slowdown with the actual Ebiten loop at 1908 × 999 and 1920 × 1080.
- [x] Filter offscreen operations while preserving order and unbounded ink.
- [x] Add one bounded viewport cache with overscan, dirty repaint, and disposal on resize/path changes.
- [x] Avoid hover layout for row background changes; keep ordinary CSS fallback for controls.
- [x] Reuse row overscan instead of rebuilding on each scroll update; migrate current scroll callbacks to return whether data changed.
- [x] Disable antialiasing for flat rectangles already snapped to pixels. GPU pixel parity passed against the earlier renderer at five scroll offsets.
- [x] Measure paired scroll FPS, process RSS, and CPU at 1080p. 0.77 → 54.25 FPS; peak RSS 306.62 → 209.44 MB; mean CPU 44.65 → 32.82% of one core. See `temp/complex-dump/1080-resources.txt`.
- [ ] Complete the longer resource soak and record the memory trend.
- [x] Run `make build`, formatting, and Go size gates. `make build` passed with `GOCACHE=/tmp/go-gpui-scroll-cache`.
- [x] Run targeted tests for root, page, replay, window, perf-complex and perf-stress packages; all passed.
- [x] Run the full `make test`; existing fetch, music and player tests requiring loopback TCP cannot run in this restricted sandbox. Other completed package results are in `temp/complex-dump/make-test-scroll-writable.txt`.
- [ ] Confirm scrolling manually on the user's desktop. Xvfb software rendering does not establish hardware frame pacing or a guaranteed 60 FPS.

The earlier completion logs cover the earlier analysis and example move. Final scroll validation logs use the `*-scroll-final.txt` suffix; targeted tests are recorded in `targeted-tests-scroll-final.txt` and `make-build-scroll-final.txt`. Variable-height rows, focus retention, and selection across removed rows remain outside this fixture's fixed-height windowing design.

Current validation boundary: paired desktop measurements and GPU parity passed before the environment switched to restricted execution. The longer soak cannot start because Xvfb cannot bind its local sockets. With a writable cache, `make build` and targeted tests passed. Full `make test` was attempted; tests requiring loopback listeners failed because socket creation is prohibited by the sandbox. Formatting, the 2000-character gate for 34 changed/new Go files, and `git diff --check` passed. Logs are in `temp/complex-dump`; no new measurements from the failed soak are claimed.
