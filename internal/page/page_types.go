package page

import (
	"context"
	"html/template"
	"image"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

// Handlers are the Go functions for input.
// A nil function ignores that input. After a function returns, the page
// is drawn again. Copy does not draw.
type Handlers struct {
	Click      func(ctx context.Context, box Box) error
	Type       func(ctx context.Context, text string) error
	Backspace  func(ctx context.Context) error
	DeleteWord func(ctx context.Context) error
	Submit     func(ctx context.Context) error
	Copy       func(ctx context.Context) (text string, ok bool, err error)
	Cut        func(ctx context.Context) (text string, ok bool, err error)
	Paste      func(ctx context.Context, text string) error
	SelectAll  func(ctx context.Context) error
	Undo       func(ctx context.Context) error
	Redo       func(ctx context.Context) error
}

// Page is one HTML template and the last picture it produced.
type Page struct {
	title      string
	tpl        *template.Template
	data       any
	handlers   Handlers
	img        image.Image
	boxes      []layout.Box
	width      int
	height     int
	minWidth   int
	minHeight  int
	maxWidth   int
	maxHeight  int
	generation uint64
	past       []string
	pastAt     int
	routes     map[string]string
}
