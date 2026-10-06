// Package theme is the live theme example.
// The screen is an HTML template that reads custom properties with var().
// The extra stylesheet passed to ownframe.Config and Page.SetTheme defines those
// properties, so a click can restyle the page without touching the template.
// ownframe opens the window. This package does not.
package theme

import (
	_ "embed"

	"github.com/chinmay-sawant/ownframe"
)

//go:embed theme.html
var themeHTML string

const (
	// DefaultWidth and DefaultHeight are the size of a newly opened window.
	DefaultWidth  = 520
	DefaultHeight = 420
)

// View is the data the template prints.
type View struct {
	Title  string
	Status string
}

// App is the themed screen.
type App struct {
	page *ownframe.Page
	view View
	dark bool
}
