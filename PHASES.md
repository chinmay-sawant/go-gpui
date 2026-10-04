# Phase status

Recorded 2026-10-04.

## v0.0.2 (on feature/v0.0.2)

| Area | What shipped |
|---|---|
| Pipeline split | `internal/render` caches the parsed tree and styled document; `SetSize` and state changes call `css.Relayout`, so a drag or a hover reparses nothing. Page counters feed `host.Inspector`. |
| Dynamic resize | `MaxWidth`/`MaxHeight` are window bounds; a drag commits one relayout per settled size (100 ms motion throttle); hover and active re-resolve after a relayout; the scroll clamp runs on relayout and navigation; `examples/resize`. |
| Incremental repaint | `Page.Invalidate`, the display-list diff, `TakeDirty`, `replay.DrawRect`, and a persistent window buffer; byte-equality tests hold incremental and full repaints to the same PNG. The bitmap fallback keeps a full repaint. |
| Hot reload | `Config.File`/`ThemeFile`, the 250 ms poll, pending retry on a broken file, state survival, `Serve` reload plus shell refresh, `examples/reload`. |
| DevTools | `Page.Stats`, `host.Inspector`, the F12 or Ctrl+Shift+I overlay with box picking, operation outlines, and the stats panel, `GET /debug/state`, `examples/devtools`. The overlay never enters `Page.PNG`, the display list, or the box list. |
| Drag and drop | `host.Dropper`, the window drop pass, `Handlers.Drop`, `examples/drop`. `-web` has no window loop and receives no drops; the wasm canvas does. |
| Printing | `Page.PDF`/`WritePDF`/`SavePDF`/`Print`, `internal/print` per OS, `GET /pdf`, `examples/print`. The live window still cannot print its own display list; the PDF is a re-render from source. |
| Packaging | `scripts/package.sh` with the `-n` dry run, archive layouts and `SHA256SUMS`, `documentation/packaging.md`. |
| Input and interaction | Caret and selection range, click-to-offset through `internal/textrun`, Tab focus traversal, drag/double/triple selection, the context menu, cursor shapes, `ScrollTo`/`ScrollBy`, touch drag and pinch, F11. |
| Engine | `css.Relayout` on gowkhtmltopdf `feature/v002-relayout`, with a sheet cache for repeated collection of one tree. |

## v0.0.2 pending

1. Push gowkhtmltopdf `feature/v002-relayout`, then drop the `go.mod` replace and pin that commit.
2. Open the PR for `feature/v0.0.2`.

## v0.0.2 limits and engine gaps

- IME: Ebiten v2.10.4 ships the experimental `exp/textinput` package, so the work is window wiring plus the page-side composing run, not an upstream ask. The route and the design are recorded in `plans/v0.0.2/input-interaction.md`.
- `@media (height)`, `(min-height)`, and `(max-height)` never match: the engine accepts only width and inline-size names (`internal/css/container.go` in gowkhtmltopdf). `TestMediaHeightFollowsTheFrame` is skipped with this reason.
- The bitmap fallback repaints the whole canvas; a partial bitmap path needs a new engine call and is out of scope for v0.0.2.
- Canvas, accessibility, video, WebGL, and auto-update stay after v0.0.2 with the reasons in `plans/v0.0.2/electron-gaps.md`.
- Up/Down in a textarea and selection beyond the visible scroll of a field stay out; the value model is a single line.
- A page that fetches during layout or runs a clock can change pixels the diff does not know about. `Invalidate()` with no argument is the escape hatch.

## v0.0.1 status (previous)

Recorded 2026-10-03.

- gowkhtmltopdf: `chore/changes-for-go-gpui` at `54a29b6` locally (remote at `7f8164f`).
- go-gpui: `chore/pending` at `614d83a` locally, from `master` `ac2ddc4`; `master` untouched.

## Done (on the local branches)

| Area | What shipped |
|---|---|
| A. Template-faithful form controls | Author attributes kept; `data-gpui-field/focus/selected/placeholder`; caret; placeholder; default stylesheet after `<head>`; login on real inputs. |
| Buttons | Engine UA face; submit-like inputs rewritten to buttons; go-gpui workaround removed. |
| B. Replay coverage | Rounded strokes, images, elliptical and masked strokes, transformed images, letter-spaced text, CSS outlines. |
| C. Binding layer | `data-bind` two-way binding; `Handlers.BeforeEdit` and `Handlers.Change`. |
| A leftovers | `FormSelected`; `Cut` writes through binding, fires Change, and is undoable in the login example. |
| D. Engine gaps | Named colors; `oklch`/`oklab`/`color-mix`/`light-dark`; `min`/`max`/`clamp`, general `calc`, `dvh`/`svh`/`lvh`; `@supports`/`@layer`/`@property`; `data:` fonts and null-transform woff2; `conic-gradient`; basic `clip-path`; input value/placeholder painting; control UA faces. |
| Focus and pointer state | `:focus`/`:focus-visible`/`:hover`/`:active`/`:checked`; `css.Options.Focus/Hover/Active`; `Page.Hover/Press/Release`; window pointer tracking; `host.Screen` additions. |

## Pending — integration

1. Push the engine branch (`54a29b6`); remote is at `7f8164f`.
2. Push `chore/pending` (or open a PR); decide when to merge it to `master`.
3. Pin the engine commit in go-gpui's `go.mod` and drop `replace ../gowkhtmltopdf`.
4. `master` in both repos still points where it was.

## Pending — engine limits

- WOFF2: transformed `glyf/loca` (Google Fonts) unsupported; TTF/OTF and null-transform WOFF2 work.
- `:disabled`, `:target` never match; `:focus-within` not implemented.
- `::placeholder`, `::selection` pseudo-elements never match; go-gpui covers them with `data-gpui-placeholder` and `data-gpui-selected`.
- `:hover`/`:active`/`:focus` match the exact id only: no ancestor hover, no state for elements without ids.
- `@layer`: per-stylesheet order; `!important` does not reverse layers.
- `clip-path`: raster-only (backgrounds/images); `@property` syntax parsed but not enforced; bare inputs have no UA width.

## Pending — replay fallbacks

- Blend/isolation groups, rotated text, fake oblique, font features, autospace.
- Elliptical fills, unknown stroke masks, non-image transforms, images without a decodable payload.

## Optional cleanups

- Drop the span rewrite once the engine exposes value mutation for editing.
- `chore/miscellaneous` is superseded by `chore/pending`.
