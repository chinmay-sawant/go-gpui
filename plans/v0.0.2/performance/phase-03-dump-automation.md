# Phase 03, dump automation

Goal: one script that captures the full debug state on demand, so every analyze and improve step reads the same file format.

## Status

Surveyed 2026-10-05. `scripts/perf-dump.sh` does not exist, `temp/perf-dumps/` does not exist, and no `page_perf_dump.go` exists anywhere in the repo. The schema below is real. It matches `GET /debug/state` in `internal/web/debug_http.go` plus `host.Stats` in `internal/host/inspect.go`. The script spec is a build task for whoever picks up the checklist.

## Dump schema

`GET /debug/state` returns this JSON. The dump script must store it verbatim and never reshape it.

| Field | Source | What it tells |
|---|---|---|
| width, height | display or image bounds | frame size the numbers belong to |
| generation | page generation | whether the page moved between samples |
| fallback | display nil, image set | bitmap path, full repaints only |
| boxes | `layout.Box` list | box count and geometry |
| stats.redraws, parses, cascades, layouts, repaints | `host.Stats` | pipeline stage counts |
| stats.boxes, ops | `host.Stats` | scene size |
| stats.lastRedraw, lastDraw | `host.Stats` | last frame cost |
| stats.lastTemplate, layoutTime, displayListTime, paintTime | `host.Stats` | stage split of the last redraw |
| stats.dirtyOps, dirtyRegions, changedOps | `host.Stats` | dirty tracking detail |
| stats.allocFrame | `host.Stats` | bytes allocated during the last redraw |
| stats.reloads, lastReloadError | `host.Stats` | hot reload noise during the sample |
| ops | per-kind display op counts | operation mix, present only with a display list |
| runtime | `RuntimeSnapshot` | heap, RSS, goroutines, GC runs, present only when Perf is on |

`runtime` is absent unless the app opted into Perf. A dump with no `runtime` key means the app ran without `WithPerf(true)`, not that memory is zero. Record that in the notes row.

## `scripts/perf-dump.sh` usage

Intended interface, modeled on the existing `perf-idle.sh` and `perf-longrun.sh` flags.

```sh
sh scripts/perf-dump.sh --url http://127.0.0.1:8128/debug/state --out temp/perf-dumps/
sh scripts/perf-dump.sh --url http://127.0.0.1:8128/debug/state --out temp/perf-dumps/ --samples 12 --interval 5
```

Behavior: poll the URL, write one timestamped JSON file per sample plus one CSV rollup with the columns `ts, redraws, parses, cascades, layouts, repaints, heap_alloc, goroutines`, same columns as the soak script. Filenames carry UTC time, app name, and commit hash. A failed fetch writes a `fetch-failed` row instead of stopping the run.

## Checklist

- [ ] Create `temp/perf-dumps/` and confirm it is ignored or scratch (check `.gitignore` before committing samples).
- [ ] Write `scripts/perf-dump.sh` with the flags above. Keep it POSIX sh like the other perf scripts.
- [ ] Dump each benchmark app plus the phase 02 stress app, desktop and `-web`.
- [ ] Confirm one dump per app contains `stats` and `runtime`. Missing keys mean Perf is off. Fix the app flags, not the schema.
- [ ] Store the dumps next to the baseline entry they support.
