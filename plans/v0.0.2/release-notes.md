## v0.0.2

Second release of **ownframe**: one Go process that shows an HTML template in a window and calls Go functions for clicks, keys, and typing. This release keeps the parsed and styled document between frames, repaints only the dirty box, and adds a DevTools dock, hot reload, drag and drop, printing, packaging, and a full input layer. There is still no JavaScript engine and no second process. `gowkhtmltopdf` parses the HTML, applies the CSS, and lays the page out; a page whose operations the replay accepts keeps vector operations and no bitmap, and any other page falls back to the engine image.

- Module: `github.com/chinmay-sawant/ownframe`
- Go: 1.26.4
- Ebiten: v2.10.4, the floor for text replay
- Engine: `gowkhtmltopdf` pinned at `v0.2.7-0.20261004151708-1a3918301a68` (`chore/changes-for-go-gpui`), no local replace
- Direct requires: `gowkhtmltopdf`, Ebiten, `golang.org/x/image`, `golang.org/x/text`
- Docs: `documentation/` (29 files)
- Examples: 31 runnable examples plus shared `examples/music` support
- Review: `plans/v0.0.2/report-html-2.0.0/v0.0.2-report.html` (six parts, one file)
- Previous release: v0.0.1

### Highlights

| Area | What you get in v0.0.2 |
|------|------------------------|
| Pipeline split | `internal/render` caches the parsed tree and the styled document; a size or state change calls `css.Relayout` instead of reparsing. |
| Dynamic resize | `MaxWidth`/`MaxHeight` bound the window; a drag commits one relayout per settled size (100 ms throttle); hover, active, and scroll re-resolve after a relayout. |
| Incremental repaint | A click, hover, or edit repaints only the dirty region through `TakeDirty` and `replay.DrawRect`; byte-equality tests hold it to the full-redraw PNG. |
| DevTools | F12 or Ctrl+Shift+I right dock with Elements JSON, Frame counters, and Ops outlines; `GET /debug/state`. |
| Hot reload | `Config.File` and `Config.ThemeFile` watch and swap the page in 250 ms, keeping state and history. |
| Drag and drop | `host.Dropper` and `Handlers.Drop` receive the dropped files of one frame. |
| Printing | `Page.PDF`/`WritePDF`/`SavePDF`/`Print` and the per-OS print path; `GET /pdf` in web mode. |
| Packaging | `scripts/package.sh` writes one release archive per system with `SHA256SUMS`. |
| Input | Caret and selection, click-to-offset, Tab traversal, context menu, cursor shapes, `ScrollTo`/`ScrollBy`, touch scroll, pinch zoom, F11. |
| Engine | `css.Relayout` and a sheet cache on the engine branch, with parity, leak, and reuse tests. |

### Pipeline split and the engine

`Page.Redraw` still executes the template, but the executed bytes now key a cache. `internal/render.Cache` holds the `*html.Document` and the styled `*css.Document` from one parse and one cascade. An unchanged source calls `css.Relayout` and lays the cached document out again; a changed source builds a new cache. `SetSize`, hover, press, release, focus, and theme changes all follow the relayout path, so a drag or a hover reparses nothing.

At 640x480, the stage benchmark put `css.Apply` at ~200 µs for 200 rules and `css.Relayout` at ~1.2 µs by reusing collected sheets. `BenchmarkRedrawHeavyWarm` fell from 1.34–1.83 ms to 489–708 µs. `Page.Stats` reports parses, cascades, layouts, repaints, and relayouts; the devtools Frame tab and `GET /debug/state` show the same counters.

The engine side landed on `gowkhtmltopdf` branch `chore/changes-for-go-gpui` at `1a39183`: the `css.Relayout(ctx, doc, width, height, focus, hover, active)` entry point, the sheet cache, parity and leak tests, and a relayout benchmark. `go.mod` pins that commit as `v0.2.7-0.20261004151708-1a3918301a68` and no longer replaces `../gowkhtmltopdf`.

### Dynamic resize

`MaxWidth` and `MaxHeight` are window bounds, enforced through `ebiten.SetWindowSizeLimits`; zero or negative means no cap, so the old 2560 default is gone. A drag commits one relayout per settled size behind a 100 ms motion throttle, and skipped frames draw the previous artifact scaled. The final drag size always commits. Media queries, `vw` units, and percentages follow the window, hover and press re-resolve from the cursor after a relayout, and the scroll position clamps after a relayout or navigation. `go run ./examples/resize` shows the columns switching, the `100vw` bar following, and a hover control.

