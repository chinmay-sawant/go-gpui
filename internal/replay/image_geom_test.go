package replay

import (
	"context"
	"math"
	"testing"

	"github.com/chinmay-sawant/blinkless/layout"

	"github.com/chinmay-sawant/ownframe/internal/render"
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

func TestImageGeomAppliesTransform(t *testing.T) {
	source := `<html><body><div style="width:20px;height:20px;` +
		`background:linear-gradient(red,blue);transform:rotate(20deg)"></div></body></html>`

	display, err := render.DisplayList(context.Background(), source, 320, 200)
	if err != nil {
		t.Fatal(err)
	}

	for _, op := range display.Ops {
		if op.Kind != layout.DisplayOpImage || !op.XformSet {
			continue
		}

		geom := imageGeom(&op, 0, 0, 10, 10)
		_, y0 := geom.Apply(0, 0)
		_, y1 := geom.Apply(10, 0)
		if math.Abs(y1-y0) < 1e-6 {
			t.Fatal("rotation did not tilt the top edge")
		}

		return
	}

	t.Fatal("no transformed image op")
}

func assertNear(t *testing.T, got, want float64) {
	t.Helper()

	if math.Abs(got-want) > 1e-9 {
		t.Fatalf("got %v, want %v", got, want)
	}
}
