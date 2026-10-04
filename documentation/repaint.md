# Repaint

A click that changes one box used to rebuild and replay the whole page. Now
the page reports the box it changed and the window repaints only that box.
This is the automatic path. The manual path, a frame callback that edits an
operation, is in [frames.md](frames.md).

## The page reports the region

`Page.TakeDirty` returns the region that changed since the previous call, in
CSS pixels, and clears it. A handler can declare the id it changed with
`Invalidate`; for handlers that forget, the page diffs the previous display
list against the new one and unions the bounds of every changed operation.
An operation removed from the list dirties its old bounds, so a vanishing
element leaves no ghost. A change the page cannot place, such as a theme
switch or a `Load`, returns the whole frame.

```go
rect, ok := page.TakeDirty()
```

`ok` false means nothing changed since the previous call. The rect is
padded, so shadows and outlines repaint too.

## The window repaints the rect

`internal/replay.DrawRect` is `Draw` for one region. It walks `Display.Order`
the same way and skips any op whose painted box does not meet the rect.

```go
replay.DrawRect(dst, display, rect, dx, dy)
```

Text needs care: an op carries its baseline in `Y`, so `DrawRect` tests the
line box from the font ascent above the baseline to `InkDescent` below it,
not `Y` alone. Strokes and lines grow by half a stroke, and an op whose ink
cannot be bounded, such as a rotated run, is always drawn.

The window keeps one canvas-sized image for the replay path. A dirty rect is
filled with the page background and replayed into that buffer, and the
visible page is blitted from the buffer to the screen each frame. The screen
is cleared between frames, so the blit covers the viewport and only the
buffer repaint is partial. Scrollbar thumbs and the `bitmap fallback` badge
draw on the screen after the blit, so a repaint never wipes them.

The buffer is rebuilt in full when the generation changes without a usable
rect, when the canvas changes size, and when the rect is the whole frame. A
scroll offset change only moves the blit, so scrolling does not replay the
list.

## What still repaints in full

- A fallback page with no display list. `imageout.RenderLayout` paints the
  whole canvas, and a partial bitmap path needs a new engine call; the
  decision is recorded in
  [the plan](../plans/v0.0.2/incremental-repaint.md). The badge stays
  visible.
- A page with a frame callback. A tick changes operations in place and
  leaves no dirty rect, so `Page.Ticking` keeps the full replay.
- `Page.PNG` paints from the template source on demand, so `GET /frame.png`
  always costs a full paint.

## Measured numbers

Against the examples in this tree, at 480x640, at this commit:

| Page | Ops | A small box selects |
|------|-----|---------------------|
| platform | 6 | 1 |
| login | 18 | 1 |
| states | 16 | 1 |
| forms | 45 | 4 |

The op filter runs at about 13 ns per op (`go test -bench BenchmarkDrawRectScan
-v ./internal/replay`, which also prints the platform counts). The platform
page scans in about 80 ns per frame. The saving is in the operations it
skips, not in the scan: a click on the platform counter replays 1 op and
skips 5, and the same click on the forms page replays 4 and skips 41.
