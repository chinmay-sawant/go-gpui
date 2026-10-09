package page

import "context"

// Handlers are the Go functions for input.
// Nil functions ignore input. Key and drag callbacks do not draw.
type Handlers struct {
	Hover HoverHandler
	Click func(ctx context.Context, box Box) error
	// Press can activate on pointer down and consume the following click.
	Press func(ctx context.Context, box Box, x, y float64) (bool, error)

	// LongPress handles a hold outside a text field.
	LongPress func(ctx context.Context, box Box) error

	// KeyDown receives a lowercase key name. It does not draw.
	KeyDown func(ctx context.Context, key string) error

	// KeyUp receives a released key name. It does not draw.
	KeyUp func(ctx context.Context, key string) error

	Swipe func(ctx context.Context, dx, dy float64) error

	// DragStart claims a pointer gesture beginning inside a box.
	DragStart func(ctx context.Context, box Box) (bool, error)
	// DragMove receives CSS-pixel deltas while a claimed pointer moves.
	DragMove func(ctx context.Context, dx, dy float64) error
	DragEnd  func(ctx context.Context) error

	// BeforeEdit can reject a change before the control is edited.
	BeforeEdit func(ctx context.Context, box Box) error

	// Change runs after editing, before drawing. Errors skip the redraw.
	Change     func(ctx context.Context, box Box) error
	Type       func(ctx context.Context, text string) error
	Backspace  func(ctx context.Context) error
	DeleteWord func(ctx context.Context) error
	Submit     func(ctx context.Context) error
	Copy       func(ctx context.Context) (text string, ok bool, err error)
	Cut        func(ctx context.Context) (text string, ok bool, err error)
	Paste      func(ctx context.Context, text string) error
	SelectAll  func(ctx context.Context) error
	Undo       func(ctx context.Context) error
	Redo       func(ctx context.Context) error

	// Drop runs when files land on the window. Each file's Read func is
	// valid only during the call.
	Drop func(ctx context.Context, files []Drop) error
}
