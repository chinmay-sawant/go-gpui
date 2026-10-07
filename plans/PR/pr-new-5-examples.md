# feat(examples): add five desktop apps and partial tick replay

Filled copy of [`skills/PR/PR_TEMPLATE.md`](../../skills/PR/PR_TEMPLATE.md).

- Base: `origin/master` (`aed3e0a`)
- Head: `chore/new-5-examples`
- Diff against `origin/master`: 1042 files before this plan file

---

## Summary

This branch adds five desktop apps: a download manager, a live log viewer, a spreadsheet, a system monitor, and Tetris. Each one keeps its own state, starts from dummy data, and is indexed as desktop-only. It also lets a ticking page repaint a dirty rectangle. The download manager uses that so a 1908x999 progress frame copies a cached picture instead of replaying the whole list, which was missing vsync and showing 12 fps.

---

## Motivation / context

- The five apps were planned under `plans/v0.0.2/examples/` and were not on `master`.
- A ticking page always replayed the full display list. At 1908x999 the download manager's progress tick cost about 15 ms of replay, and a real list rebuild cost about 25 ms. Together those missed a 16.7 ms frame and landed near 80 ms, which the frame tab reports as 12 fps.
- Issues: none. This work has no tracker issue.

---

## Changes

### Five desktop examples

- `examples/tetris`: SRS rotation, 7-bag randomizer, fixed-step tick, held-key repeat, next preview, SQLite scores.
- `examples/live-log-viewer`: multi-file tail, filters, follow and pause, entry detail, bounded export, SQLite history.
- `examples/download-manager`: bounded workers, pause and resume, retry, cancel, durable history.
- `examples/system-monitor`: CPU, memory, disk, and network panels, paged processes, dummy or live collectors.
- `examples/spreadsheet`: sparse grid, formula bar, SUM and AVERAGE, undo, CSV, SQLite workbooks.
- `examples/readme.md` lists all five as desktop-only. Light and dark theme contrast is repaired in the same apps.

### Partial tick replay

- `Page.UseFrameDirty` opts one tick into the partial buffer. `Page.MarkRect` names a CSS-pixel strip that has no element id. `Page.FrameDirty` reports the opt-in and clears it.
- A tick that never calls `UseFrameDirty` still replays the whole list.
- The download manager calls `UseFrameDirty` every tick, marks the progress row that changed, and marks the header counts, detail body, and footer on separate frames. One dirty rect is the union of that frame, so a header and a footer together would repaint the history between them.
- `documentation/frames.md`, `documentation/repaint.md`, and `documentation/performance.md` describe the opt-in.

### Tests

- Each example has headless coverage for its core, storage, and UI.
- `TestProgressDirtyStaysOnTheRow` checks that a progress tick at 1908x999 stays within an 80 px dirty rect.
- `TestEndToEndDummy` waits until a manually added URL appears in the active set or on the history page. A short dummy file can finish before a fixed tick count.

---

## Impact

| Area | Impact |
|------|--------|
| **Performance** | Download-manager progress frames repaint one row and blit the cache. Other ticking pages are unchanged. A frame that changes the list shape still rebuilds the display list. |
| **Memory** | An opted-in tick keeps the same content-sized buffer a non-ticking page already used. |
| **Behavior / correctness** | Pause, select, and remove still receive the click. The row strip has no element id. Theme colors for progress and graph fills are unchanged. |
| **API / CLI** | Adds `UseFrameDirty`, `FrameDirty`, and `MarkRect` on `Page`. No signature removed. |
| **Dependencies** | None. |
| **Binary size / build time** | Five example trees. The library change is three small files. |

---

## Breaking changes / migration

| Item | Migration |
|------|-----------|
| None | Existing pages keep the full replay until they call `UseFrameDirty`. |

---

## Test plan

- [x] `make test` passed on `f93b980` (`go test -p 1 ./... ./examples/...`)
- [ ] `make lint` / `go vet` was not run as its own target
- [ ] `make build` was not run
- [ ] `make run` does not apply to these headless example tests
- [ ] `make reference-metrics` does not apply. No detector surface changed.

### Commands

```sh
make test
```

---

## Screenshots / sample output

No new screenshot. The five apps are covered by headless tests. A live 1908x999 frame-tab reading was not taken in this change.

---

## Related issues

- None. No issue is closed or linked by this PR.

---

## PR metadata checklist (author)

- [x] Self-assigned (`--assignee @me`)
- [x] Labels applied
- [x] Related issues filled: there is no tracker issue
- [x] Filled body committed under `plans/PR/pr-new-5-examples.md`

---

## Follow-ups (out of scope)

- Measure Draw on a live 1908x999 window. A shape-changing frame can still miss one refresh.
- The footer counter can lag one frame while progress rows are moving.

---

## Reviewer checklist

- [ ] Behavior matches summary and test plan
- [ ] No unrelated changes in diff
- [ ] Public API / CLI changes documented
- [ ] New rules have fixture coverage when applicable
- [ ] PR has assignee and labels
- [ ] Related issues use correct Closes/Relates keywords
- [ ] No secrets or generated artifacts committed
- [ ] Diff-stat-by-extension table pasted at the bottom

---

## Diff stat by extension

| Extension | Files | Insertions | Deletions |
| --- | ---: | ---: | ---: |
| `.css` | 8 | 323 | 0 |
| `.go` | 1010 | 58678 | 1 |
| `.html` | 9 | 618 | 0 |
| `.md` | 16 | 1342 | 5 |
| **Total** | **1043** | **60961** | **6** |
