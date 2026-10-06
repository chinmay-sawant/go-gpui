package page

import "context"

// Handlers are the Go functions for input.
// A nil function ignores that input. After a function returns the page
// is drawn again, except after KeyDown, KeyUp, and Copy.
type Handlers struct {
	Hover HoverHandler
	Click func(ctx context.Context, box Box) error

	// LongPress runs when a press is held in place. A text field under the
	// point selects its word and skips the handler.
	LongPress func(ctx context.Context, box Box) error

	// KeyDown runs when a key goes down. key is a lowercase name such as
	// "space", "arrowup", "a", or "1". It does not draw: change state and
	// let a SetTick callback paint, or call Redraw.
	KeyDown func(ctx context.Context, key string) error

	// KeyUp runs when a key goes up, with the same names as KeyDown. It
	// does not draw.
	KeyUp func(ctx context.Context, key string) error

	// BeforeEdit runs before a control's value, checked state, or
	// selection changes, before the built-in edit and before Change. An
	// error aborts the edit and skips the redraw.
	BeforeEdit func(ctx context.Context, box Box) error

	// Change runs after a control's value, checked state, or selection
	// changed, before the page is drawn again. An error skips the redraw
	// and is returned.
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
