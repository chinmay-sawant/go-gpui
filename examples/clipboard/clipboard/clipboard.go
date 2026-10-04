// Package clipboard is the clipboard example.
// The buttons and the Ctrl+C/X/V/A/Z/Y chords call the page clipboard API.
// The library keeps the field values, and the OS clipboard holds the text
// the chords and the Copy, Cut, and Paste buttons use. gpui opens the
// window. This package does not.
package clipboard

import (
	_ "embed"

	"github.com/chinmay-sawant/go-gpui"
)

//go:embed clipboard.html
var clipboardHTML string

const (
	// DefaultWidth and DefaultHeight are the size of a newly opened window.
	DefaultWidth  = 560
	DefaultHeight = 640

	// pasteFromGo is the text the paste button inserts.
	pasteFromGo = "from Go"

	undoLimit = 64
)

// View is the status line the template prints.
type View struct {
	Status string
}

// App is the clipboard screen.
// last is the field a button acts on after the click blurred the form.
type App struct {
	page *gpui.Page
	view View
	last string
	undo []snap
	redo []snap
}
