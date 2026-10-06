// Package controls is the form-controls example.
// The screen is an HTML template. A click focuses a text field, toggles a
// checkbox, checks a radio, cycles a select, or opens the file dialog for
// the doc field. ownframe opens the window. This package does not.
package controls

import (
	_ "embed"

	"github.com/chinmay-sawant/ownframe"
)

//go:embed controls.html
var controlsHTML string

const (
	// DefaultWidth and DefaultHeight are the size of a newly opened window.
	DefaultWidth  = 640
	DefaultHeight = 640
)

// View is the status line the template prints.
type View struct {
	Status string
}

// App is the controls screen.
type App struct {
	page *ownframe.Page
	view View
}
