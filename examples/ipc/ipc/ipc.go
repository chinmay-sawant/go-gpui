// Package ipc is the in-process IPC example.
// The screen sends strings with gpui.Send and gpui.Request to a listener and
// a handler registered in New, all inside this one process.
// gpui opens the window. This package does not.
package ipc

import (
	_ "embed"

	"github.com/chinmay-sawant/go-gpui"
)

//go:embed ipc.html
var ipcHTML string

const (
	// DefaultWidth and DefaultHeight are the size of a newly opened window.
	DefaultWidth  = 640
	DefaultHeight = 520
)

// View is the log and status the template prints.
type View struct {
	Log    string
	Status string
}

// App is the IPC screen. The stop fields cancel the registrations from New.
type App struct {
	page       *gpui.Page
	view       View
	stopLog    func()
	stopDouble func()
}
