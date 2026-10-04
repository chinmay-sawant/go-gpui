package host

import "context"

// Drop is one file or directory dropped on the window. Path is absolute on
// desktop and empty in a browser; Read serves the bytes only during the Drop
// call.
type Drop struct {
	Name  string
	Size  int64
	IsDir bool
	Path  string
	Read  func() ([]byte, error)
}

// Dropper is a screen that accepts dropped files.
type Dropper interface {
	Drop(ctx context.Context, files []Drop) error
}
