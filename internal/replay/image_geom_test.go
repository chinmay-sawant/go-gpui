package replay

import (
	"math"
	"testing"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

func TestImageGeomScalesAndTranslates(t *testing.T) {
	op := layout.DisplayOp{X: 10, Y: 20, W: 30, H: 40}
	geom := imageGeom(&op, 5, -3, 15, 20)

	x, y := geom.Apply(0, 0)
	assertNear(t, x, op.X*pxPerPt+5)
	assertNear(t, y, op.Y*pxPerPt-3)

	x, y = geom.Apply(15, 20)
	assertNear(t, x, (op.X+op.W)*pxPerPt+5)
	assertNear(t, y, (op.Y+op.H)*pxPerPt-3)
}

func TestImageGeomKeepsTranslateWhenBoundsMissing(t *testing.T) {
	op := layout.DisplayOp{X: 4, Y: 6, W: 30, H: 40}
	geom := imageGeom(&op, 1, 2, 0, 0)

	x, y := geom.Apply(15, 20)
	assertNear(t, x, 15+op.X*pxPerPt+1)
	assertNear(t, y, 20+op.Y*pxPerPt+2)
}

func assertNear(t *testing.T, got, want float64) {
	t.Helper()

	if math.Abs(got-want) > 1e-9 {
		t.Fatalf("got %v, want %v", got, want)
	}
}
