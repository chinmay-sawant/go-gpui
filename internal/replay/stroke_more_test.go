package replay

import (
	"testing"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

func TestStrokeGeometryElliptical(t *testing.T) {
	op := &layout.DisplayOp{
		Kind: layout.DisplayOpStrokeRect,
		X:    0, Y: 0, W: 20, H: 10, Width: 1,
		Radius: 4, RadiusY: 2,
	}

	path, _, ok := strokeGeometry(op, 0, 0)
	if !ok {
		t.Fatal("elliptical stroke rejected")
	}

	if path.Bounds().Empty() {
		t.Fatal("elliptical path is empty")
	}
}

func TestStrokeGeometryMasked(t *testing.T) {
	op := &layout.DisplayOp{
		Kind: layout.DisplayOpStrokeRect,
		X:    0, Y: 0, W: 20, H: 10, Width: 1,
		StrokeMask: 1,
	}

	path, _, ok := strokeGeometry(op, 0, 0)
	if !ok {
		t.Fatal("masked stroke rejected")
	}

	if path.Bounds().Dx() <= 0 {
		t.Fatal("masked path has no length")
	}
}

func TestStrokeGeometryUnknownMask(t *testing.T) {
	op := &layout.DisplayOp{
		Kind: layout.DisplayOpStrokeRect,
		X:    0, Y: 0, W: 20, H: 10, Width: 1,
		StrokeMask: 16,
	}

	if _, _, ok := strokeGeometry(op, 0, 0); ok {
		t.Fatal("unknown mask accepted")
	}
}

func TestStrokeGeometryWidthConversion(t *testing.T) {
	cases := []struct {
		name  string
		width float64
		want  float32
	}{
		{name: "points scale", width: 1, want: float32(pxPerPt)},
		{name: "three points", width: 3, want: float32(3 * pxPerPt)},
		{name: "zero clamps to one", width: 0, want: 1},
		{name: "negative clamps to one", width: -2, want: 1},
		{name: "thin clamps to one", width: 0.5, want: 1},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			op := &layout.DisplayOp{
				Kind: layout.DisplayOpStrokeRect,
				X:    0, Y: 0, W: 20, H: 10, Width: tc.width,
			}

			_, width, ok := strokeGeometry(op, 0, 0)
			if !ok {
				t.Fatal("stroke rejected")
			}

			if width != tc.want {
				t.Fatalf("width = %v, want %v", width, tc.want)
			}
		})
	}
}
