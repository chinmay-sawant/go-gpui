# Phase 01, baseline

Goal: write down what the page costs today, so every later improvement has something to beat.

## How to run the dump

Display-free benchmarks first. These need no window and run anywhere.

```sh
sh scripts/perf-baseline.sh --out /tmp/opencode/perf-baseline-$(date +%F).txt
```

That script runs `BenchmarkRedrawStages`, `BenchmarkRedrawWarm`, `BenchmarkRedrawCold`, and `BenchmarkRedrawHeavyWarm` in `internal/page`, `BenchmarkDrawRectScan` in `internal/replay`, and `BenchmarkClickCount` in `examples/platform`, all with `-benchmem`.

Then the idle watch, one app at a time. Leave the window untouched for the full sample.

```sh
sh scripts/perf-idle.sh --example examples/perf-benchmarks/normal --seconds 60
sh scripts/perf-idle.sh --example examples/perf-benchmarks/normal --web --seconds 60
```

Then the soak for the release scenario, 10 to 30 minutes.

```sh
sh scripts/perf-longrun.sh --minutes 15 --out /tmp/opencode/perf-soak.csv
```

Copy the results into the table below and into `documentation/performance-baseline.md`. Label each entry with the commit hash and machine name.

## Current numbers

Surveyed 2026-10-05, filled from headless runs on commit 82ca687 (plus the `perf-longrun.sh` key/paren fix below), machine Chinmay. Cold-redraw rows come from `temp/perf-dumps/dump-20261005T150517Z.json` via `sh scripts/perf-dump.sh`; stage rows from `temp/perf-baseline-2026-10-05.txt` via `sh scripts/perf-baseline.sh`; idle rows from 30 s desktop and `-web` runs on 2026-10-05; soak row from a 1-minute undriven `-web` soak of large. Frame-percentile and heap rows stay TBD for the reasons in the table. Prior art lives in the v0.0.2 history section of `documentation/performance-baseline.md`. A fourth app, stress-240 (1508 boxes, 1980 ops, 95.17 ms cold redraw, 91.46 ms display-list, 20.95 MB alloc), has no column yet; it becomes one when the table grows a stress column.

| Metric | Normal UI | Large UI | Flappy |
|---|---|---|---|
| Idle RSS, MB | 28.6 web; desktop crashes (see notes) | 52.0 web; desktop crashes (see notes) | 28.0 web; 164 desktop |
| Idle CPU, percent | ~0 web; no desktop reading | 0.5-3.5 web; no desktop reading | ~0 web; 16-23 desktop under Xvfb |
| Go heap, MB | TBD (runtime nil in prod, blocked on RuntimeSnapshot wiring) | TBD (same) | TBD (same) |
| Avg frame, ms | TBD (frame ring not exposed; needs DevTools read or endpoint) | TBD (same) | TBD (same) |
| p95 frame, ms | TBD (same) | TBD (same) | TBD (same) |
| p99 frame, ms | TBD (same) | TBD (same) | TBD (same) |
| Alloc per cold redraw | 373 KB | 26.15 MB | 362 KB |
| Cold redraw, ms | 0.63 | 88.76 | 0.40 |
| Display-list build, ms | 0.47 | 84.29 | 0.30 |
| Layout, ms | 0.08 | 1.66 | 0.045 |
| Paint, ms | 0 (vector path; bitmap Lay 4.0-7.7 ms on small fixture) | 0 (vector path) | 0 (vector path) |
| Soak verdict | 1-min undriven web soak: flat; desktop soak blocked by atlas crash | 1-min undriven web soak: counters frozen at 1/1/1/1/1, no errors; heap columns empty (runtime nil); 15-min driven soak still open | 1-min numbers via large run; flappy desktop held 30 s with no growth |
| Commit | 82ca687 + perf-longrun.sh key/paren fix | 82ca687 + perf-longrun.sh key/paren fix | 82ca687 + perf-longrun.sh key/paren fix |
| Machine | Chinmay | Chinmay | Chinmay |

Stage benches from `temp/perf-baseline-2026-10-05.txt` (`scripts/perf-baseline.sh`, same commit): TemplateExecute 578 ns, HTMLParse 1493 ns, CSSApply 4432 ns, DisplayList 169 us, Lay 7.7 ms, RedrawWarm 143 us, RedrawCold 163 us, RedrawHeavyWarm 333 us, DrawRectScan 150-166 ns at 0 allocs, ClickCount 607 us.

Blockers found while running, all reproduced: desktop windows with tall content panic in Ebiten (`atlas: the image being put on an atlas is too big`, normal 480x16494 within 30 s, large 800x36579 within 10 s) because the persistent buffer is sized to the content; `perf-longrun.sh` had a missing close paren and lowercase stat keys, so every sample wrote `fetch-failed` until fixed; `/debug/state` runtime is always nil outside tests, so heap and goroutine columns stay empty on `-web`.

## Checklist

- [x] Run `scripts/perf-baseline.sh` and paste the bench output next to this table. Done 2026-10-05, output in `temp/perf-baseline-2026-10-05.txt`, stage numbers quoted under the table.
- [x] Run `scripts/perf-idle.sh` for each of normal, large, and flappy, desktop and `-web`. Done 2026-10-05 at 30 s per capture (not 60 s): desktop normal and large crash in Ebiten with the atlas-too-big panic, desktop flappy holds with CPU 16-23 percent and RSS 164 MB under Xvfb, all three `-web` runs sit at redraws 1 with flat RSS. Go heap and frame percentiles could not be captured headless, see table.
- [x] Run `scripts/perf-longrun.sh` for the release soak scenario and record the verdict. Done 2026-10-05 at 1 minute, undriven, `-web` large: counters frozen at 1/1/1/1/1 with no errors, heap columns empty because runtime is nil in prod. Required fixing a missing close paren and lowercase stat keys in the script first. A 15-minute driven soak and any desktop soak stay open, blocked by the atlas crash.
- [x] Fill every Commit and Machine cell. Done: 82ca687 plus the script fix, machine Chinmay.
- [ ] Mirror the filled table into `documentation/performance-baseline.md`.
