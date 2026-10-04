// Package resize is the window resize example.
// The page has two columns, a media query that switches them at 640 px, a
// 100vw bar, a paragraph that rewraps, and a hover control. Dragging the
// window edge lays the page out again at the new size.
// gpui opens the window. This package does not.
package resize

import (
	_ "embed"

	"github.com/chinmay-sawant/go-gpui"
)

//go:embed resize.html
var resizeHTML string

const (
	// DefaultWidth and DefaultHeight are the size of a newly opened window.
	DefaultWidth  = 760
	DefaultHeight = 560

	// MinWidth and MinHeight are the smallest frame the screen will draw.
	// Below 300 px the columns have nothing left to compress and text
	// collapses, so the window refuses to shrink past it. The content still
	// scrolls when it is taller than the frame.
	MinWidth  = 300
	MinHeight = 160
)

// App is the resize screen.
type App struct {
	page *gpui.Page
}
