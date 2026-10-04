// Package print is the PDF export example. The page is a small quarterly
// report. Save PDF writes a file, Print hands one to the OS print path, and
// -web serves the same bytes at GET /pdf.
package print

import (
	_ "embed"

	"github.com/chinmay-sawant/go-gpui"
)

//go:embed report.html
var reportHTML string

const (
	// DefaultWidth and DefaultHeight are the first frame size.
	DefaultWidth  = 720
	DefaultHeight = 640
)

// View is what the template prints.
type View struct {
	Title  string
	Status string
	Web    bool
}

// App is the report screen.
type App struct {
	page *gpui.Page
	view View
	// SavePath is where the Save PDF button writes. New puts it in the
	// user home directory; a test points it at a temporary directory.
	SavePath string
}

// Page returns the page the window or the server shows.
func (a *App) Page() *gpui.Page {
	return a.page
}

// View returns the data the template prints.
func (a *App) View() View {
	return a.view
}
