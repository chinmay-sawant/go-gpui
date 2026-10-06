// Package scrolling is the mouse-wheel scrolling example.
// The screen is a column of forty rows in a frame shorter than the content.
// The window scrolls the page and drags the scrollbar thumbs; this package
// only lays the template out. ownframe opens the window. This package does not.
package scrolling

import (
	_ "embed"

	"github.com/chinmay-sawant/ownframe"
)

//go:embed scrolling.html
var scrollingHTML string

const (
	// DefaultWidth and DefaultHeight are the size of a newly opened window.
	// The column is taller than DefaultHeight, so the page scrolls.
	DefaultWidth  = 360
	DefaultHeight = 480
)

// App is the scrolling screen.
type App struct {
	page *ownframe.Page
}
