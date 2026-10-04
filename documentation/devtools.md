# DevTools

The window can draw an inspector over the page. It outlines the element under
the cursor, pins one to read its tag and geometry, outlines the display
operations, and prints the frame numbers. The overlay is window chrome, like
the scrollbar thumbs and the fallback badge. It never enters `Page.PNG`, the
display list, or the box list.

## Turning it on

Press F12 or Ctrl+Shift+I. The window owns both keys while a page has an
inspector, so `Handlers.KeyDown` and `Handlers.KeyUp` never see them, and
the matching release is swallowed too. Pressing the key again closes the
overlay and sends one hover for the current cursor, so a stale `:hover` does
not stick.

`Config.DevTools` starts the overlay on, and `Page.SetDevTools` and
`Page.DevTools` read and change the flag at runtime. Use those in a browser
and on a phone: a browser may keep F12 for itself, and a phone has no
keyboard.

```go
page, err := gpui.New(gpui.Config{
    HTML:     html,
    Width:    720,
    Height:   560,
    DevTools: true,
})
```

## The panel

The panel sits at the bottom left, above the horizontal scrollbar strip, and
grows with its longest line. The badge in the top right stays visible. Each
line:

| Line | Meaning |
|---|---|
| `[ ] ops mode (o)` | The operation view. Click the line or press `o` to toggle it. |
| `fps` and `tps` | `ebiten.ActualFPS` and `ebiten.ActualTPS`. |
| `frame` and `draw` | Wall time between two `Draw` calls and the time around `drawContent`. The display lookup in `sync.go` can dominate the second number. |
| `window` and `app` | The window size and `Page.Size()`. |
| `scroll`, `stretched`, `fallback`, `seq` | The scroll offset, whether the page is scaled to the window, whether the last frame came from a bitmap, and the generation the window synced. |
| `redraws` | The `host.Inspector.Stats` counters: redraws, parses, cascades, layouts, and repaints. |
| `boxes` and `ops` | The box and operation counts, plus the last redraw and the last draw. |
| `reloads` and `reload error` | Hot reload counts when a page reports them. `-` means none. |
| `hover` and `pin` | The tag, id, geometry, action, text, and operation count of the picked box. |

## Picking

While the overlay is on, the pointer belongs to the overlay. Moving the
cursor outlines the innermost box under it, and the page picture stays still.
The window sends no hover, so the pixels the numbers describe do not change.
A click pins the box, and a click on the same box or anywhere on the panel
clears the pin. The panel hit-tests its own rows, never `Page.Boxes()`.

Alt+click forwards the press and the click to the page, so a control can
still be exercised while inspecting. Closing the overlay clears the hovered
and pinned boxes, and a new generation drops either one when its id is gone
from `Boxes()`.

## Operations

Ops mode outlines every operation in paint order. The colour names the kind:

| Kind | Colour |
|---|---|
| fill | `#1a56db` |
| stroke | `#e5484d` |
| line | `#d97706` |
| text | `#176b45` |
| image | `#7c3aed` |
| grid run | `#0891b2` |

The bounds come from the operation's `X`, `Y`, `W`, and `H`, converted with
`Display.PixelPerPoint`. A text operation carries its baseline in `Y`, so the
outline uses the line box from its face metrics. `DisplayOpNoop` and
`DisplayOpUnknown` paint nothing and get no outline. A fallback page has no
display list, so the panel prints `bitmap fallback` and draws nothing.

The panel counts each kind and the total. A growing number widens the panel
instead of clipping it.

## A custom screen

The window type-asserts `host.Inspector` on the screen it already holds:

```go
type Inspector interface {
    DevTools() bool
    SetDevTools(on bool)
    Stats() Stats
}
```

A screen without it gets no overlay and no compile change. `Stats` is the
snapshot the panel prints. A screen that also implements
`SetDrawTime(time.Duration)` receives the window's draw time, and
`Stats().LastDraw` reports it. That hook is separate from `host.Inspector`,
so the interface keeps its three methods.

## Web mode

`GET /debug/state` returns the same data as JSON: size, generation, fallback,
boxes, stats, and a per-kind operation count map. The route is in
[web.md](web.md).

## Limits

There are no computed styles to show. The engine styles a document internally
and publishes no per-element reader, so the panel shows geometry and
attributes, not `font-size` or `margin`. Suppressing pointer input freezes
the page's hover and press state at their last value. F12 is a browser key in
a wasm build, and `Config.DevTools` or `SetDevTools` is the cross-platform
way in.
