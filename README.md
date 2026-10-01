# go-gpui

A window for an HTML page. You write the Go data and the handlers.
A click on `data-action` calls your Go function.

1. The login template is HTML.
2. `gpui.Page.Redraw` calls `gowkhtmltopdf/html.Parse`, `gowkhtmltopdf/css.Apply`, and `gowkhtmltopdf/layout.Lay`.
3. The desktop and browser-canvas window draw that layout image directly.
4. The `-web` page encodes PNG only for `GET /frame.png`.

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

## Layout

```
internal/page/            the template, the picture, and the input handlers
internal/window/          the native window, and the same loop on a phone or in a browser build
internal/web/             the picture page on 127.0.0.1
examples/login/           the sign-in program
examples/login/login/     the sign-in template and its Go handlers
browser/index.html        page that loads the WebAssembly build
scripts/browser.sh        builds that page and serves it
skills/                   copied from the gowkhtmltopdf skills folder
```

The sign-in program does not import the window or the web package.
It calls `gpui.Run` or `gpui.Serve`.

Feature notes for this branch live in `documentation/features.md`.
That set covers the screen paint, in-process IPC, HTML history,
local crash files, `Fetch` / `XHR`, and the OS clipboard.

## Desktop window

From this directory:

```
go run ./examples/login
```

That opens a normal window with a title bar. Drag an edge to resize it.
The smallest size is 320 by 400. After you stop dragging, the HTML screen
is drawn again at the new size. Click a field and type. Ctrl-C copies that
field, Ctrl-V pastes, Ctrl-X cuts, Ctrl-A selects it, and Ctrl-Z undoes.
Enter signs in. Backspace deletes. A wrong password prints an error. The
demo password prints Signed in.

The picture page is still there:

```
go run ./examples/login -web
```

Open http://127.0.0.1:8091/. `-addr` changes that address.
That page encodes PNG only for GET /frame.png.

## Browser build

The window loop also compiles to WebAssembly:

```
sh scripts/browser.sh
```

Open http://127.0.0.1:8092/. Resizing the browser changes the frame.
This is the same screen as the desktop window. The canvas draws the
layout image directly. `-web` is a separate, simpler page. It encodes
PNG only for GET /frame.png.

## Phone

`examples/login/mobile` registers that same screen with Ebitengine's mobile view.
It does not call `Run`. From this directory, with the Android SDK
or Xcode installed:

```
go install github.com/hajimehoshi/ebiten/v2/cmd/ebitenmobile@latest
ebitenmobile bind -target android -javapkg com.chinmaysawant.gogpui -o go-gpui.aar ./examples/login/mobile
ebitenmobile bind -target ios -o go-gpui.xcframework ./examples/login/mobile
```

The Android bind writes an `EbitenView`. The iOS bind writes a view
controller. Taps are clicks. The view fills the screen, so rotating the
phone or changing the split changes the frame.

Demo login: `secret` / `secret`.

`go.mod` requires the published `github.com/chinmay-sawant/gowkhtmltopdf` module.
