// Package crash is the crash-report example.
// The report button writes a report without a panic. The panic button
// panics; ownframe.Run recovers, writes a report, and returns an error that
// names the file. ownframe opens the window. This package does not.
package crash

import (
	_ "embed"

	"github.com/chinmay-sawant/ownframe"
)

//go:embed crash.html
var crashHTML string

const (
	// DefaultWidth and DefaultHeight are the size of a newly opened window.
	DefaultWidth  = 560
	DefaultHeight = 480

	// reportTitle is the title of every report this example writes.
	reportTitle = "crash example"
)

// View is the status line the template prints.
type View struct {
	Status string
}

// App is the crash-report screen.
type App struct {
	page *ownframe.Page
	view View
}
