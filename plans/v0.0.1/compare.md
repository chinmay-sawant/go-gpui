# Compare with Electron

Recorded for v0.0.1 from a scan of this repo on 2026-10-01, at commit 980f021. The pending list and the landed list were updated on 2026-10-02 after the feature merges. `skills/` was ignored. This is the gap list, plus the screen-paint fact that the list is easy to misread.

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
- Clipboard write and read keep an in-memory copy. On this branch, X11, Windows, and macOS cgo also call the OS clipboard API. Wayland does not.
- Hosts are the Ebiten window, the same loop as wasm (`scripts/browser.sh`, port 8092), Ebitengine mobile bind, and `-web` on 127.0.0.1:8091. `-web` only serves the picture, click, type, and backspace.
- The login example owns field text. The library does not implement inputs. The button is `data-action=login`.

## Pending against Electron

- No Chromium document, V8, preload, `contextBridge`, or Node.
- IPC on this branch is in-process only. There is still no cross-process bridge and no JavaScript.
- Navigation on this branch is `Load`, `Back`, `Forward`, and `data-action` routes. There is still no `loadURL`, no document `<a href>`, and no custom protocol. `examples/login/login/login_history.go` is an undo stack for the password field.
- No multi-window, BrowserView, frameless window, always-on-top, fullscreen, native menu, tray, notifications, or native file dialogs.
- `Fetch` and `XHR` are one `net/http` call, http and https only. There is still no session, cookies, cache, or web storage.
- Crash reports are a local text file. Nothing is uploaded. There are still no DevTools, no auto-update, and no installer. Packaging is the wasm serve script and the documented `ebitenmobile` bind.
- Forms are single-line text painted as `div`s, handled in Go. No `input`, select, checkbox, radio, file input, or textarea.
- No video, audio, document canvas, WebGL, `contenteditable`, file drag-and-drop, or context menu.
- No IME. Typed text is `ebiten.AppendInputChars`.
- No accessibility tree, spellcheck, printing, multi-monitor placement, or deep links.
- Shortcuts stay inside the window, and they only edit text. There are no OS-global shortcuts.
- No sandbox, CSP, or context isolation. `-web` does not authenticate.
- Wayland has no data-device client, so a Wayland session uses the memory copy. Android, iOS, and wasm do too. X11, Windows, and macOS cgo call the OS clipboard API.

## Landed on feature/v0.0.1

Merged from separate worktrees on 2026-10-02. No new modules. How to call each one is in `documentation/features.md`.

- `internal/render.Paint` is the screen path. It calls `layout.Lay` and does not write a PDF.
- In-process `Send`, `Listen`, `Handle`, and `Request`.
- `Load`, `Back`, `Forward`, `Route`, and `HTML`.
- `Run` and `BindMobile` write a local crash file and return its path.
- `Fetch` and `XHR` over `net/http`.
- Clipboard `Write` and `Read` call X11, Win32, or `NSPasteboard`. Wayland stays on the memory copy. Tests call `UseMemory`.
