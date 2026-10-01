# Compare with Electron

Recorded for v0.0.1 from a scan of this repo on 2026-10-01, at commit 980f021. `skills/` was ignored. This is the gap list, plus the screen-paint fact that the list is easy to misread.

## What this repo is

`gpui.New` parses one `html/template` string. `Redraw` fills that template, then calls `gowkhtmltopdf` `html.Parse`, `css.Apply`, and `layout.Lay`. The desktop window, the wasm canvas, and the phone bind all draw `layout.Result.Image()` through Ebiten. `Page.PNG` encodes that same image for `GET /frame.png` and for tests. There is no DOM, no JavaScript, and no second process.

`go.mod` directly requires `github.com/chinmay-sawant/gowkhtmltopdf` and `github.com/hajimehoshi/ebiten/v2`. The new work on this branch does not add modules.

## The screen is not HTML, then PDF, then an image

`internal/page/page_draw.go` never calls `Document.WritePDF`, `Document.PDF`, or `ImageDocument`. `layout.Lay` calls `imageout.RenderLayout` and stores the `image.Image` it returns.

The module is still named gowkhtmltopdf, and its image painter uses `pdf.Font` and `pdf.Registry` for faces. Those are font tables inside the library. They are not a PDF file that this window rasterizes.

`PNG()` runs only after that image exists. The `-web` host needs those bytes because a browser cannot hold the Go `image.Image`. The native window does not.

Painting stays on `html.Parse`, `css.Apply`, and `layout.Lay`. go-gpui owns the call. It does not grow its own layout engine, and it does not start calling the PDF writer to feed the window.

## What already works

- Public calls are `New`, `SetData`, `Handle`, `Run`, `Serve`, and `BindMobile`. Input is a fixed set of callbacks. `Click`, `Type`, `Backspace`, `DeleteWord`, `Submit`, `Copy`, `Cut`, `Paste`, `SelectAll`, `Undo`, `Redo`. A nil callback skips that input. `Copy` does not redraw.
- One decorated, resizable Ebiten window. Minimum frame 320 by 400 in the login example. Wheel scroll is 48 CSS pixels when the picture is larger than the window. A picture that matches the window stays 1:1.
- Text chords are Ctrl or Cmd with C, V, X, A, Z, Y, plus Shift-Insert, Ctrl-Insert, Shift-Delete, and Ctrl-Backspace. AltGr still types.
- Clipboard write and read try `wl-copy` and `wl-paste` when `WAYLAND_DISPLAY` is set, else `xclip` when `DISPLAY` is set. Both paths keep an in-memory copy. That is not a native clipboard API.
- Hosts are the Ebiten window, the same loop as wasm (`scripts/browser.sh`, port 8092), Ebitengine mobile bind, and `-web` on 127.0.0.1:8091. `-web` only serves the picture, click, type, and backspace.
- The login example owns field text. The library does not implement inputs. The button is `data-action=login`.

## Pending against Electron

- No Chromium document, V8, preload, `contextBridge`, or Node.
- No IPC. Electron IPC is how the Node main process and the page process pass messages. This program is one Go process, and the only bridge today is those input callbacks.
- No navigation. `New` takes one HTML string. No `loadURL`, no document `<a href>`, no back or forward, no custom protocol. `examples/login/login/login_history.go` is an undo stack for the password field. `-web` `href` values are image-map clicks, not page loads.
- No multi-window, BrowserView, frameless window, always-on-top, fullscreen, native menu, tray, notifications, or native file dialogs.
- No session, cookies, cache, `fetch`, XHR, or web storage. The only HTTP is the PNG host and the wasm file load.
- No DevTools, auto-update, or crash reporter. Packaging is the wasm serve script and the documented `ebitenmobile` bind. There is no installer.
- Forms are single-line text painted as `div`s, handled in Go. No `input`, select, checkbox, radio, file input, or textarea.
- No video, audio, document canvas, WebGL, `contenteditable`, file drag-and-drop, or context menu.
- No IME. Typed text is `ebiten.AppendInputChars`.
- No accessibility tree, spellcheck, printing, multi-monitor placement, or deep links.
- Shortcuts stay inside the window, and they only edit text. There are no OS-global shortcuts.
- No sandbox, CSP, or context isolation. `-web` does not authenticate.
- Clipboard is `wl-copy` / `xclip` plus memory, not the platform clipboard API. Windows, macOS, Android, iOS, and wasm have no OS clipboard path.

## In progress on feature/v0.0.1

Separate worktrees, then merged back here. No new third-party modules.

- Screen paint stays a direct `layout.Lay` image, called from this repo, still not a PDF round-trip.
- In-process IPC, `Send` / `Listen` / `Request`, so app code can pass messages without a Node process.
- `Load`, back, forward, and `data-action` routes.
- A local crash report file. Nothing is uploaded.
- `Fetch` and `XHR` over `net/http`, http and https only.
- Clipboard through the OS API on Linux, Windows, and macOS. Memory remains the fallback for tests and for targets with no desktop clipboard.
