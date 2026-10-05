# Phase 02, complex UI

Goal: a stress example that pushes layout, paint, and dirty tracking past what the three benchmark apps cover, so the dump in phase 03 has a hard case.

## Status

Surveyed 2026-10-05. `examples/perf-stress/` does not exist. `examples/perf-benchmarks/` has normal, large, and flappy with shared code in `benchutil`, each runnable on desktop and with `-web`. The spec below describes the missing stress app. Building it is code work and belongs to whoever picks up the checklist, not to this plan.

## Stress example spec

Location: `examples/perf-stress/`, one module path under the existing examples module. It must run both ways.

```sh
go run ./examples/perf-stress
go run ./examples/perf-stress -web
```

Contents, each chosen to strain one stage:

- A thousand-row list with varied nesting depth, to strain layout and the display list build.
- A large stylesheet with several hundred rules including `@media` width blocks, to strain cascade and `css.Relayout`.
- Text-heavy rows with mixed scripts, to strain shaping and replay text runs.
- One ticking region (equalizer or seek bar style, like the flappy app) that edits a few operations per frame, to strain per-frame allocation.
- One control that adds and removes repeated rows or cards, to strain cache invalidation.
- One text field with selection, to keep the caret and selection spans in the measured path.

Perf opt-in from the start, following `documentation/performance.md`. The app sets `WithPerf(true)` and serves pprof in web mode through `ServeOptions{Perf: true}`.

## Pass criteria

Run the phase 03 dump against the stress app on the reference machine and check each row.

| Scenario | Pass |
|---|---|
| Idle 60 s, no input | CPU near 0 percent, RSS flat |
| Slow window-edge drag | one relayout per committed size, no full reparse per mouse event |
| Wheel top to bottom and back | no long-frame burst, scroll stays smooth |
| Single control toggle | 1 dirty rect, small repainted count, large skipped count |
| Typing with selection | caret follows the glyph, no full-frame fallback per keystroke |
| Tick animation on | p99 frame under 16.67 ms |
| Add and remove 100 rows | heap returns to pre-add level after GC |
| 15 min mixed soak | RSS and goroutines flat after warmup |

## Checklist

- [ ] Create `examples/perf-stress/` with the contents above, desktop and `-web` runs.
- [ ] Opt into Perf in the app so the Frame tab and `GET /debug/state` expose counters.
- [ ] Drive each scenario in the table and record pass or fail with the commit hash.
- [ ] File failures as phase 04 findings with the scenario name attached.
