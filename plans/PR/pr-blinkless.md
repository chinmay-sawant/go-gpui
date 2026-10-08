# feat: lay out pages with blinkless

Filled copy of [`skills/PR/PR_TEMPLATE.md`](../../skills/PR/PR_TEMPLATE.md).

- Base: `origin/master` (`ba937db`)
- Head: `feat/blinkless`
- Diff against `origin/master`: 150 files, including this plan file

---

## Summary

ownframe now lays pages out with blinkless `v0.0.0-20261008162417-1078f00d42d9`. The library and the examples module both require that commit. Replayable pages keep the display list. Other pages paint that list into a local bitmap so PNG and fallback frames still encode outside an Ebiten game. PDF calls return `ErrNoPDF` because blinkless does not write a PDF.

---

## Motivation / context

- The layout engine moved from `gowkhtmltopdf` to blinkless. The pin is the published pseudo-version, with no `replace`.
- `v0.0.2` is not on the remote. Examples require `github.com/chinmay-sawant/ownframe v0.0.0` and replace it with the parent checkout.
- Issues: none. This work has no tracker issue.

---

## Changes

### Engine

- Both `go.mod` files require `github.com/chinmay-sawant/blinkless v0.0.0-20261008162417-1078f00d42d9`.
- Go imports use `blinkless` and its `css`, `html`, and `layout` packages.
- `layout.DisplayList` is the placement call. `css.Relayout` still re-places a styled document.

### Bitmap and PDF

- `internal/bitmap` paints fills, images, and text into an `image.NRGBA` for PNG and for pages the window cannot replay.
- The window still replays a vector page onto the Ebiten canvas.
- `PDF`, `WritePDF`, `SavePDF`, and `Print` return `ErrNoPDF`.

### Module graph

- `examples/go.mod` replaces `github.com/chinmay-sawant/ownframe` with `../`.
- `go.work` no longer pins `v0.0.2`.

---

## Impact

| Area | Impact |
|------|--------|
| **Performance** | The window replay path is unchanged. PNG and fallback frames paint on the CPU. |
| **Memory** | A non-replayable page keeps one NRGBA the size of the display. |
| **Behavior / correctness** | Theme, images, text, and translated pinned bars still show in PNG tests. Letter-spacing, strokes, and non-translation image transforms are not in the CPU painter. |
| **API / CLI** | `ErrNoPDF` is the PDF result. `PDFOptions` fields stay on the type and are unused. |
| **Dependencies** | Direct engine requirement is blinkless at the commit above. |
| **Binary size / build time** | One module swap. No new toolchain. |

---

## Breaking changes / migration

| Item | Migration |
|------|-----------|
| Engine module | Require blinkless `v0.0.0-20261008162417-1078f00d42d9`. Drop `gowkhtmltopdf`. |
| PDF | Callers of `PDF`, `WritePDF`, `SavePDF`, and `Print` receive `ErrNoPDF`. |
| Examples module | The workspace no longer supplies `v0.0.2`. The examples `replace` points at the parent checkout. |

---

## Test plan

- [x] `make test`
- [x] `go test -count=0 -p 1 . ./examples/login/ ./internal/page/`
- [ ] `make lint` / `go vet`
- [ ] `make build` (when binary output is part of the change)
- [ ] `make run` wall time vs baseline (hard < 400ms; soft ±50ms of reference)
- [ ] `make reference-metrics` / gopdfsuit hard metrics if detector surface changed

### Commands

```sh
make test TEST_P=1
go test -count=0 -p 1 . ./examples/login/ ./internal/page/
```

`make test TEST_P=1` passed on this tree. The count=0 compile of the root package, `internal/page`, and `examples/login` passed before that.

---

## Screenshots / sample output

```
(no window run for this engine swap)
```

---

## Related issues

- None. No tracker issue exists for this engine switch.

---

## PR metadata checklist (author)

- [x] Self-assigned (`--assignee @me`)
- [x] Labels applied
- [ ] Related issues filled with real ticket IDs
- [x] Filled body committed under `plans/PR/pr-blinkless.md`

---

## Follow-ups (out of scope)

- Publish an ownframe release tag and drop the examples `replace`.
- CPU bitmap letter-spacing, strokes, and non-translation image transforms.

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
| `.go` | 139 | 540 | 338 |
| `.md` | 6 | 158 | 17 |
| `.mod` | 2 | 6 | 3 |
| `.sum` | 2 | 4 | 4 |
| `.work` | 1 | 0 | 2 |
| **Total** | **150** | **708** | **364** |
