// Package host is the window and the browser page.
// Both talk to a screen. The program supplies that screen.
package host

import (
	"context"
	"image"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

// Ticker is a screen that wants one call per frame. The window calls Tick
// before it draws the frame. A screen without Tick is skipped.
type Ticker interface {
	Tick(ctx context.Context) error
}

// Screen is one HTML page the hosts can show.
// The page owns the template, the data, and what a click or a key does.
type Screen interface {
	Title() string
	Size() (int, int)
	MinSize() (int, int)
	Clamp(width, height int) (int, int)
	SetSize(width, height int)
	Redraw(ctx context.Context) error
	Image() image.Image
	Display() *layout.Display
	PNG() []byte
	Generation() uint64
	Boxes() []layout.Box
	Click(ctx context.Context, x, y float64) error
	Hover(ctx context.Context, x, y float64) error
	Press(ctx context.Context, x, y float64) error
	Release(ctx context.Context) error
	Type(ctx context.Context, text string) error
	Backspace(ctx context.Context) error
	DeleteWord(ctx context.Context) error
	Submit(ctx context.Context) error
	Copy(ctx context.Context) (string, bool, error)
	Cut(ctx context.Context) (string, bool, error)
	Paste(ctx context.Context, text string) error
	SelectAll(ctx context.Context) error
	Undo(ctx context.Context) error
	Redo(ctx context.Context) error
}
