// Package bindhooks is the binding hooks example.
// The screen is an HTML template. Each data-bind control writes its value
// into View, BeforeEdit can veto an edit, and Change writes the status line.
// gpui opens the window. This package does not.
package bindhooks

import (
	_ "embed"

	"github.com/chinmay-sawant/go-gpui"
)

//go:embed bindhooks.html
var bindhooksHTML string

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
	Status string
}

// App is the binding-hooks screen.
type App struct {
	page *gpui.Page
	view View
}
