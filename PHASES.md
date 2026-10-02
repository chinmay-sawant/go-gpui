# Phase status

Recorded 2026-10-02. Latest commit: `3072785` ("forms: let templates style real inputs"),
pushed to `origin/master`. Working tree clean.

## Phase status

| Phase | Status | What it covers |
|---|---|---|
| **A. Template-faithful form controls** | Done, pushed | Real inputs keep author CSS; `data-gpui-field/focus/selected/placeholder`; caret; placeholder text; defaults as a `<style>` after `<head>`; login migrated to real `<input>`s; docs updated. `gofmt` clean, `go test ./...` passes, all Go files under 2000 chars. |
| **B. GPU replay coverage** | Pending | Add `OpStrokeRect` (rounded borders) and `OpImage` (gradients/background images) to `internal/replay`. Today any rounded border or gradient forces the whole page to the 1x bitmap fallback + badge. |
| **C. Binding/directive layer** | Pending (decision needed) | Two-way binding between template data and form values, per-element events, directives. Only needed if "drop in a template, get behavior automatically" goes beyond `html/template` + handlers. |
| **D. Upstream engine CSS gaps** | Pending (decision needed) | In `gowkhtmltopdf`: `:focus`/`:hover`/`:checked` matching, `::placeholder`, `oklch`/`color-mix`/`clamp`, `@supports`/`@layer`, `min()/max()`, woff2 fonts, button/select UA styling, native input-value painting. |

## Small leftovers inside Phase A (optional)

- Selection state is internal; there is no `FormSelection`-style accessor for apps.
- The span rewrite still exists because the engine does not paint input values natively; fixing that upstream would remove it.
- Field values are not interpolated from template data (documented); apps read `FormValue`.
- Login undo/redo cannot capture **Cut**, because the library does not call the app handler when a field is focused. Snapshot-based undo covers typing/backspace/delete-word/paste only.

## Open decisions before B/C/D

1. Is the bitmap fallback acceptable for design-heavy pages, or is B a hard requirement?
2. Do you want C at all, or is `html/template` + handlers the intended contract?
3. Are we allowed to patch `gowkhtmltopdf` for D?
