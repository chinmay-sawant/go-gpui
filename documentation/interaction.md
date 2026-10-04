# Interaction

The window owns the focus keys, the pointer gestures, the context menu, the cursor shape, touch gestures, fullscreen, and the scroll offset. The page owns the caret, the selection range, the focus order, and the scroll requests. A page that implements none of the optional surfaces keeps the v0.0.1 behavior.

## Focus

Tab and Shift+Tab move focus in document order when the screen implements `host.Focuser` and a field takes the focus. The window calls `FocusNext` or `FocusPrev`, then reads `FocusID`. An empty id means no field moved, so the key reaches the page handler as before. When a field moves, the window consumes the press and its release, so `Handlers.KeyDown` does not see `"tab"`.

The page sets the same `:focus-visible` host state a click sets, so the focus ring draws at each stop. Escape closes the context menu first. With the menu closed it calls `Focus(ctx, "")` and clears the ring. Escape itself still reaches the handler. `internal/window/focus_keys.go` holds the window side.

## Caret and selection

A press in a field calls `SelectAt(ctx, x, y)`, which focuses the field under the point and puts the caret at the glyph. It runs before the click handler, so a handler that focuses or selects a field keeps it. A move with the button down calls `Drag(ctx, x, y)`, so one press, move, release selects a span. Dragging past the top or bottom edge scrolls one wheel step per frame, clamped to the content.

Two rapid presses at one spot call `SelectWordAt`; a third calls `SelectLineAt`. The gap is 400 ms and the spot may move 4 px. `internal/window/drag_select.go` and `internal/window/click_watch.go` hold the window side. The page holds the caret and the range ([editing.md](editing.md)).

## Context menu

A right click asks the page for rows with `ContextMenu() []host.MenuItem` and draws them at the cursor. Row ids are `cut`, `copy`, `paste`, `select-all`, `undo`, and `redo`, and an action calls the screen method of the same name. Paste reads the clipboard the way the Ctrl+V chord does. A disabled row draws grey and does nothing.

A left click on a row runs the action and closes the menu. A left click elsewhere or Escape closes it, and a new right click reopens it at the new point. The menu is chrome. `internal/window/menu_draw.go` paints it after the scrollbar thumbs, it never enters `Page.PNG`, the display list, or a `Redraw`, and no page method runs when it opens or closes.

## Cursor shape

When the screen implements `host.CursorShape`, the window reads `CursorShape()` after each hover and sets the Ebiten cursor only when the shape changed. Text fields and textareas map to the I-beam, links and controls to the hand, scrollbar thumbs to the matching resize cursor, and everything else to the default arrow. wasm and mobile keep their own cursor and ignore the call. The last shape stays on the shell for tests.

## Touch

A finger that lifts without moving sends press, click, and release as a tap. A finger that moves past 8 px drags the page instead: the content follows the finger, clamped to the content ends. Two fingers pinch a zoom that starts at the span of the first two fingers and stays between 0.25 and 4. The replay and bitmap paths draw through the zoom, and `contentPoint` divides it out, so a click under a pinch lands on the same box. A tap that the system also reports as a mouse click does not tap twice.

## Programmatic scroll

`Page.ScrollTo(x, y)` and `Page.ScrollBy(dx, dy)` store one request. The window consumes it through `host.ScrollRequester`, clamps it to the content size, draws the thumbs from the new offset on the next frame, and clears it. The offset stays owned by the shell, so the wheel, a thumb drag, a touch drag, and a page request all move the same value ([scrolling.md](scrolling.md)).

## Fullscreen

F11 toggles `ebiten.SetFullscreen` on desktop and the window consumes the key. wasm and mobile keep their own fullscreen, the call is a no-op there, and F11 reaches the page. `internal/window/fullscreen.go` holds the window side.

## Limits

- Up and Down in a textarea, and selection beyond the visible scroll of a field, are out of scope. The value model is a single line.
- IME is not in this release. Ebiten v2.10.4 has no preedit API, and [features.md](features.md) records the ask.
- The context menu is not native. It is a shell rectangle drawn with the badge face.