### Incremental repaint

`Page.Invalidate` marks a box, its children, and its previous rectangle. `Redraw` diffs the old and new display lists, and `TakeDirty` hands the window a padded region. `replay.DrawRect` replays only the operations that meet that region, and the window keeps a content-sized persistent buffer and blits it. A diff over 8 changed ops or one third of the frame falls back to a full frame; the bitmap fallback and ticking pages still repaint in full. `Invalidate()` with no argument marks the whole frame for pages that fetch during layout or run a clock. Tests byte-compare an interactive page against a full `Redraw` for login, platform, states, forms, scroll, and theme.

### DevTools

F12 or Ctrl+Shift+I toggles a full-height right dock. The Elements tab shows the hovered or pinned box as pretty JSON with fold and unfold. The Frame tab groups right-aligned counters under WINDOW, RENDERING, PIPELINE, and RELOAD. The Ops tab lists every operation in paint order with a colour per kind; clicking a row outlines the operation at its painted bounds. `Page.Stats` and `host.Inspector` expose the counters, `Config.DevTools` starts the overlay on, and `Page.SetDevTools` toggles it. The overlay is window chrome: it never enters `Page.PNG`, the display list, or the box list. `go run ./examples/devtools` (F12 toggles; `-web` serves `GET /debug/state`). The engine publishes no per-element style reader, so the panel shows geometry and attributes, not computed styles. The dock handles the mouse only; touch users turn it on programmatically.

### Hot reload

`Config.File` and `Config.ThemeFile` read the template and theme from disk and watch them. The window polls every 250 ms before it syncs the frame. A broken edit keeps the last good picture and retries on the next poll; a deleted file keeps the picture and reports once. A reload replaces the current history entry in place, keeps the history length, typed values, focus, hover and active state, images, and tick, and fires no `Change` callback. `Serve` polls before answering and refreshes `/frame.png`; `-reload=false` and `Config.DisableHotReload` turn the watch off. `go run ./examples/reload` edits a temp copy of `index.html`.

### Drag and drop

`host.Dropper` receives the dropped entries of one frame from `ebiten.DroppedFiles()`; a screen without it ignores files. `Handlers.Drop` handles them on the page, and `Drop.Read` is valid only during the call. A directory stays one entry. `-web` has no window loop and receives no drops; the wasm canvas runs the same pass. `go run ./examples/drop` prints the dropped paths and shows dropped images.

### Printing

`Page.PDF`, `Page.WritePDF`, `Page.SavePDF`, and `Page.Print` render the template to PDF again and hand it to the OS print path: `lp` or `xdg-open` on Linux, `osascript` or `open` on macOS, PowerShell `Start-Process -Verb Print` on Windows. wasm, Android, and iOS return `ErrNoPrinter`; `OWNFRAME_PRINT_DEBUG=1` logs the fallback reasons. `Serve` adds `GET /pdf` for a screen that renders PDF bytes. The live window cannot print its own display list, `SetImage` entries are absent from the PDF, and pagination can differ. `go run ./examples/print`.

### Packaging

`scripts/package.sh [-n] <example>` builds one example into `dist/`: a Linux tar.gz with the binary and a `.desktop` file, a macOS `.app` zip, a Windows `.exe` zip, or a wasm zip. The script uses `-trimpath` and stripped flags and writes `SHA256SUMS`. `-n` prints the names and entries without building. `documentation/packaging.md` has the layouts and the signing caveats, including that macOS archives need a macOS host.

### Input and interaction

- A form field carries a caret and an anchor, both rune offsets; a selection normalizes to `[min, max)` and keys replace it. A click focuses the field and places the caret at the clicked glyph through `internal/textrun`. Drag, double-click, and triple-click extend or select by word and line, and an edge drag auto-scrolls.
- Tab and Shift+Tab move focus through `host.Focuser` in document order, skipping disabled controls and `tabindex < 0`. The window consumes the key only when a field can take focus. Escape closes the context menu first, then clears focus.
- The context menu (`host.ContextMenu`) offers cut, copy, paste, select all, undo, and redo from live page state. The cursor shape (`host.CursorShape`) is an I-beam over text, a hand over a link or button, and a resize arrow over a scrollbar thumb.
- `ScrollTo` and `ScrollBy` queue a request that the window clamps; touch drags scroll with an 8 px slop, and pinch zoom clamps between 0.25 and 4 while clicks stay on the same box. F11 toggles desktop fullscreen.
- Out of scope: textarea Up/Down, selection beyond a field's visible scroll, and IME. Ebiten v2.10.4 ships `exp/textinput`, so IME is window wiring plus a page-side composing run, not an upstream ask. It is the first input item of the next cycle.

