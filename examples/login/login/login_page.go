package login

import (
	"context"

	"github.com/chinmay-sawant/go-gpui"
)

// Page returns the gpui page Run and Serve display.
func (a *App) Page() *gpui.Page {
	return a.page
}

// View returns a copy of the current login fields.
func (a *App) View() View {
	return a.view
}

// Size returns the frame size in CSS pixels.
func (a *App) Size() (int, int) {
	return a.page.Size()
}

// SetSize stores the frame size used by the next Redraw.
func (a *App) SetSize(width, height int) {
	a.page.SetSize(width, height)
}

// Generation increases by one on every successful Redraw.
func (a *App) Generation() uint64 {
	return a.page.Generation()
}

// Redraw fills the template and renders the current size.
func (a *App) Redraw(ctx context.Context) error {
	return a.page.Redraw(ctx)
}

// PNG returns the last PNG, or nil when nothing has been drawn.
func (a *App) PNG() []byte {
	return a.page.PNG()
}

// Boxes returns the last hit-test boxes.
func (a *App) Boxes() []gpui.Box {
	return a.page.Boxes()
}

// Click hit-tests the page and applies the login action.
func (a *App) Click(ctx context.Context, x, y float64) error {
	return a.page.Click(ctx, x, y)
}

// Type appends text to the focused field.
func (a *App) Type(ctx context.Context, text string) error {
	return a.page.Type(ctx, text)
}

// Backspace drops the last rune of the focused field.
func (a *App) Backspace(ctx context.Context) error {
	return a.page.Backspace(ctx)
}

// Submit checks the email and password.
func (a *App) Submit(ctx context.Context) error {
	return a.page.Submit(ctx)
}
