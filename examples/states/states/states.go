// Package states is the CSS states example.
// The screen is an HTML template. Hover, press, focus, and check controls
// match :hover, :active, :focus, and :checked, and the host redraws with
// that state. ownframe opens the window. This package does not.
package states

import (
	_ "embed"

	"github.com/chinmay-sawant/ownframe"
)

//go:embed states.html
var statesHTML string

const (
	// DefaultWidth and DefaultHeight are the size of a newly opened window.
	DefaultWidth  = 520
	DefaultHeight = 520
)

// View is the data the template prints.
type View struct {
	Status string
}

// App is the CSS states screen.
type App struct {
	page *ownframe.Page
	view View
}
