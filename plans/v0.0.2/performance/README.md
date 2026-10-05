# Performance validation loop

Recorded 2026-10-05. This folder turns `documentation/performance.md` into a repeatable loop with five phases and a checklist per phase.

Status: phase checklists are open. Baseline numbers are TBD until the dump automation in phase 03 exists. Missing pieces found during the survey are recorded as checklist items, not assumed.

## The loop

```
complex-UI -> dump -> analyze -> improve -> re-dump
     ^                                            |
     +--------------------------------------------+
```

Phase 02 builds the complex UI. Phase 03 dumps numbers from it. Phase 04 names the bottleneck. Phase 05 fixes one thing. Then the loop returns to phase 03 and re-dumps with the same app, same machine, same script. A change lands only when the re-dump beats the prior dump.

## Phases

| Plan | Job |
|---|---|
| [phase-01-baseline.md](phase-01-baseline.md) | Record current numbers or TBD slots, one row per metric |
| [phase-02-complex-ui.md](phase-02-complex-ui.md) | Build the stress example and define pass criteria |
| [phase-03-dump-automation.md](phase-03-dump-automation.md) | Define the dump schema and the `perf-dump.sh` script |
| [phase-04-analysis.md](phase-04-analysis.md) | Slot in bottleneck findings and re-analyze on demand |
| [phase-05-improvements.md](phase-05-improvements.md) | Ordered backlog, each item with file, gain, and verify command |

## Thresholds

A 60 fps window has 16.67 ms per frame. These apply to the complex UI on the reference machine once phase 01 fills in actuals.

| Metric | Pass | Investigate |
|---|---|---|
| p99 frame time | under 16.67 ms | over budget on any app |
| p95 frame time | under 12 ms | over budget twice in a row |
| Idle CPU, untouched window | near 0 percent | sustained above 2 percent |
| Idle RSS across 60 s | flat | climbs through the window |
| Single-control click | 1 dirty rect, rest skipped | full-frame fallback |
| Soak, 15 min | heap and RSS flat after warmup | steady climb, goroutines climb |
| Alloc per frame, non-ticking page | near 0 after warmup | grows with op count |

## Ground rules

- Markdown only in this folder. No code changes land from these checklists. Code work goes through phase 05 items with their own verify commands.
- Measure on one machine per release. Label every number with the commit hash. Numbers from different machines do not compare.
- Perf sampling stays opt-in. Dumps assume `WithPerf(true)` on the page and `ServeOptions{Perf: true}` in web mode, per `documentation/performance.md`.
- Leave cells empty when a script or counter is unavailable. Say so in the notes row instead of guessing.

## Definition of done

- Phase 01 has one filled table per benchmark app, or TBD slots with the reason named.
- Phase 02 has a stress example that runs on desktop and with `-web`, plus pass criteria per scenario.
- Phase 03 has a dump script that writes timestamped files under `temp/perf-dumps/`, and the schema matches `GET /debug/state`.
- Phase 04 names the stage for the worst number, with profile evidence.
- Phase 05 has every fix quoted with before and after numbers from the same script and app.
