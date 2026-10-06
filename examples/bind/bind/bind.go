// Package bind is the binding example.
// The screen is an HTML template. Each data-bind control writes its value
// into View, and the template prints the struct on the next draw.
// ownframe opens the window. This package does not.
package bind

import (
	_ "embed"

	"github.com/chinmay-sawant/ownframe"
)

//go:embed bind.html
var bindHTML string

const (
	// DefaultWidth and DefaultHeight are the size of a newly opened window.
	DefaultWidth  = 640
	DefaultHeight = 640
)

// View is the data the template prints and the bound controls write.
type View struct {
	Name   string
	Email  string
	Agree  bool
	Plan   string
	Color  string
	Status string
}

// App is the binding screen.
type App struct {
	page *ownframe.Page
	view View
}
