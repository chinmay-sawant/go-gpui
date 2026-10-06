package replay

import (
	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
	"github.com/chinmay-sawant/ownframe/internal/render"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"math"
)

// roundedPath traces the border centerline with per-corner elliptical arcs.
func roundedPath(op *layout.DisplayOp, dx, dy float64) vector.Path {
	rx, ry := render.RadiiXY(op)

	x := float32(op.X*pxPerPt + dx)
	y := float32(op.Y*pxPerPt + dy)
	w := float32(op.W * pxPerPt)
	h := float32(op.H * pxPerPt)
	rx4 := [4]float32{
		float32(rx[0] * pxPerPt), float32(rx[1] * pxPerPt),
		float32(rx[2] * pxPerPt), float32(rx[3] * pxPerPt),
	}
	ry4 := [4]float32{
		float32(ry[0] * pxPerPt), float32(ry[1] * pxPerPt),
		float32(ry[2] * pxPerPt), float32(ry[3] * pxPerPt),
	}

	var path vector.Path

	path.MoveTo(x+rx4[0], y)
	path.LineTo(x+w-rx4[1], y)
	ellipseTo(&path, x+w-rx4[1], y+ry4[1], rx4[1], ry4[1], -math.Pi/2, 0)
	path.LineTo(x+w, y+h-ry4[2])
	ellipseTo(&path, x+w-rx4[2], y+h-ry4[2], rx4[2], ry4[2], 0, math.Pi/2)
	path.LineTo(x+rx4[3], y+h)
	ellipseTo(&path, x+rx4[3], y+h-ry4[3], rx4[3], ry4[3], math.Pi/2, math.Pi)
	path.LineTo(x, y+ry4[0])
	ellipseTo(&path, x+rx4[0], y+ry4[0], rx4[0], ry4[0], math.Pi, 3*math.Pi/2)
	path.Close()

	return path
}

// ellipseTo appends an elliptical arc as short segments. The path must
// already sit on the arc's start point.
func ellipseTo(path *vector.Path, cx, cy, rx, ry, from, to float32) {
	const steps = 8

	for i := 1; i <= steps; i++ {
		a := from + (to-from)*float32(i)/steps
		path.LineTo(cx+rx*float32(math.Cos(float64(a))), cy+ry*float32(math.Sin(float64(a))))
	}
}
