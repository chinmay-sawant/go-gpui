// Package shapes is the shapes example.
// The gallery is one HTML template: a rounded fill, an elliptical stroke, a
// masked left border, a circle, letter-spaced text, and the same data-URI
// image three times. ownframe opens the window. This package does not.
package shapes

import (
	_ "embed"

	"github.com/chinmay-sawant/ownframe"
)

//go:embed shapes.html
var shapesHTML string

const (
	// DefaultWidth and DefaultHeight are the size of a newly opened window.
	DefaultWidth  = 720
	DefaultHeight = 640
)

// View is the status line the template prints.
type View struct {
	Status string
}

// App is the shapes screen.
type App struct {
	page *ownframe.Page
	view View
}
