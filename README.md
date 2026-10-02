# go-gpui

A window for an HTML page. You write the Go data and the handlers.
A click on `data-action` calls your Go function.

1. The login template is HTML.
2. `gpui.Page.Redraw` asks `gowkhtmltopdf` for the placement as a display list. A page `render.Replayable` accepts keeps the vector operations and no bitmap; any other page falls back to `layout.Lay` and an image.
3. The desktop and browser-canvas window replays the display list, or draws the fallback image directly.
4. The `-web` page encodes PNG on demand only for `GET /frame.png`.

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
internal/page/            the template, the picture, the display list, and the input handlers
internal/replay/          paints the display list on the Ebiten canvas
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
That set covers the screen paint, the retained display list, in-process IPC,
HTML history, local crash files, `Fetch` / `XHR`, form controls, and the OS
clipboard.

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
This is the same screen as the desktop window. The canvas replays the
display list, or draws the fallback image directly. `-web` is a separate,
simpler page. It encodes PNG only for GET /frame.png.

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

## Display list

`go.mod` requires a `gowkhtmltopdf` version that exports `layout.DisplayList`
and `Display.Boxes`. The display-list operations live behind that module's
`internal/` rule, so only a version carrying the export can hand them over.

This tree carries a local `replace` to `../gowkhtmltopdf` because the `Boxes`
field is not in the pinned pseudo-version yet. Drop the replace and bump the
pin once upstream carries it.

`internal/render.DisplayList` returns the placement as vector operations and no
picture. `internal/window` replays them through `internal/replay`; a page with
an operation the replay does not support keeps the `render.Paint` bitmap. See
[documentation/screen.md](documentation/screen.md).
