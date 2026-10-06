// Package web is the web mode example.
// The screen is an HTML template. Clicks and keystrokes are Go functions,
// and ownframe.Serve exposes the same page over HTTP. ownframe opens the window.
// This package does not.
package web

import (
	_ "embed"

	"github.com/chinmay-sawant/ownframe"
)

//go:embed web.html
var webHTML string

const (
	// DefaultWidth and DefaultHeight are the size of a newly opened window.
	DefaultWidth  = 640
	DefaultHeight = 480
)

// View is the counter and note the template prints.
type View struct {
	Count int
	Note  string
}

// App is the web mode screen.
type App struct {
	page *ownframe.Page
	view View
}
