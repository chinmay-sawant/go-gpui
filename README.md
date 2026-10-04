# go-gpui

**Write the screen in HTML. Write the logic in Go. Ship one binary to the desktop, the browser, and the phone.**

No Chromium, no WebKit, no wrapper around either. The HTML and CSS work is done by [gowkhtmltopdf](https://github.com/chinmay-sawant/gowkhtmltopdf), a layout engine written from scratch in Go.

You pass a template to `New`, store its data with `SetData`, and register handlers for clicks and keys. A click on `data-action` calls your Go function. `examples/login` is the sign-in program.

```go
page, err := gpui.New(gpui.Config{
    Title:  "Hello",
    HTML:   `<h1>{{.Title}}</h1>`,
    Width:  480,
    Height: 640,
})
page.SetData(struct{ Title string }{"Hello"})
gpui.Run(context.Background(), page)
```

## Why this exists

Writing a desktop UI in Go leaves three roads. Electron gives you the whole web platform and ships Chromium and Node with every app, so the installer runs past a hundred megabytes, every message crosses a process bridge, and you track Chromium security releases on someone else's schedule. A Rust GPUI builds the UI in code, in a systems language, behind a render trait and its own layout engine. This library takes the third road: keep HTML and CSS as the UI layer, keep Go for the logic, and ship neither a browser engine nor a JavaScript runtime.

That road needs a layout engine, and the engine is [gowkhtmltopdf](https://github.com/chinmay-sawant/gowkhtmltopdf), a separate project where the HTML parser, the CSS cascade, and the layout pass are written in Go. Blink is how Electron gets a full engine, by shipping all of Chromium. This one lays documents out in Go instead, and hands the placement back as a list of operations. It is not a Blink replacement in features. Of the 818 CSS properties in the webref catalog, 407 are implemented and none is half-finished, audited on 2026-09-21 ([theming.md](documentation/theming.md)). No video, canvas, or WebGL, which is why those are absent here too.

`Redraw` fills your `html/template`, then gowkhtmltopdf parses it, applies the CSS, and lays it out. The window replays that placement as vector operations on the Ebiten canvas, or blits the painted image when an operation has no replay. One process, one binary, no V8, no preload script, no `node_modules`. The same page definition runs on the desktop, as WebAssembly in a browser, and through an Android or iOS bind.

The price is worth stating plainly. The screen is a laid-out picture rather than a live DOM, so there is no JavaScript, no DevTools, and no video, canvas, or WebGL. That list is in [features.md](documentation/features.md), and [compare-electron.md](documentation/compare-electron.md) plus [compare-rust-gpui.md](documentation/compare-rust-gpui.md) put this library next to both alternatives.

## Running the examples

`go run ./examples/<name>` opens a window. `go run ./examples/<name> -web` serves the same screen over HTTP instead. [examples/readme.md](examples/readme.md) lists every example with one line each. The sign-in demo accepts `secret` and `secret`.

```
go run ./examples/login              # desktop window
go run ./examples/login -web         # picture page on 127.0.0.1:8091, -addr changes it
sh scripts/browser.sh                # the same window as WebAssembly on 127.0.0.1:8092
```

Drag an edge to resize; the smallest size is 320 by 400 and the screen is laid out again at the new size. Click a field and type. Ctrl-C, Ctrl-V, Ctrl-X, Ctrl-A, and Ctrl-Z work on the focused field, Enter submits, and the wheel scrolls a page larger than the window.

For a phone, with the Android SDK or Xcode installed:

```
go install github.com/hajimehoshi/ebiten/v2/cmd/ebitenmobile@latest
ebitenmobile bind -target android -javapkg com.chinmaysawant.gogpui -o go-gpui.aar ./examples/login/mobile
ebitenmobile bind -target ios -o go-gpui.xcframework ./examples/login/mobile
```

## Where to read next

Each topic has one file. Nothing here repeats what those files already say.

| Topic | Read |
|-------|------|
| Every feature, grouped, and the list of what is absent | [features.md](documentation/features.md) |
| A dated scan with file citations | [features-examples.md](documentation/features-examples.md) |
| Config, sizing, DPI, audio, window limits | [window.md](documentation/window.md) |
| Desktop, WebAssembly, phone, and the capability matrix | [platforms.md](documentation/platforms.md) |
| Web mode routes and the PNG endpoint | [web.md](documentation/web.md) |
| Template to image, and `Page.PNG` | [screen.md](documentation/screen.md) |
| Display-list replay and the bitmap fallback | [screen.md](documentation/screen.md#replay) |
| Per-frame work with `Page.SetTick` | [frames.md](documentation/frames.md) |
| Keys and text chords | [keys.md](documentation/keys.md) |
| Clicks, hover, taps, and hit-test boxes | [pointer.md](documentation/pointer.md) |
| Wheel scrolling and scrollbar thumbs | [scrolling.md](documentation/scrolling.md) |
| `Config.Theme` and `SetTheme` | [theming.md](documentation/theming.md) |
| In-process `Send`, `Listen`, `Handle`, `Request` | [ipc.md](documentation/ipc.md) |
| `Load`, `Back`, `Forward`, `data-action` routes | [navigation.md](documentation/navigation.md) |
| `Fetch` and `XHR` over `net/http` | [fetch.md](documentation/fetch.md) |
| Copy and paste, including the Wayland fallback | [clipboard.md](documentation/clipboard.md) |
| Typing, undo, redo, select all | [editing.md](documentation/editing.md) |
| `input`, `textarea`, and `select` values | [forms.md](documentation/forms.md) |
| `data-bind` to struct fields | [binding.md](documentation/binding.md) |
| Panic reports on disk | [crash.md](documentation/crash.md) |
| Every example | [examples/readme.md](examples/readme.md) |
| Against Electron, and against the Rust framework | [compare-electron.md](documentation/compare-electron.md), [compare-rust-gpui.md](documentation/compare-rust-gpui.md) |

## What is not here

- No Chromium, V8, Node, preload script, or cross-process IPC. One Go process, one template.
- One window. No tray, native menus, notifications, or second window.
- No accessibility tree, IME, spellcheck, printing, or OS-global shortcuts.
- No DevTools, auto-update, installer, or uploaded crash dump.
- `Fetch` and `XHR` are one http or https request, with no cookies, cache, or session.

## Layout

```
internal/page/       the template, the picture, the display list, the input handlers
internal/render/     html.Parse, css.Apply, layout.Lay, or the display list
internal/replay/     paints the display list on the Ebiten canvas
internal/frame/      finds fill and text operations for a SetTick callback
internal/window/     the native window, and the same loop on a phone or in a browser build
internal/web/        the picture page on 127.0.0.1
internal/ipc/        in-process Send, Listen, Handle, Request
internal/fetch/      one http or https request
internal/clipboard/  the OS clipboard, with an in-memory copy
internal/crash/      a local panic report
internal/filepick/   the desktop open dialog
examples/login/      the sign-in program
browser/index.html   the page that loads the WebAssembly build
scripts/browser.sh   builds that page and serves it
```

## Development

`make build` compiles every package and links nothing. `make test` runs `go test -p 1 ./...`. Leave `go build ./...` alone; it links an executable per example and leaves the binaries here.

`go.mod` carries a local `replace` to `../gowkhtmltopdf`, because the display-list path needs `layout.DisplayList` and `Display.Boxes` and the pinned version does not export them. Keep [gowkhtmltopdf](https://github.com/chinmay-sawant/gowkhtmltopdf) as a sibling directory next to this one, or the build fails. Text replay needs Ebiten v2.10.4 or newer.
