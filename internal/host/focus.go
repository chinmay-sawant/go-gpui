package host

import "context"

// Focuser is a screen that moves focus between its controls. The window
// consumes Tab and Shift+Tab only when the screen implements it.
type Focuser interface {
	FocusNext(ctx context.Context) error
	FocusPrev(ctx context.Context) error
	Focus(ctx context.Context, id string) error
	FocusID() string
}
