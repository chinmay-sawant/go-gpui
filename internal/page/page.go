// Package page turns an HTML template into a picture and hit boxes.
package page

import (
	"errors"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

// Box is one element a click can land on.
// X, Y, W, H are CSS pixels, origin at the top left of the picture.
type Box = layout.Box

const defaultMax = 2560

var (
	// ErrEmptyHTML means New was given a blank template.
	ErrEmptyHTML = errors.New("gpui: empty html")

	// ErrBadSize means a width or a height is unusable.
	ErrBadSize = errors.New("gpui: bad size")

	// ErrNilPage means Run, Serve, or BindMobile was called without a page.
	ErrNilPage = errors.New("gpui: nil page")

	errNilContext = errors.New("gpui: nil context")

	// errNoImage is the resolver answer for a src SetImage does not hold.
	errNoImage = errors.New("gpui: no image for src")
)

// Config is the template, the optional theme, and the frame size for a new
// page. Theme is an extra stylesheet applied after the template's own styles;
// empty means no theme.
// Width and Height are the first frame, in CSS pixels.
// MinWidth and MinHeight are the smallest frame. Zero means 1.
// MaxWidth and MaxHeight ask the host to cap the window; zero means 2560.
// They do not cap the picture, so the layout follows the window.
type Config struct {
	Title string
	HTML  string
	// File reads the page source from disk at New and watches it. Setting
	// File and HTML together is an error.
	File  string
	Theme string
	// ThemeFile reads the theme stylesheet from disk and watches it.
	ThemeFile string
	// DisableHotReload turns the file watch off. It is ignored without File.
	DisableHotReload bool
	// DevTools starts the window overlay on.
	DevTools  bool
	Width     int
	Height    int
	MinWidth  int
	MinHeight int
	MaxWidth  int
	MaxHeight int
}
