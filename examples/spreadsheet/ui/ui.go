// Package ui is the spreadsheet example screen. It renders one bounded
// window of the active sheet on an ownframe page, keeps the selection, the
// active editor, and the cell tiles in its own state, and sends edits to a
// Backend. The core workbook packages implement Backend; tests use a fake.
package ui

import (
	"context"
	"errors"
	"time"

	"github.com/chinmay-sawant/ownframe"
)

// Options configures a new screen.
type Options struct {
	Backend Backend
	Title   string
	Width   int
	Height  int
	DataDir string
	Perf    bool
}

// App is one spreadsheet screen.
type App struct {
	page *ownframe.Page
	b    Backend
	opts Options

	sheets []Sheet
	info   map[string]Sheet
	active string
	rev    uint64

	sel   map[string]Selection
	used  map[string]Area
	edit  *editing
	tiles *tileCache
	win   Window

	fetchWin Window
	fetchGen uint64

	dark   bool
	status string
	csv    csvState

	work *worker

	viewW, viewH     int
	scrollX, scrollY int

	view View

	shift, ctrl, alt bool

	clickAt            time.Time
	clickedR, clickedC int

	closed bool
}

// Page returns the ownframe page Run displays.
func (a *App) Page() *ownframe.Page { return a.page }

// View returns the last built template data.
func (a *App) View() View { return a.view }

// selection returns the active sheet's selection, clamped to the sheet.
func (a *App) selection() Selection {
	sh := a.sheet()
	s, ok := a.sel[a.active]
	if !ok {
		s = newSelection(0, 0)
	}

	return s.Clamp(max(sh.Rows, 1), max(sh.Cols, 1))
}

// Redraw rebuilds the view and draws the page.
func (a *App) Redraw(ctx context.Context) error {
	a.refresh()

	return a.page.Redraw(ctx)
}

// refresh rebuilds the template data from the current state.
func (a *App) refresh() {
	a.viewW, a.viewH = a.page.Size()
	a.view = a.buildView()
}

var errNilBackend = errors.New("ui: nil backend")
