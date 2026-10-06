// Package editing is the text-editing example.
// The screen is an HTML template. Copy, cut, paste, select-all, undo, and
// redo edit the note, and the note survives a redraw. ownframe opens the window.
// This package does not.
package editing

import (
	_ "embed"

	"github.com/chinmay-sawant/ownframe"
)

//go:embed editing.html
var editingHTML string

const (
	// DefaultWidth and DefaultHeight are the size of a newly opened window.
	DefaultWidth  = 640
	DefaultHeight = 640

	undoLimit = 64
)

// View is the stamp and status the template prints.
type View struct {
	Stamp  int
	Status string
}

// App is the editing screen.
type App struct {
	page *ownframe.Page
	view View
	undo []string
	redo []string
}
