package replay

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

// drawLine draws one border segment. The engine centers the stroke on the
// line and extends it half a stroke past each endpoint, so one filled
// rectangle reproduces an axis-aligned segment exactly. A diagonal segment,
// such as a checked checkbox's tick, is stroked instead, with square caps
// for the same extension; a filled rectangle would flatten it into a bar.
func drawLine(dst *ebiten.Image, op *layout.DisplayOp, dx, dy float64) {
	x, y, w, h, width := op.PaintLineGeometry()

	stroke := int(math.Round(width * pxPerPt))
	if stroke < 1 {
		stroke = 1
	}

	if w != 0 && h != 0 {
		drawSegment(dst, op, x, y, x+w, y+h, float32(stroke), dx, dy)

		return
	}

	half := float64(stroke) / 2 / pxPerPt

	var box image.Rectangle
	if op.H <= 0 {
		box = snap(x-half, y-half, w+2*half, 2*half)
	} else {
		box = snap(x-half, y-half, 2*half, h+2*half)
	}

	vector.FillRect(
		dst,
		float32(box.Min.X)+float32(dx), float32(box.Min.Y)+float32(dy),
		float32(box.Dx()), float32(box.Dy()), rgba(op), false,
	)
}

// snap converts a point-space rectangle to whole canvas pixels, matching the
// engine's own raster geometry.
func snap(x, y, w, h float64) image.Rectangle {
	return image.Rect(
		int(math.Round(x*pxPerPt)), int(math.Round(y*pxPerPt)),
		int(math.Round((x+w)*pxPerPt)), int(math.Round((y+h)*pxPerPt)),
	)
}
