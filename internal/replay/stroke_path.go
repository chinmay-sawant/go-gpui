package replay

import (
	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// strokeGeometry builds the centerline path of an OpStrokeRect and its width
// in canvas pixels. A non-zero StrokeMask builds straight side segments; an
// unknown mask bit returns ok false. Zero builds a rounded path whose corners
// may be elliptical.
func strokeGeometry(op *layout.DisplayOp, dx, dy float64) (vector.Path, float32, bool) {
	width := float32(op.Width * pxPerPt)
	if width < 1 {
		width = 1
	}

	if op.StrokeMask != 0 {
		return maskPath(op, dx, dy), width, op.StrokeMask&^0x0F == 0
	}

	return roundedPath(op, dx, dy), width, true
}

// maskPath traces the selected sides as straight segments.
func maskPath(op *layout.DisplayOp, dx, dy float64) vector.Path {
	x := float32(op.X*pxPerPt + dx)
	y := float32(op.Y*pxPerPt + dy)
	w := float32(op.W * pxPerPt)
	h := float32(op.H * pxPerPt)

	var path vector.Path

	if op.StrokeMask&1 != 0 {
		path.MoveTo(x, y)
		path.LineTo(x+w, y)
	}
	if op.StrokeMask&2 != 0 {
		path.MoveTo(x+w, y)
		path.LineTo(x+w, y+h)
	}
	if op.StrokeMask&4 != 0 {
		path.MoveTo(x, y+h)
		path.LineTo(x+w, y+h)
	}
	if op.StrokeMask&8 != 0 {
		path.MoveTo(x, y)
		path.LineTo(x, y+h)
	}

	return path
}
