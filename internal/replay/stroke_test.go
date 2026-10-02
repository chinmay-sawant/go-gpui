package replay

import (
	"image"
	"testing"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

func TestStrokeGeometryZeroRadiiBounds(t *testing.T) {
	op := &layout.DisplayOp{
		Kind: layout.DisplayOpStrokeRect,
		X:    0, Y: 0, W: 30, H: 15, Width: 2,
	}

	path, width, ok := strokeGeometry(op, 3, 7)
	if !ok {
		t.Fatal("zero-radius stroke rejected")
	}

	want := image.Rect(3, 7, 43, 27)
	if got := path.Bounds(); got != want {
		t.Fatalf("bounds = %v, want %v", got, want)
	}

	if wantWidth := float32(2 * pxPerPt); width != wantWidth {
		t.Fatalf("width = %v, want %v", width, wantWidth)
	}
}

func TestStrokeGeometryCircularRadiusPath(t *testing.T) {
	op := &layout.DisplayOp{
		Kind: layout.DisplayOpStrokeRect,
		X:    10, Y: 5, W: 30, H: 15, Width: 1, Radius: 6,
	}

	path, _, ok := strokeGeometry(op, 0, 0)
	if !ok {
		t.Fatal("circular-radius stroke rejected")
	}

	bounds := path.Bounds()
	if bounds.Empty() {
		t.Fatal("circular-radius path is empty")
	}

	want := image.Rect(13, 6, 54, 27)
	if bounds != want {
		t.Fatalf("bounds = %v, want %v", bounds, want)
	}
}
