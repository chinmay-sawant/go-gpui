// Package history is the navigation history example.
// The screen is an HTML template. Load, Back, and Forward move a history
// list, and data-action routes load a page registered with Route. ownframe
// opens the window. This package does not.
package history

import (
	_ "embed"

	"github.com/chinmay-sawant/ownframe"
)

//go:embed history.html
var historyHTML string

const (
	// DefaultWidth and DefaultHeight are the size of a newly opened window.
	DefaultWidth  = 520
	DefaultHeight = 520
)

// View is the data the template prints.
type View struct {
	Status string
}

// App is the history screen.
type App struct {
	page *ownframe.Page
	view View
}
