// Package devtools is the devtools overlay example.
// The page carries one of every operation kind the replay draws: a rounded
// fill, a rounded border, a straight rule, shaped text, an image through
// SetImage, and a table row. The overlay starts on through Config.DevTools.
// ownframe opens the window. This package does not.
package devtools

import (
	_ "embed"

	"github.com/chinmay-sawant/ownframe"
)

//go:embed devtools.html
var devtoolsHTML string

const (
	// DefaultWidth and DefaultHeight are the size of a newly opened window.
	DefaultWidth  = 720
	DefaultHeight = 560
)

// View is the counter the template prints.
type View struct {
	Count int
}

// App is the devtools example screen.
type App struct {
	page *ownframe.Page
	view View
}
