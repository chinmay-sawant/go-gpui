# Performance baseline

One table per release. Fill it on the same machine when you can, and label every entry with the commit hash. Numbers from different machines do not compare.

Full procedure lives in [performance.md](performance.md).

## How to fill it

Run the automated script first. It exercises each benchmark app in `examples/perf-benchmarks` and prints the rows below.

```sh
sh scripts/perf-baseline.sh
```

Then run the manual steps. Do the 30 to 60 second idle watch from [performance.md](performance.md) for each app, and the soak scenario that matches the release. Write the soak verdict and the machine name by hand.

Copy the template section for each release. Leave cells empty when a script or counter is unavailable, and say so in the notes row instead of guessing.

## Template

| Metric | Normal UI | Large UI | Flappy |
|---|---|---|---|
| Idle RSS, MB | | | |
| Idle CPU, percent | | | |
| Go heap, MB | | | |
| Avg frame, ms | | | |
| p95 frame, ms | | | |
| p99 frame, ms | | | |
| Alloc per frame, KB | | | |
| Layout, ms | | | |
| Paint, ms | | | |
| Soak verdict | | | |
| Commit | | | |
| Machine | | | |
| Notes | | | |

## v0.0.2 prior numbers

These rows predate the baseline scripts. They come from `plans/v0.0.2/release-notes.md` and `repaint.md` at 480x640 and 640x480 on the author's machine. They are machine specific and use older harnesses, so treat them as history, not as targets.

| Metric | Value | Source |
|---|---|---|
| css Apply, 200 rules | about 200 us | release notes pipeline split |
| css Relayout, cached sheets | about 1.2 us | release notes pipeline split |
| Redraw heavy warm | 489 to 708 us, down from 1.34 to 1.83 ms | `BenchmarkRedrawHeavyWarm` |
| DrawRect op filter | about 12 ns per op | `BenchmarkDrawRectScan` in `internal/replay` |
| Click through handler, execute, relayout, dirty rect | about 0.5 ms | `BenchmarkClickCount` in platform example |
| platform `#count`, 11 ops | 5 repainted, 6 skipped | repaint.md measured numbers |
| login `#message`, 18 ops | 3 repainted, 15 skipped | repaint.md measured numbers |
| states `#hover-btn`, 16 ops | 4 repainted, 12 skipped | repaint.md measured numbers |
| forms `#remember`, 45 ops | 5 repainted, 40 skipped | repaint.md measured numbers |
