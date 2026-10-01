# go-gpui

Local HTML screens on top of gowkhtmltopdf. A screen is an HTML file.
`internal/login` fills the `{{ }}` holes, asks `gowkhtmltopdf/screen` for a
PNG and the element rectangles, and calls a Go function when a click lands
on `data-action`.

## Layout

```
cmd/go-gpui/main.go       starts the process
internal/login/           the screen, the template, and the Go actions
internal/window/          the native window, and the same loop on a phone or in a browser build
internal/web/             the older page on 127.0.0.1, used only with -web
mobile/                   Android and iOS bind target
browser/index.html        page that loads the WebAssembly build
scripts/browser.sh        builds that page and serves it
skills/                   copied from the gowkhtmltopdf skills folder
```

`internal/login` does not import the window or the web package. Both call
`Redraw`, `Click`, `Type`, `Backspace`, and `Submit`.

## Desktop window

From this directory:

```
go run ./cmd/go-gpui
```

That opens a normal window with a title bar. Drag an edge to resize it.
The smallest size is 320 by 400. After you stop dragging, the HTML screen
is drawn again at the new size. Click a field and type. Enter signs in.
Backspace deletes.

The older browser page is still there:

```
go run ./cmd/go-gpui -web
```

Open http://127.0.0.1:8091/. `-addr` changes that address.

## Browser build

The window loop also compiles to WebAssembly:

```
sh scripts/browser.sh
```

Open http://127.0.0.1:8092/. Resizing the browser changes the frame.
This is the same screen as the desktop window. `-web` is a separate,
simpler page that only shows the PNG.

## Phone

`mobile` registers that same screen with Ebitengine's mobile view.
It does not call `RunGame`. From this directory, with the Android SDK
or Xcode installed:

```
go install github.com/hajimehoshi/ebiten/v2/cmd/ebitenmobile@latest
ebitenmobile bind -target android -javapkg com.chinmaysawant.gogpui -o go-gpui.aar ./mobile
ebitenmobile bind -target ios -o go-gpui.xcframework ./mobile
```

The Android bind writes an `EbitenView`. The iOS bind writes a view
controller. Taps are clicks. The view fills the screen, so rotating the
phone or changing the split changes the frame.

Demo login: `ada@example.com` / `secret`.

`go.mod` replaces `github.com/chinmay-sawant/gowkhtmltopdf` with the sibling
checkout `../gowkhtmltopdf`. Nothing here is pushed.
