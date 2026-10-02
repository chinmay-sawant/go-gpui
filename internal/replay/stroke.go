package replay

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"

	"github.com/chinmay-sawant/go-gpui/internal/render"
)

// drawStrokeRect strokes one complete rounded rectangle. v1 covers full
// strokes with circular corners; masked and elliptical ops are skipped,
// because render.Replayable keeps those pages on the bitmap path.
func drawStrokeRect(dst *ebiten.Image, op *layout.DisplayOp, dx, dy float64) {
	path, width, ok := strokeGeometry(op, dx, dy)
	if !ok {
		return
	}

	stroke := &vector.StrokeOptions{Width: width, LineJoin: vector.LineJoinRound}
	options := &vector.DrawPathOptions{AntiAlias: true}
	options.ColorScale.ScaleWithColor(rgba(op))

	vector.StrokePath(dst, &path, stroke, options)
}

// strokeGeometry builds the centerline path of a full OpStrokeRect and its
// width in canvas pixels, resolving circular corners like a rounded fill. A
// masked stroke or an elliptical corner returns ok false.
func strokeGeometry(op *layout.DisplayOp, dx, dy float64) (path vector.Path, width float32, ok bool) {
	if op.StrokeMask != 0 {
		return vector.Path{}, 0, false
	}

	radii, circular := render.FillRadii(op)
	if !circular {
		return vector.Path{}, 0, false
	}

	x := float32(op.X*pxPerPt + dx)
	y := float32(op.Y*pxPerPt + dy)
	w := float32(op.W * pxPerPt)
	h := float32(op.H * pxPerPt)
	r := [4]float32{
		float32(radii[0] * pxPerPt), float32(radii[1] * pxPerPt),
		float32(radii[2] * pxPerPt), float32(radii[3] * pxPerPt),
	}

	path.MoveTo(x+r[0], y)
	path.LineTo(x+w-r[1], y)
	arcTo(&path, x+w, y, x+w, y+r[1], r[1])
	path.LineTo(x+w, y+h-r[2])
	arcTo(&path, x+w, y+h, x+w-r[2], y+h, r[2])
	path.LineTo(x+r[3], y+h)
	arcTo(&path, x, y+h, x, y+h-r[3], r[3])
	path.LineTo(x, y+r[0])
	arcTo(&path, x, y, x+r[0], y, r[0])
	path.Close()

	width = float32(op.Width * pxPerPt)
	if width < 1 {
		width = 1
	}

	return path, width, true
}
