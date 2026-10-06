# Phase 04, analysis

Goal: read the dumps, name the slowest stage, and back it with a profile before anyone changes code.

## How to re-analyze

Run the phase 03 dump against the app and scenario in question, then open the worst sample.

1. Compare `stats.lastTemplate`, `layoutTime`, `displayListTime`, and `paintTime`. The largest stage is the suspect.
2. Check `dirtyOps`, `dirtyRegions`, and `changedOps`. A full-frame rect on a small edit means the diff fell back. More than eight changed ops or coverage over a third of the frame are the known fallback triggers in `documentation/repaint.md`.
3. Check `allocFrame` and `runtime.heap_alloc`. Allocation on a static page points at per-frame work that should not exist.
4. Take a profile in web mode with Perf on. Use the port the app actually serves.

```sh
go tool pprof -top http://127.0.0.1:8091/debug/pprof/heap
go tool pprof -top http://127.0.0.1:8091/debug/pprof/profile?seconds=30
```

5. Change one thing, re-dump with the same script and app, and quote before and after numbers. No optimization lands without that pair.

## Bottleneck findings

Surveyed 2026-10-05. No dumps exist yet, so every slot is open. Fill one row per finding with the dump file, the profile, and the commit hash.

| # | Stage | Evidence | Dump file | Profile | Status |
|---|---|---|---|---|---|
| 1 | Template execute | TBD | TBD | TBD | open |
| 2 | Layout / relayout | TBD | TBD | TBD | open |
| 3 | Display list build | TBD | TBD | TBD | open |
| 4 | Paint / replay | TBD | TBD | TBD | open |
| 5 | Dirty diff fallback | TBD | TBD | TBD | open |
| 6 | Per-frame allocation | TBD | TBD | TBD | open |

## Checklist

- [ ] Re-dump the failing scenario from phase 02 with `scripts/perf-dump.sh`.
- [ ] Name the stage from the `stats` split, not from a guess.
- [ ] Attach a CPU or heap profile for any finding marked confirmed.
- [ ] Move each confirmed finding into phase 05 with file pointer and expected gain.
- [ ] Re-analyze after every phase 05 fix lands, same app and script.
