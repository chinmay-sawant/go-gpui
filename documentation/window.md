# Window and sizing

`New` takes the first frame size from `Config`. `Run` opens a window at that size. When the user resizes the window, the window lays the page out again at the new size and scrolls a page larger than the window.

## Config

```go
page, err := gpui.New(gpui.Config{
    Title:     "Inbox",
    HTML:      `<h1>Inbox</h1>`,
    Theme:     `h1 { color: navy }`,
    Width:     480,
    Height:    320,
    MinWidth:  320,
    MinHeight: 200,
    MaxWidth:  1280,
    MaxHeight: 800,
})
```

`Width` and `Height` are the first frame, in CSS pixels. `MinWidth` and `MinHeight` are the smallest frame; zero means 1. `MaxWidth` and `MaxHeight` are the largest window the host asks the OS for; zero means 2560. They do not cap the picture. `New` clamps the first frame up to the minimum.

`New` returns `ErrBadSize` when `Width` or `Height` is not positive, or when a min is greater than its max. A blank template returns `ErrEmptyHTML` first ([screen.md](screen.md)). `Theme` is the extra stylesheet ([theming.md](theming.md)).

## Page size

`Size` returns the frame size in CSS pixels. `MinSize` returns the smallest frame and `MaxSize` the window bound from `Config`. `SetSize(width, height)` stores the size the next `Redraw` uses, pulled up to the minimum; it does not draw. `Clamp(width, height)` returns a size pulled up to the minimum without storing it. Neither one caps a frame at the maximum.

`Generation` increases by one on every successful `Redraw`, on both the replay and bitmap paths. `SetSize`, `SetData`, and `SetTheme` do not move it on their own. `internal/window/sync.go` caches the drawn artifact with the generation it came from and rebuilds it when `Generation` differs.

## The window

`Run` opens a decorated, resizable window. The initial size is the page `Size`, the OS minimum is `MinSize`, and the OS maximum is `MaxSize` when the screen has one. A desktop window cannot grow past that bound; the layout follows whatever size the window has. On a WebAssembly build the browser gives the canvas size, and `MaxSize` does not apply. On a phone the system gives the screen size, so rotation and split screen work the same way.

## Resize

Ebiten calls `Layout` with the outside size when the window changes. `Layout` pulls a size below 1 up to 1, records it, and returns it. The next update commits it in `resize`:

- A size that held still for one update is laid out at once.
- A size that keeps changing relayouts at most every 100 ms, so text rewraps while the drag runs.
- Between commits the previous frame is scaled to the window.

The last size of a drag always commits, even when the throttle window is still open, and a burst of window events between two updates costs one relayout. The commit calls `SetSize` and `Redraw`, then re-resolves the hover and pressed ids against the new boxes. `internal/window/sync.go` rebuilds the drawn artifact only when the generation or the size changed, so a frame that matches both is drawn from the cache.

A file-backed page relayouts when its source changes too ([hot-reload.md](hot-reload.md)). A reload keeps the frame size, and a page that got shorter pulls the scroll offset back inside the content.

## Fit

`stretched` in `internal/window/view.go` reports that the drawn frame does not match the window. Since the cap became a window bound, that only happens while a relayout is in flight. A committed frame matches the window, so it draws at 1:1, one CSS pixel per window pixel, and scrolls when its content overflows.

The replay path draws through an offscreen buffer at the frame size and scales the buffer. The bitmap path scales the image directly. The wheel is ignored while the picture is stretched, so a scaled picture never scrolls ([scrolling.md](scrolling.md)). The `bitmap fallback` badge comes from the cached artifact, so it does not blink during a drag.

## DPI

The library has no DPI code. No code reads a device scale factor or applies one to a size. `Config.Width`, `SetSize`, and the window all use Ebiten device-independent pixels, so a frame that matches the window is one CSS pixel per window pixel.

## Title

`Config.Title` is the window title. Empty means `go-gpui`. `Page.Title` returns it. `Run` reads it once when the window opens, and there is no runtime setter, so a title change after that has no effect.

## Pacing and update order

Ebiten updates the window at its default rate, about 60 TPS. The library never calls `SetTPS`, so the rate is not tunable here. Every `Update` runs keys, resize, pointer, drop, tick, a reload poll, sync, the page scroll request, and wheel in that order (`internal/window/game.go`). Resize runs before pointer, so a hit test uses the boxes of the frame being drawn. The tick is `Page.SetTick`; its call shape and cost are in [frames.md](frames.md). The reload poll asks a file-backed page for changes at most every 250 ms ([hot-reload.md](hot-reload.md)).

## Errors

A tick callback that returns an error stops the window, and `Run` returns it. Cancelling the context does the same. Any handler error behaves the same way: `Update` returns it to Ebiten, the window closes, and `Run` returns it.

## Audio

`Run` and `BindMobile` create a 48 kHz Ebiten audio context once, when one is not already running (`audio_context.go`). `Serve` does not. The library has no audio API; the examples play through [examples/music](../examples/music) ([features.md](features.md)).

## Showing the page

- `Run` opens the desktop window, or the browser canvas on wasm, and blocks until it closes.
- `BindMobile` registers the page with Ebitengine's mobile view. Call it from the package that `ebitenmobile bind` compiles, and do not call `Run` from that package. It returns after registering.
- `Serve` listens on `addr` and blocks. The page at `/` shows the latest picture, and clicks and the type form call the same handlers as `Run` ([forms.md](forms.md)).

All three return `ErrNilPage` when the page is nil. `Run` and `BindMobile` recover a panic and return the crash report path ([crash.md](crash.md)). Build targets are in [platforms.md](platforms.md).

## Limits

- No second window, no window icon, and no window position. Fullscreen, the cursor shape, touch scroll and pinch, and programmatic scroll are in [interaction.md](interaction.md).
- No runtime title setter.
- `Run` prints `opening a window` to stdout once, before the loop starts.
