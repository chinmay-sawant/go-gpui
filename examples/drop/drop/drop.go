// Package drop is the drag-and-drop example.
// A drop of any file prints the file's path. The window delivers the files;
// gpui opens the window. This package does not. A file control takes a path
// from the picker when a drag cannot reach the window.
package drop

import (
	_ "embed"

	"github.com/chinmay-sawant/go-gpui"
)

//go:embed drop.html
var dropHTML string

const (
	// DefaultWidth and DefaultHeight are the size of a newly opened window.
	DefaultWidth  = 560
	DefaultHeight = 460
)

// View is the list of paths the template prints.
type View struct {
	Paths []string
}

// App is the drop screen.
type App struct {
	page *gpui.Page
	view View
}
