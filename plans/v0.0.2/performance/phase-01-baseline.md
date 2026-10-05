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

Surveyed 2026-10-05, filled from headless runs on commit 6441769 (plus uncommitted hardening in the working tree), machine Chinmay. Cold-redraw rows come from `temp/perf-dumps/dump-20261005T150517Z.json` via `sh scripts/perf-dump.sh`; stage rows from `BenchmarkRedrawStages` rerun the same day. Idle, frame-percentile, and soak rows stay TBD because they need a live window (`scripts/perf-idle.sh`, `scripts/perf-longrun.sh`), which no worker ran. Prior art lives in the v0.0.2 history section of `documentation/performance-baseline.md`. A fourth app, stress-240 (1508 boxes, 1980 ops, 95.17 ms cold redraw, 91.46 ms display-list, 20.95 MB alloc), has no column yet; it becomes one when the table grows a stress column.

| Metric | Normal UI | Large UI | Flappy |
|---|---|---|---|
| Idle RSS, MB | TBD (needs live window) | TBD (needs live window) | TBD (needs live window) |
| Idle CPU, percent | TBD (needs live window) | TBD (needs live window) | TBD (needs live window) |
| Go heap, MB | TBD (needs live window) | TBD (needs live window) | TBD (needs live window) |
| Avg frame, ms | TBD (needs live window) | TBD (needs live window) | TBD (needs live window) |
| p95 frame, ms | TBD (needs live window) | TBD (needs live window) | TBD (needs live window) |
| p99 frame, ms | TBD (needs live window) | TBD (needs live window) | TBD (needs live window) |
| Alloc per cold redraw | 373 KB | 26.15 MB | 362 KB |
| Cold redraw, ms | 0.63 | 88.76 | 0.40 |
| Display-list build, ms | 0.47 | 84.29 | 0.30 |
| Layout, ms | 0.08 | 1.66 | 0.045 |
| Paint, ms | 0 (vector path; bitmap Lay 4.0 ms on small fixture) | 0 (vector path) | 0 (vector path) |
| Soak verdict | TBD (needs live window) | TBD (needs live window) | TBD (needs live window) |
| Commit | 6441769 + local hardening | 6441769 + local hardening | 6441769 + local hardening |
| Machine | Chinmay | Chinmay | Chinmay |

## Checklist

- [ ] Run `scripts/perf-baseline.sh` and paste the bench output next to this table.
- [ ] Run `scripts/perf-idle.sh` for each of normal, large, and flappy, desktop and `-web`.
- [ ] Run `scripts/perf-longrun.sh` for the release soak scenario and record the verdict.
- [ ] Fill every Commit and Machine cell. No number without a hash.
- [ ] Mirror the filled table into `documentation/performance-baseline.md`.
