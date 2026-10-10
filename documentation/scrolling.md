# Scrolling

A page taller or wider than the window scrolls on the mouse wheel, a touch drag, or a request from the page. The window moves the page and owns the offset.

## Wheel

`internal/window/fit.go` moves the offset 48 px per wheel notch. A positive wheel moves toward the start of the page, a negative wheel toward the end, and the horizontal wheel pans sideways. The offset is clamped to the content, so a wheel or a drag cannot pass an edge. Holding Ctrl (or Command) while turning the wheel zooms the page instead of scrolling it; [keys.md](keys.md) has the zoom keys and the range.

`internal/window/scrollbar.go` measures the content for that clamp: the canvas size, grown to cover every box that overflows it. A fixed-width child wider than the window scrolls too. The replay buffer uses the same measure, so scrolled content stays painted.

While the picture is stretched to the window, `internal/window/view.go` ignores the wheel and `internal/window/scrollbar.go` and `internal/window/scrollbar_drag.go` hide the thumbs.

## Thumbs

A page that overflows gets a thumb overlay, 8 px thick and at least 24 px long, in `#BDBDBD`. The vertical thumb sits on the right edge and the horizontal one on the bottom edge. Only the thumb is painted, so the page background stays visible under it.

Press a thumb to start a drag. The drag keeps the distance between the cursor and the thumb's start and continues when the cursor leaves the thumb. The window consumes the drag, so it does not also click the page. A press on the track does not page-jump; it misses the thumb and falls through to the content. Thumb drag is mouse-only.

## Coordinates

The window maps a cursor before the page sees it ([pointer.md](pointer.md)). A scrolling page adds the offset, so page coordinates start at the content's top left. A stretched picture scales the cursor instead, so a point halfway across the window is halfway across the picture. Both mappings are in `internal/window/fit.go`.

## Page requests

`Page.ScrollTo(x, y)` and `Page.ScrollBy(dx, dy)` store one request. The window consumes it through `host.ScrollRequester`, clamps it to the content, and clears it. The offset stays owned by the shell, so a request, the wheel, a thumb drag, and a touch drag all move the same value. [interaction.md](interaction.md) has the calls.

## Touch

A finger that moves past 8 px drags the page, clamped to the content ends. After lift, the page continues with a decaying fling. Mobile updates follow the display refresh. `Page.SetTouchScrollOptions` tunes finger sensitivity and how quickly the fling slows; both settings have bounded defaults. Two fingers pinch a zoom that stays between 0.25 and 4. [interaction.md](interaction.md) has the gesture details.

## Limits

- The offset survives `Load`, `Back`, and `Forward` ([navigation.md](navigation.md)).
  The window clamps it to the content after any relayout or redraw, so a
  shorter page cannot show empty space below it.
- Desktop and wheel scrolling keep their existing step behavior. Touch scroll
  inertia applies only to touch gestures.

The [scrolling example](../examples/scrolling) is a column of 40 rows in a 360x480 window.
