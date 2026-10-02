# Phase status

Recorded 2026-10-02.

- gowkhtmltopdf: `chore/changes-for-go-gpui` at `7f8164f`, pushed.
- go-gpui: `master` at `ac2ddc4`, pushed (all feature branches merged).

## Done

| Phase | What shipped |
|---|---|
| A. Template-faithful form controls | Author attributes kept; `data-gpui-field/focus/selected/placeholder`; caret; placeholder; default stylesheet after `<head>`; login on real inputs. |
| Buttons | `<button>` gets a face and a hit box; submit-like inputs rewritten to buttons. |
| B. Replay coverage | Rounded strokes (`StrokeMask == 0`, circular radii) and identity-transform images replay on the GPU. |
| C. Binding layer | `data-bind` two-way binding to a pointer struct; `Handlers.Change`. |
| A leftovers | `FormSelected`; `Cut` writes through binding and fires `Change`. |
| D. Engine gaps (gowkhtmltopdf) | Full named colors; `oklch`/`oklab`/`color-mix`/`light-dark`; `min`/`max`/`clamp` and general `calc`; `dvh`/`svh`/`lvh`; `@supports`/`@layer`/`@property`; `data:` fonts and null-transform woff2; `conic-gradient`; basic `clip-path`; input value/placeholder painting; UA faces for button/select/textarea. |

## Pending

| Item | Notes |
|---|---|
| Pin the engine commit | go-gpui still uses `replace ../gowkhtmltopdf`. Pin `7f8164f` and drop the replace once the branch lands upstream. |
| Replay fallbacks | Masked/partial strokes, elliptical radii, transformed images, blend groups, outlines, letter-spacing text. |
| `:focus` pseudo-class | Engine needs a focus-state option; go-gpui passes the focused id. `:hover`/`:active` cannot work statically. |
| Login pre-cut undo | The example still cannot snapshot before `Cut`. |
| Button workaround | go-gpui's default `buttonCSS` face is redundant with the engine UA; the submit-input rewrite stays. |

## Engine-side limits

- woff2: null-transform fonts only.
- `@layer`: per-stylesheet order; `!important` does not reverse layers.
- `clip-path`: raster-only (backgrounds/images).
- Bare inputs have no UA width/height.
- `@property` syntax parsed but not enforced.
- `:hover`, `:active`, `:checked`, `::placeholder`, `::selection` do not match; `:focus` is being added.
