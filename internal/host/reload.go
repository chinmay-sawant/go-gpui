package host

import "context"

// Reloader is a screen whose source files are watched. The window polls it;
// a true return means the page changed and the window should show the new
// generation.
type Reloader interface {
	PollReload(ctx context.Context) (bool, error)
}
