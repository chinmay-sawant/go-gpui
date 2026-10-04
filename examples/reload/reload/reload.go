// Package reload is the hot reload example.
// The page is read from index.html on disk. Editing that file redraws the
// open window within the poll interval, and the counter keeps its value.
// gpui opens the window. This package does not.
package reload

import "github.com/chinmay-sawant/go-gpui"

const (
	// SourcePath is the file the example window watches.
	SourcePath = "examples/reload/index.html"

	// DefaultWidth and DefaultHeight are the size of a newly opened window.
	DefaultWidth  = 640
	DefaultHeight = 560
)

// View is the data the template prints.
type View struct {
	Count int
}

// App is the reload screen.
type App struct {
	page *gpui.Page
	view View
}
