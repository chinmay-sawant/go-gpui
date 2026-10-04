# Pointer

The window sends hover, press, release, and click events to a page. `internal/window/pointer.go` calls the page methods each frame, and a caller can call the same methods directly.

## Page methods

`Page.Hover(ctx, x, y)` sets the hovered id to the innermost box with an id at the point. A point that hits no box with an id clears it. It redraws only when the id changed. A hover change repaints the union of the box it left and the box it entered, not the rest of the frame.

`Page.Press(ctx, x, y)` sets the pressed id the same way and redraws only when it changed. `Page.Release(ctx)` clears the pressed id and redraws only when it was set.

`Page.Click(ctx, x, y)` hit-tests the last box that contains the point, which is the innermost box in document order. Then:

- A form control activates before the handler runs and never follows a route ([forms.md](forms.md)).
- Any other hit blurs the focused control.
- After a non-control hit, a non-empty `data-action` with a registered route loads that HTML and skips the handler ([navigation.md](navigation.md)).
- A nil `Click` handler ignores the event.
- An error from the handler or the load returns before the redraw. Otherwise the page redraws after the handler.

## Window mapping

The window sends hover every frame. A scrollbar interaction consumes the event before hover runs, so hover pauses while a thumb is pressed or dragged. Press and click both fire on the mouse-down edge; release fires on the mouse-up edge.

A press in a field also calls `SelectAt`, and each move with the button down calls `Drag`, so one press, move, release selects a span. Two rapid presses call `SelectWordAt` and a third calls `SelectLineAt`. Dragging past the top or bottom edge scrolls while the button stays down. [interaction.md](interaction.md) has the details.

A right click asks the page for its context menu rows and draws them at the cursor. A left click on a row runs its action; a left click elsewhere or Escape closes the menu.

The cursor shape follows the hovered element when the page implements `host.CursorShape`: an I-beam over a field, a hand over a link or button, a resize cursor over a scrollbar thumb.

The coordinates arrive in page space, with the scroll offset added or the stretch scale applied ([scrolling.md](scrolling.md)).

A touch that lifts without moving sends press, click, and release as a tap. A moved touch drags the page, and two fingers pinch a zoom. A mouse click in the same frame suppresses the tap, so one gesture is not delivered twice.

## CSS states

The hovered and pressed ids become the host states `:hover` and `:active`. `:focus` comes from clicks and from Tab, and the focus ring draws at each keyboard stop ([interaction.md](interaction.md)). [forms.md](forms.md) lists the state attributes and the focus rules.
