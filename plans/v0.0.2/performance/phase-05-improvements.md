# Phase 05, improvements

Goal: an ordered backlog where each item names the file, the expected gain, and the command that proves it.

## Rules

- One fix per item. Re-dump after each one with the same script and app.
- Quote before and after numbers in the commit message, per `documentation/performance.md`.
- Keep the profile files next to the baseline entry when the change lands.

## Backlog

Surveyed 2026-10-05. No phase 04 findings exist, so these are candidates ordered by likely gain on the stress app. Promote confirmed findings above these. Drop any candidate the profiles clear.

| # | Fix | File pointer | Expected gain | Verify command |
|---|---|---|---|---|
| 1 | Shrink the dirty fallback rate on repeated-row edits, so add and remove repaint rows not frames | `internal/page` dirty and diff path, `internal/replay` `DrawRect` | fewer full-frame fallbacks in `stats`, lower paint time | `sh scripts/perf-baseline.sh` plus `sh scripts/perf-dump.sh --url http://127.0.0.1:8128/debug/state --out temp/perf-dumps/` before and after |
| 2 | Cut per-frame allocation on ticking pages, reuse buffers across frames | `internal/replay` draw path, `internal/frame` op lookup | lower `allocFrame` and flatter heap on the flappy and stress apps | `go tool pprof -top http://127.0.0.1:8091/debug/pprof/heap` plus phase 03 dump before and after |
| 3 | Keep hover and press on the cached relayout path, no reparse or recollect on pointer moves | `internal/render` cached document, `internal/page` state changes | `parses` and `cascades` stay flat during pointer-heavy scenarios | `sh scripts/perf-longrun.sh --minutes 15` with click scenario, watch `parses` and `cascades` columns |
| 4 | Narrow drag relayout cost, one layout per committed size at most | `internal/window` drag throttle and `internal/page` size path | layout time bounded during resize soak | `sh scripts/perf-longrun.sh --minutes 15` with resize scenario, compare `layouts` and p99 |
| 5 | Trim text shaping and replay cost on text-heavy rows | `internal/replay` text runs, `internal/textrun` offsets | lower paint time on the stress app list | `sh scripts/perf-baseline.sh` plus stress app dump before and after |
| 6 | Confirm the bitmap fallback stays rare, log when it triggers | `internal/window/draw.go` display-or-fallback branch | fallback count near zero outside known cases | phase 03 dump, check `fallback` key across samples |

## Checklist

- [ ] Confirm or replace each candidate with a phase 04 profile.
- [ ] Land fixes one at a time, each with before and after numbers.
- [ ] Re-run `make test` after each fix, per the repo ground rules.
- [ ] Close the loop: re-dump, update phase 01, mark the README thresholds pass or fail.
