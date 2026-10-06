## Summary

Ship go-gpui v0.0.2: the window keeps the parsed and styled page between frames, repaints only the dirty region, and gains a Chrome-style DevTools dock, hot reload, drag and drop, printing, packaging, and a full input layer. The branch is 108 commits and 818 files ahead of `master` (+32,231 / −3,939). The engine adds `css.Relayout`; `go.mod` drops the local `replace` and pins the pushed `chore/changes-for-go-gpui` commit `1a3918301a68` as `v0.2.7-0.20261004151708-1a3918301a68`. The library still writes no PDF.

## Motivation / context

- Plans: `plans/v0.0.2/README.md` defines the six workstreams; every phase checklist is closed (249/249 boxes).
- Two of the six share one root cause: `Page.Redraw` kept nothing it computed, so every click, hover, drag, and resize paid a full template execute, `html.Parse`, `css.Apply`, whole-page layout, and whole-page replay.
- Review pages: `plans/v0.0.2/report-html-2.0.0/v0.0.2-report.html` (six verified parts in one file).
- No issue is linked; the repo has no issues, same as v0.0.1.

## Changes

### Pipeline split and the engine (foundation)

- `internal/render` (new): `Cache` holds the executed source, the parsed `*html.Document`, and the styled `*css.Document` from one parse and one cascade. `Redraw` compares the executed source; a hit calls `Cache.Relayout` and lays the cached document out again, a miss builds a new cache and counts one parse and one cascade.
- `SetSize` and state changes (hover, press, release, focus, theme) relayout without reparsing; counters feed `host.Inspector`, `Page.Stats`, and the devtools Frame tab.
- Stage benchmarks at 640x480: `css.Apply` is ~200 µs for 200 rules, `css.Relayout` reuses collected sheets in ~1.2 µs; `BenchmarkRedrawHeavyWarm` fell from 1.34–1.83 ms to 489–708 µs.
- Engine (`gowkhtmltopdf`): `css.Relayout(ctx, doc, w, h, focus, hover, active)` plus a sheet cache on branch `chore/changes-for-go-gpui`, with parity, leak, gate, reuse, and import tests and a benchmark. Pushed at `1a39183`; `go.mod` pins it and no longer replaces `../gowkhtmltopdf`.
- `TestResizeReusesParseAndCascade` pins 1 parse, 1 cascade, 10 redraws/layouts/repaints.

### Dynamic resize

- `MaxWidth`/`MaxHeight` are window bounds now, enforced through `ebiten.SetWindowSizeLimits`; zero or negative means no cap (the old 2560 default is gone).
- A drag commits one relayout per settled size behind a 100 ms motion throttle; skipped frames draw the previous artifact scaled, and the final size always commits.
- Hover and active re-resolve after a relayout; scroll clamps after relayouts and navigation; media queries, `vw` units, and percentages follow the window.
- Example: `go run ./examples/resize` (web on 8127).

### Incremental repaint

- `Page.Invalidate` marks the box, its children, and its previous rect; `Redraw` diffs old and new display lists, and `TakeDirty` hands the window a padded region.
- `replay.DrawRect` replays only the operations that meet the region; the window keeps a content-sized persistent buffer and blits it.
- A diff over 8 changed ops or one third of the frame falls back to the full frame. The bitmap fallback and ticking pages still repaint in full; `Invalidate()` with no argument is the escape hatch for pages that fetch during layout or run a clock.
- Byte-equality tests hold interactive pages (login, platform, states, forms, scroll, theme) to the same PNG as a full `Redraw`.

### DevTools

- F12 or Ctrl+Shift+I toggles a full-height right dock. The overlay never enters `Page.PNG`, the display list, or the box list.
- Tabs: Elements (picked box as pretty JSON with fold), Frame (window, rendering, pipeline, and reload counters), Ops (paint-order list with per-kind colours; clicking an op outlines it at its painted bounds).
- API: `Page.Stats`, `host.Inspector`, `Config.DevTools`, `Page.SetDevTools`; `GET /debug/state` in web mode; `go run ./examples/devtools` (web on 8122).
- Limit: the engine publishes no per-element style reader, so the panel shows geometry and attributes, not computed styles.

### Hot reload

