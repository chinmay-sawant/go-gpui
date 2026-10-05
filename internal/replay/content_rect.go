package replay

import (
	"image"
	"math"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

// contentRect is the painted page in CSS pixels: the canvas plus any box
// that overflows it, the region the window's replay buffer covers.
func contentRect(display *layout.Display) image.Rectangle {
	r := canvasRect(display)
	if display == nil {
		return r
	}

	for _, b := range display.Boxes {
		r = r.Union(contentBox(b))
	}

	return r
}

// contentBox rounds one element box outward to whole CSS pixels.
func contentBox(b layout.Box) image.Rectangle {
	return image.Rect(
		int(math.Floor(b.X)), int(math.Floor(b.Y)),
		int(math.Ceil(b.X+b.W)), int(math.Ceil(b.Y+b.H)),
	)
}
