// Package drop is the drag-and-drop example.
// A dropped PNG or JPEG shows through SetImage; a dropped text file shows its
// first lines. The window delivers the files. gpui opens the window. This
// package does not.
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

	// imageSrc is the source name the template's image points at.
	imageSrc = "dropped"

	// maxLines caps the preview of a dropped text file.
	maxLines = 12

	// noFile is the first status line.
	noFile = "No file yet."
)

// View is the status line and the preview the template prints.
type View struct {
	Status string
	Image  bool
	Lines  []string
}

// App is the drop screen.
type App struct {
	page *gpui.Page
	view View
}