- `Config.File` and `Config.ThemeFile` read from disk and are watched; `Config.DisableHotReload` and `Page.SetHotReload` turn the 250 ms poll off.
- A broken edit keeps the last good picture and retries on the next poll. A reload replaces the history entry in place, keeps history length, typed values, focus, hover/active, images, and tick, and fires no `Change` callback.
- `Serve` reloads before answering and the shell refreshes `/frame.png`; `go run ./examples/reload` (web on 8123).

### Drag and drop

- `host.Dropper`, the window drop pass over `ebiten.DroppedFiles()`, and `Handlers.Drop`; `Drop.Read` is valid only during the call. A directory stays one entry. `-web` receives no drops; the wasm canvas does.
- Example: `go run ./examples/drop` (web on 8124).

### Printing

- `Page.PDF`/`WritePDF`/`SavePDF`/`Print`; `internal/print` uses `lp` or `xdg-open`, `osascript` or `open`, and PowerShell per OS; wasm and mobile return `ErrNoPrinter`; `GPUI_PRINT_DEBUG=1` logs fallbacks; `Serve` adds `GET /pdf`.
- The PDF is a re-render from source. The live window cannot print its display list, `SetImage` entries are absent, and pagination can differ from the screen.
- Example: `go run ./examples/print` (web on 8125).

### Packaging

- `scripts/package.sh [-n] <example>` writes one release archive per system (Linux tar.gz with a `.desktop` file, macOS `.app` zip, Windows `.exe` zip, wasm zip) with `-trimpath` and stripped flags, plus `SHA256SUMS`. The dry run builds nothing.
- Guide: `documentation/packaging.md`.

### Input and interaction

- Caret and selection range on a rune-offset model; click-to-offset through `internal/textrun`; drag, double-click, and triple-click selection with edge auto-scroll; keys and chords edit the range.
- Tab and Shift+Tab traverse focus through `host.Focuser`; Escape closes the context menu, then clears focus.
- Context menu through `host.ContextMenu` (cut, copy, paste, select all, undo, redo), cursor shapes through `host.CursorShape` (I-beam, hand, resize), programmatic scroll through `host.ScrollRequester`, touch-drag scrolling, pinch zoom (0.25–4), and F11 fullscreen.
- The caret is placed before the click handler runs (`5bcdf97`). Example: `go run ./examples/input` (web on 8126).
- Out: textarea Up/Down, selection past a field's visible scroll, and IME (Ebiten v2.10.4 `exp/textinput` is the route; next cycle).

### Repository layout

- `examples/` is its own Go module, `github.com/chinmay-sawant/go-gpui/examples`; a committed `go.work` joins it with the root module and pins the root at `v0.0.2` while that tag is unpublished.
- The library module zip drops the examples: 626 files and ~1.5 MB instead of ~50 MB, most of which was the `desktop-cat` PNG set. `examples/go.mod` has no `replace`, so versioned `go install` stays clean.
- `go.mod`, `examples/go.mod`, and `go.work` require Go 1.26.4.

### Late work after the report pages

- DevTools right-dock commits (`c5ff5a6`..`140dc48`) and the op-outline fix (`067ba1a`).
- Examples: `spotify-player` expanded to eight screens and `audio-player` removed; the `wispr-flow-dashboard` app shell with one package per screen (166 files); the `desktop-cat` transparent companion with native media notifications (75 files); a clipboard Paste button; the flappy bird game moved to template files.
- Page and replay fixes: zero max means no cap, a click resolves through an anonymous child box, diagonal display lines stroke as segments, the replay buffer sizes to page content, and rewritten fields emit `data-gpui-placeholder`.

## Impact

| Area | Impact |
|---|---|
| **Performance** | A drag or hover relayouts instead of reparsing; a click repaints the dirty box; `BenchmarkRedrawHeavyWarm` 1.34–1.83 ms → 489–708 µs. |
| **Memory** | One cached tree and styled document per parse; one content-sized replay buffer in the window. |
| **Behavior / correctness** | Partial repaints are byte-equal to full redraws in tests; fallback and ticking pages stay full repaint. |
| **API / CLI** | 29 new `Page` methods, 4 new `Config` fields, `Handlers.Drop`, optional `host` interfaces, and `WindowOptions` for the desktop-cat example. `host.Screen` is unchanged. |
| **Dependencies** | `go.mod` pins the pushed engine commit; no new direct dependency. The examples module carries the audio indirects, and the library module drops `go-mp3` and no longer ships the examples in its zip. |
| **Binary size / build time** | Not measured in this PR; desktop and `GOOS=js GOARCH=wasm` builds compile. |

