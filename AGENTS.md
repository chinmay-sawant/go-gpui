# go-gpui

Go library that opens an HTML template in a window. The sign-in program lives in `examples/login` and is one caller of the library.

## Where code lives

Package `gpui` at the module root is the public API. `New`, `SetData`, and `Handle` build a page. `Run` opens the desktop or wasm window. `Serve` shows the picture page. `BindMobile` registers the phone view. The root files are aliases and those entry points. The page implementation is `internal/page`.

`internal/window` is the Ebiten loop for desktop, phone, and `GOOS=js GOARCH=wasm`. `browser/index.html` loads that wasm build. `scripts/browser.sh` builds and serves it.

`internal/web` is the `go run ./examples/login -web` page. It encodes PNG for `GET /frame.png`.

`internal/host` defines `Screen`. The window and web packages call that interface. They do not import package `gpui`.

`internal/clipboard` reads and writes the OS clipboard, and keeps an in-memory copy. Tests call `UseMemory`.

`internal/render` paints the screen. `Redraw` calls `render.DisplayList` first. A page `render.Replayable` accepts keeps a `layout.Display`; any other page calls `render.Paint`, which is `html.Parse`, `css.Apply`, and `layout.Lay`. Neither path writes a PDF.

`internal/replay` draws the display list on the Ebiten canvas: fills with circular corners, axis-aligned border lines, grid runs, and shaped text. `internal/window/draw.go` picks the display list or the fallback image; a fallback frame shows a `bitmap fallback` badge in its top-right corner.

`internal/ipc` is in-process `Send`, `Listen`, `Handle`, and `Request`. `ipc.go` re-exports them.

`Page.Load`, `Back`, `Forward`, `Route`, and `HTML` live in `internal/page`. `nav.go` exports `ErrNoHistory`.

`internal/crash` writes a local report. `Run` and `BindMobile` recover a panic and return that path. `crash.go` exports `Report` and `SetCrashDir`.

`internal/fetch` sends one http or https request and stores no cookies. `fetch.go` exports `Fetch`, `XHR`, `FetchResponse`, and `ErrScheme`.

`github.com/chinmay-sawant/gowkhtmltopdf` parses the HTML, applies the CSS, and lays the page out. A replayable page keeps `layout.Display` operations and no picture; any other page paints `Page.Image`. The window replays or blits accordingly. Text replay needs Ebiten v2.10.4 or newer, because gowkhtmltopdf requires `go-text/typesetting` v0.3.4 and older Ebiten builds its font face without the lookup cache v0.3.4 added. `go.mod` replaces the module with `../gowkhtmltopdf` until upstream carries `Display.Boxes`; drop the replace and bump the pin then.

Read `documentation/features.md` before changing paint, IPC, navigation, crash reports, fetch, or the clipboard. The call shapes and the limits are in the other files under `documentation/`.

A page taller or wider than the window scrolls on the mouse wheel. A picture that matches the window stays at one CSS pixel per window pixel. The sign-in example accepts the email `secret` and the password `secret`.

## Go file size

Every Go file is at most 2000 characters. Count with `wc -m`. That number includes newlines. When a file would exceed 2000 characters, split it inside the same package and keep the behavior. One file in the package carries the package comment. The other files start with the package clause.

This limit applies to Go files only. HTML, CSS, Markdown, and scripts have no character cap.

## Before you finish

Run `gofmt` on every Go file you edit. From this directory, `go test ./...` passes.

Leave `~/.Xauthority` untouched. The window toolkit logs a missing authority file and still opens the window.

Keep clipboard tests off the desktop clipboard. Call `clipboard.UseMemory` in those tests.
