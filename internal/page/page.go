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

// Config is the template and the frame size for a new page.
// Width and Height are the first frame, in CSS pixels.
// MinWidth and MinHeight are the smallest frame. Zero means 1.
// MaxWidth and MaxHeight cap the picture so a large monitor does not
// allocate a matching PNG. Zero means 2560.
type Config struct {
	Title     string
	HTML      string
	Width     int
	Height    int
	MinWidth  int
	MinHeight int
	MaxWidth  int
	MaxHeight int
}