## Breaking changes / migration

| Item | Migration |
|---|---|
| `MaxWidth`/`MaxHeight` bound the window now, and zero means no cap | Pass explicit values for a cap; drop the old 2560 literals. |
| `examples/audio-player` removed | Use `examples/spotify-player`, which absorbed its screens. |
| Engine pin moved to `chore/changes-for-go-gpui` `1a39183` | `go mod download`; no sibling checkout needed. |
| New optional `host` interfaces (`Dropper`, `Focuser`, `ContextMenu`, `CursorShape`, `ScrollRequester`, `Reloader`, `Inspector`) | Add a method to opt in; `host.Screen` is unchanged. |

## Test plan

- [x] `make test` (`go test -p 1 ./...`) — 381 test files.
- [x] `make build` (`go vet -p 1 ./...`) — passes with the pinned engine and no replace.
- [x] `GOOS=js GOARCH=wasm` compile check.
- [x] Headless Chrome screenshots for the DevTools dock and the merged report pages.
- [ ] CI is not configured in this repo.

### Commands

```sh
make build
make test
go run ./examples/resize
go run ./examples/devtools
go run ./examples/input
```

## Screenshots / sample output

- `examples/resize`: drag the window edge; the columns switch, the `100vw` bar follows, and hover survives the relayout.
- `examples/devtools`: F12 toggles the dock; hover shows JSON, click pins a box, an op row outlines the operation.
- `examples/input`: click places the caret, drag selects, Tab traverses, right-click opens the menu, touch scrolls and pinches.
- `plans/v0.0.2/report-html-2.0.0/v0.0.2-report.html`: the six report pages with file and line citations (not part of the branch).

## Related issues

None. No tickets exist in this repo; `plans/v0.0.2/` is the reference.

## PR metadata checklist (author)

- [ ] Self-assigned with `--assignee "@me"`.
- [ ] Labels applied (`enhancement`, `documentation`).
- [ ] Body copy kept at `plans/v0.0.2/pr/pr-v0.0.2.md`.
- [ ] Ticket IDs linked: not applicable, no issues exist.

Suggested open command (not run):

```sh
gh pr create \
  --base master \
  --head feature/v0.0.2 \
  --title "feat: ship go-gpui v0.0.2" \
  --body-file plans/v0.0.2/pr/pr-v0.0.2.md \
  --assignee "@me" \
  --label enhancement \
  --label documentation
```

## Follow-ups (out of scope)

- IME: wire Ebiten `exp/textinput` and a page-side composing run.
- Computed styles in the DevTools panel (needs an engine per-element style reader).
- DevTools touch support, the dock viewport-shrink decision, and a golden dock screenshot test.
- `@media (height)`, `(min-height)`, and `(max-height)` matching in the engine.
- Partial repaint for the bitmap fallback (needs a new engine call).
- Canvas, accessibility, video, WebGL, and auto-update (reasons in `plans/v0.0.2/electron-gaps.md`).
- Textarea Up/Down and selection beyond a field's visible scroll.

## Reviewer checklist

- [ ] Behavior matches the summary and test plan.
- [ ] No unrelated changes in the diff: the branch covers the v0.0.2 feature set and the examples that shipped with it.
- [ ] Public API changes documented in `documentation/`.
- [ ] PR has assignee and labels.
- [ ] No secrets or generated artifacts committed.
- [ ] Diff-stat-by-extension table pasted at the bottom.

## Diff stat by extension

Generated with `bash scripts/pr-diff-stat.sh master` after the final commit (`v0.0.1` predates this range).

| Extension | Files | Insertions | Deletions |
| --- | ---: | ---: | ---: |
| `.css` | 36 | 1527 | 499 |
| `.go` | 562 | 22444 | 2805 |
| `.html` | 52 | 3740 | 133 |
| `.json` | 1 | 255 | 0 |
| `.jsonl` | 1 | 28 | 0 |
| `.md` | 48 | 3497 | 333 |
| `.mod` | 2 | 47 | 3 |
| `.png` | 30 | Binary | Binary |
| `.sh` | 4 | 279 | 6 |
| `.sum` | 2 | 116 | 6 |
| `.svg` | 76 | 265 | 150 |
| `.work` | 1 | 8 | 0 |
| No extension | 3 | 25 | 4 |
| **Total** | **818** | **32231** | **3939** |
