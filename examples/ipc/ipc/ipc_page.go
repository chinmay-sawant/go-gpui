package ipc

import (
	"context"

	"github.com/chinmay-sawant/go-gpui"
)

// Page returns the gpui page Run and Serve display.
func (a *App) Page() *gpui.Page {
	return a.page
}

// View returns the log and status lines.
func (a *App) View() View {
	return a.view
}

// Cancel removes the demo.log listener and the demo.double handler.
// A later Send or Request on those channels finds nothing.
func (a *App) Cancel() {
	if a.stopLog != nil {
		a.stopLog()
		a.stopLog = nil
	}

	if a.stopDouble != nil {
		a.stopDouble()
		a.stopDouble = nil
	}
}

// Redraw fills the template and renders the current size.
func (a *App) Redraw(ctx context.Context) error {
	return a.page.Redraw(ctx)
}

// Boxes returns the last hit-test boxes.
func (a *App) Boxes() []gpui.Box {
	return a.page.Boxes()
}

// Click hit-tests the page and runs the button under the point.
func (a *App) Click(ctx context.Context, x, y float64) error {
	return a.page.Click(ctx, x, y)
}

// Type edits the focused control. The library owns the field text.
func (a *App) Type(ctx context.Context, text string) error {
	return a.page.Type(ctx, text)
}
