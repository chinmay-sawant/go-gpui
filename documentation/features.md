# Features

v0.0.1 keeps one Go process and one HTML template. There is no JavaScript engine and no second process. `go.mod` requires `gowkhtmltopdf` and Ebiten directly, plus `golang.org/x/text`, which is already in Ebiten's tree for language tags. These features did not add modules.

## Screen paint

`Page.Redraw` fills the `html/template`, then `internal/render.DisplayList` parses the HTML and applies the CSS. A page whose operations `render.Replayable` accepts keeps the result as a `layout.Display` and no bitmap; any other page falls back to `internal/render.Paint`, which calls `layout.Lay` and returns an `image.Image`. The window replays the display list with `internal/replay`, or blits the bitmap on a fallback page. `Page.PNG` rasterizes on demand for `GET /frame.png` and for tests, and caches the bytes until the next `Redraw`.

`internal/render.DisplayList` stops before the engine's paint step, so it never rasterizes. `Display.Boxes` carries the same hit-test boxes `Lay` returns, so a replayed page needs no bitmap for clicks.

Detail is in [screen.md](screen.md).

## IPC

`Send`, `Listen`, `Handle`, and `Request` pass strings between callers in this process. An empty channel does nothing. `Request` returns `ErrNoHandler` when nobody is handling that channel. There is no socket and no page-process bridge.

Detail is in [ipc.md](ipc.md).

## Navigation

`Load` replaces the template and pushes history. `Back` and `Forward` move through that history. `Route` maps a `data-action` value to HTML, and a click on that action loads it instead of calling the click handler. `HTML` returns the current template source. There is no `loadURL`, no `<a href>` navigation, and no custom protocol.

Detail is in [navigation.md](navigation.md).

## Crash reports

`Run` and `BindMobile` recover a panic, write a text file, and return an error that includes the file path. `Report` writes the same kind of file without a panic. Nothing is uploaded.

Detail is in [crash.md](crash.md).

## Fetch and XHR

`Fetch` sends GET. `XHR` sends a method, headers, and a body. Only `http` and `https` are accepted. A redirect to any other scheme returns `ErrScheme`. The client stores no cookies and no cache.

Detail is in [fetch.md](fetch.md).

## Clipboard

`Write` and `Read` keep an in-memory copy and also talk to the OS clipboard. Linux with `DISPLAY` uses the X11 `CLIPBOARD` selection. Windows uses `CF_UNICODETEXT`. macOS with cgo uses `NSPasteboard`. Wayland, Android, iOS, wasm, and macOS without cgo stay on the memory copy. Tests call `UseMemory` and do not touch the desktop clipboard. No helper program is started.

Detail is in [clipboard.md](clipboard.md).

## Forms

An `input`, `textarea`, or `select` with an id is stored on the page. A click focuses a text field or a textarea, toggles a checkbox, checks a radio, or cycles a select. Typing edits the focused text field even when the type handler is nil. A file input stores a typed name and does not open a dialog. `SetFormValue` and `SetFormChecked` do not redraw. A `<button>` gets a default face when the author does not style it, and a submit-like `input` is rewritten to a `button`. The login example uses a button.

Detail is in [forms.md](forms.md).

## Still absent

These Electron pieces are not in this branch. The scan that listed them is [../plans/v0.0.1/compare.md](../plans/v0.0.1/compare.md).

- Chromium, V8, preload, `contextBridge`, and Node.
- Cross-process IPC, native menus, tray, notifications, file dialogs, and more than one window. A file input stores a typed name and does not open a dialog.
- Session, cookies, cache, and web storage.
- DevTools, auto-update, installer, and an uploaded crash dump.
- Video, audio, document canvas, WebGL, file drag-and-drop, and a context menu.
- IME, an accessibility tree, spellcheck, printing, deep links, and OS-global shortcuts.
- Sandbox, CSP, and context isolation.
