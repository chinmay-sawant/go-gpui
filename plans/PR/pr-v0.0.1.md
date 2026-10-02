## Summary

Ship go-gpui v0.0.1: one Go process that shows an HTML template in an Ebiten window with Go handlers for clicks and typing. The page is laid out by `gowkhtmltopdf`; the window replays the layout as vector operations, or draws the engine's bitmap when an operation cannot be replayed. Forms, clipboard, in-process IPC, HTML history, fetch, and local crash reports ship in the same branch.

## Motivation / context

- Plan: [`plans/v0.0.1/compare.md`](plans/v0.0.1/compare.md), the Electron gap list that defines this version.
- This branch is the complete v0.0.1 workstream: 18 commits, from the first screen-paint commit to the display-list replay.
- No issue is linked; the repo had no issues or PRs before this one.

## Changes

### Rendering pipeline

- `internal/render`: `Paint` keeps the bitmap path (`html.Parse`, `css.Apply`, `layout.Lay`); `DisplayList` builds the same placement without rasterizing; `Replayable` and `FillRadii` decide whether the replay can draw every operation. No PDF is written anywhere.
- `internal/replay` (new): draws the display list in paint order. Fills with circular corners, axis-aligned border lines via `PaintLineGeometry`, collapsed table grids, and shaped text through Ebiten `text/v2` from `op.Font.Bytes()` (baseline placed by face ascent, `text-transform` applied, fake bold as a one-pixel double strike). Colors and alpha come from `Opacity()`.
- `internal/window`: `syncImage` picks the display list or one Ebiten image per generation and disposes the replaced texture; `Draw` replays with the scroll offset or blits; `pointer` sizes hit testing and scrolling from `Display.Width/Height`; `badge` paints a `bitmap fallback` notice in the top-right corner when the bitmap path is active.
- `internal/page`: `Redraw` stores the `layout.Display` and its `Boxes` for replayable pages (and no bitmap), or the `image.Image` and `Lay` boxes for the fallback. `PNG` rasterizes a replayed page on demand, once, and caches until the next redraw. `Prepare` checks `Image` and `Display` instead of encoding a PNG to test drawnness.

### Forms and input

- Form controls (input, select, checkbox, radio, file, textarea) are scanned into live state, rewritten into the HTML before each redraw, and edited through Type, Backspace, DeleteWord, Submit, Copy, Cut, Paste, SelectAll, Undo, and Redo.

### Runtime features

- `internal/ipc`: in-process `Send`, `Listen`, `Handle`, `Request`, re-exported by `ipc.go`.
- `internal/page` navigation: `Load`, `Back`, `Forward`, `Route`, and history, with `ErrNoHistory`.
- `internal/fetch`: `Fetch` and `XHR` for http and https only, no cookies or cache, `ErrScheme` for other schemes.
- `internal/clipboard`: memory copy plus OS integration (X11 selection on Linux, Unicode text on Windows, NSPasteboard on Darwin with cgo), with `UseMemory` for tests.
- `internal/crash`: local report files, wired into `Run` and `BindMobile`.

### Examples, build scripts, docs

- `examples/login` (sign-in, secret/secret) and `examples/forms` (every control), plus `browser/index.html` and `scripts/browser.sh` for the wasm build.
- `documentation/` feature guides, `RENDERING-FINDINGS.md`, and `rendering-flow.html`, a self-contained diagram of the draw path.

## Impact

| Area | Effect |
|---|---|
| Performance | Replayable pages skip whole-page rasterization per redraw; frames draw vector paths and cached glyphs. |
| Memory | A replayable page holds no RGBA page buffer; the window disposes the Ebiten texture it replaces. |
| Behavior/correctness | A page with any unsupported operation keeps the engine bitmap path, so mixed pages never render half vector and half bitmap. |
| API/CLI | `host.Screen` gains `Display() *layout.Display`. `Page` gains `Display()`; `Page.Image()` returns nil on a replayed page. `Run`, `Serve`, `BindMobile`, handlers, and examples keep their shapes. |
| Dependencies | Ebiten v2.9.8 to v2.10.4, `golang.org/x/image` and `golang.org/x/text` direct. `go.mod` carries a temporary local `replace` for `Display.Boxes`. |
| Binary size/build time | Not measured in this PR; desktop and `GOOS=js GOARCH=wasm` builds compile. |

## Breaking changes / migration

| Item | Migration |
|---|---|
| `host.Screen` gains `Display() *layout.Display` | Add the method to a custom screen, or keep using `*gpui.Page`, which implements it. |
| `Page.Image()` is nil after redrawing a replayable page | Read `Page.Display()` for operations, or `Page.PNG()` when bytes are needed. |
| `go.mod` replaces gowkhtmltopdf with `../gowkhtmltopdf` | Builds need the sibling checkout while `Display.Boxes` is not in a published commit; drop the replace once the upstream branch merges. |

## Test plan

- [x] `go test ./...`
- [x] `go vet ./...`
- [x] `GOOS=js GOARCH=wasm go build ./...`
- [x] Desktop capture: `examples/login` replays with no badge; a rounded-border page shows the `bitmap fallback` badge in the top-right corner.
- [ ] CI is not configured in this repo.

### Commands

```sh
go test ./...
go vet ./...
GOOS=js GOARCH=wasm go build ./...
```

## Screenshots / sample output

Desktop captures taken while building this branch (not committed):
- Replay path: `examples/login`, vector text and borders, no badge.
- Fallback path: a rounded-border page, bitmap output with the `bitmap fallback` badge.

## Related issues

None. No tickets exist in this repo; `plans/v0.0.1/compare.md` is the reference.

## PR metadata checklist (author)

- [x] Self-assigned with `--assignee "@me"`.
- [x] Labels applied (`enhancement`, `documentation`).
- [x] Body copy committed at `plans/PR/pr-v0.0.1.md`.
- [ ] Ticket IDs linked: not applicable, no issues exist.

## Follow-ups (out of scope)

- Push and merge the gowkhtmltopdf branch `chore/changes-for-go-gpui`, then drop the `replace` and bump the pin.
- Move more operations into `internal/replay`: stroke rects, images, non-identity transforms, blend groups, letter-spacing, and rotation.
- Avoid the extra layout on fallback pages (`DisplayList` then `Paint`).

## Reviewer checklist

- [ ] No unrelated changes in the diff: the branch covers the v0.0.1 feature set only.
- [ ] No secrets or generated artifacts committed.
- [ ] `go test ./...` passes with the sibling gowkhtmltopdf checkout present.
- [ ] The local replace is understood as temporary.

## Diff stat by extension

| Extension | Files | Insertions | Deletions |
|-----------|-------|------------|-----------|
| `.go` | 122 | 5834 | 180 |
| `.md` | 13 | 773 | 11 |
| `.html` | 2 | 409 | 0 |
| `.sum` | 1 | 22 | 22 |
| `.mod` | 1 | 11 | 10 |
| **Total** | 139 | 7049 | 223 |
