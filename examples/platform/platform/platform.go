// Package platform is the platform example.
// One page opens in a desktop window, a browser canvas, or a mobile view.
// Desktop and browser builds call Run; the mobile package calls BindMobile.
// ownframe opens the window. This package does not.
package platform

import (
	_ "embed"

	"github.com/chinmay-sawant/ownframe"
)

//go:embed platform.html
var platformHTML string

const (
	// DefaultWidth and DefaultHeight are the size of a newly opened window.
	DefaultWidth  = 640
	DefaultHeight = 480
)

// View is the click counter the template prints.
type View struct {
	Count int
}

// App is the platform screen.
type App struct {
	page *ownframe.Page
	view View
}