### Examples

31 runnable examples plus shared `examples/music` support. New in v0.0.2: `flappy-bird` (8121), `devtools` (8122), `reload` (8123), `drop` (8124), `print` (8125), `input` (8126), `resize` (8127), and `desktop-cat` (8128, transparent window and media notifications). `spotify-player` (8119) grew to eight screens and absorbed `audio-player`, which was removed. `wispr-flow-dashboard` (8117) became a full app shell with one package per screen. `clipboard` (8113) added a Paste button, `platform` (8115) repaints only its counter box, and `dino` (8120) moved its game HTML and CSS into templates. `examples/readme.md` indexes every folder and its `-web` port.

### Documentation

`documentation/` grew to 29 files. New guides: `devtools.md`, `hot-reload.md`, `repaint.md`, `drag-drop.md`, `printing.md`, `packaging.md`, `interaction.md`, `compare-electron.md`, and `compare-rust-gpui.md`. Updated: `features.md`, `features-examples.md`, `forms.md`, `frames.md`, `keys.md`, `platforms.md`, `pointer.md`, `screen.md`, `scrolling.md`, `theming.md`, `web.md`, `window.md`, `editing.md`, and `README.md`. `plans/v0.0.2/report-html-2.0.0/` holds six part pages and the merged `v0.0.2-report.html`.

### Limits

- No JavaScript engine, Chromium, preload, `contextBridge`, or Node. One process.
- No cross-process IPC, native menus, tray, notifications, or more than one window. No window icon.
- No session, cookies, cache, or web storage.
- No auto-update or installer.
- No video, document canvas, or WebGL.
- No library audio API and no `<audio>` element. The examples play through `examples/music`.
- No IME, accessibility tree, spellcheck, deep links, or OS-global shortcuts.
- No sandbox, CSP, or context isolation.
- File dialogs exist on desktop `Run` only; there is no soft keyboard on mobile.
- Web mode has no tick, key events, clipboard, submit, undo, select-all, delete-word, audio, file dialog, or crash recovery.
- The bitmap fallback repaints the whole canvas; a partial bitmap path needs a new engine call.
- `@media (height)`, `(min-height)`, and `(max-height)` never match; the engine accepts width and inline-size only.
- Textarea Up/Down and selection beyond a field's visible scroll stay out; the value model is a single line.
- The DevTools panel shows geometry and attributes, not computed styles.
- Printing re-renders from source: no display list, no `SetImage` entries, and pagination can differ.
- Drag and drop is frame-scoped and `-web` receives no drops.
- Transformed WOFF2 files (Google Fonts) are unsupported; TTF, OTF, and null-transform WOFF2 work.

### Install and run

```
go run ./examples/login
go run ./examples/login -web
sh scripts/browser.sh
```

The library:

```go
page, err := ownframe.New(ownframe.Config{
    Title:  "Hello",
    HTML:   `<h1>{{.Title}}</h1>`,
    Width:  480,
    Height: 640,
})
if err != nil {
    return err
}
page.SetData(struct{ Title string }{"Hello"})
return ownframe.Run(context.Background(), page)
```

### Verification

- `make build` (`go vet -p 1 ./...`) passes with the pinned engine and no replace.
- `make test` (`go test -p 1 ./...`) passes; 381 test files across 100 Go packages.
- Desktop and `GOOS=js GOARCH=wasm` builds compile.
- Partial repaints are byte-equal to full redraws across six pages.
- The DevTools dock was verified with headless Chrome screenshots; the six report pages cite the code and the tests.

### What's changed

- 112 commits since `v0.0.1`, 102 of them non-merge.
- Nine feature branches merged into `feature/v0.0.2`: foundation, resize, repaint page-side, repaint shell-side, hot reload, input window-side, input page-side, print, and devtools; drop landed by fast-forward.
- The pull request from `feature/v0.0.2` to `master` is not open yet; link it here once it exists.
- Full history: https://github.com/chinmay-sawant/ownframe/commits/feature/v0.0.2
