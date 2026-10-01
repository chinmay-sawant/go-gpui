// Package gpui shows an HTML page in a window.
//
// A program passes a template to New, stores the template data with SetData,
// and registers Handlers for clicks and keys. Run opens the window. Serve
// shows the same picture in a browser. BindMobile registers the page for an
// Android or iOS bind. Redraw parses the HTML, applies the CSS, and lays
// the page out through gowkhtmltopdf. This package has no layout of its own.
package gpui

import "github.com/chinmay-sawant/go-gpui/internal/page"

// Box is one element a click can land on.
// X, Y, W, and H are CSS pixels, origin at the top left of the picture.
type Box = page.Box

// Config is the template and the frame size for a new page.
// Width and Height are the first frame, in CSS pixels.
// MinWidth and MinHeight are the smallest frame. Zero means 1.
// MaxWidth and MaxHeight cap the picture so a large monitor does not
// allocate a matching PNG. Zero means 2560.
type Config = page.Config

// Handlers are the Go functions for input.
// A nil function ignores that input. After a function returns, the page
// is drawn again. Copy does not draw.
type Handlers = page.Handlers

// Page is one HTML template and the last picture it produced.
type Page = page.Page

var (
	// ErrEmptyHTML means New was given a blank template.
	ErrEmptyHTML = page.ErrEmptyHTML

	// ErrBadSize means a width or a height is unusable.
	ErrBadSize = page.ErrBadSize

	// ErrNilPage means Run, Serve, or BindMobile was called without a page.
	ErrNilPage = page.ErrNilPage
)

// New parses html as an html/template page.
func New(cfg Config) (*Page, error) {
	return page.New(cfg)
}
