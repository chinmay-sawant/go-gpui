// Package form is the form example.
// The screen is an HTML template. The library stores the control values.
// gpui opens the window. This package does not.
package form

import (
	_ "embed"

	"github.com/chinmay-sawant/go-gpui"
)

//go:embed form.html
var formHTML string

const (
	// DefaultWidth and DefaultHeight are the size of a newly opened window.
	DefaultWidth  = 640
	DefaultHeight = 640
)

// View is the status line the template prints.
type View struct {
	Status string
}

// App is the form screen.
type App struct {
	page *gpui.Page
	view View
}
