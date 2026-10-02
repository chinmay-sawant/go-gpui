# Phase status

Recorded 2026-10-03.

- gowkhtmltopdf: `chore/changes-for-go-gpui` at `54a29b6` locally (remote at `7f8164f`).
- go-gpui: `chore/pending` at `614d83a` locally, from `master` `ac2ddc4`; `master` untouched.

## Done (on the local branches)

| Area | What shipped |
|---|---|
| A. Template-faithful form controls | Author attributes kept; `data-gpui-field/focus/selected/placeholder`; caret; placeholder; default stylesheet after `<head>`; login on real inputs. |
| Buttons | Engine UA face; submit-like inputs rewritten to buttons; go-gpui workaround removed. |
| B. Replay coverage | Rounded strokes, images, elliptical and masked strokes, transformed images, letter-spaced text. |
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

- Blend/isolation groups, outlines, rotated text, fake oblique, font features, autospace.
- Elliptical fills, unknown stroke masks, non-image transforms, images without a decodable payload.

## Optional cleanups

- Drop the span rewrite once the engine exposes value mutation for editing.
- `chore/miscellaneous` is superseded by `chore/pending`.
