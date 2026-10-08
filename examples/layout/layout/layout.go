// Package layout is the layout example.
// The screen is an HTML template. blinkless parses it, applies the CSS,
// and places every element; the hit-test boxes carry each element's id, tag,
// and geometry. ownframe opens the window. This package does not.
package layout

import (
	_ "embed"

	"github.com/chinmay-sawant/ownframe"
)

//go:embed layout.html
var layoutHTML string

const (
	// DefaultWidth and DefaultHeight are the size of a newly opened window.
	DefaultWidth  = 720
	DefaultHeight = 600
)

// View is the status line the template prints.
type View struct {
	Status string
}

// App is the layout screen.
type App struct {
	page *ownframe.Page
	view View
}
