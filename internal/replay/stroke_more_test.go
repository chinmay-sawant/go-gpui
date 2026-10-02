package replay

import (
	"testing"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

func TestStrokeGeometryRejectsElliptical(t *testing.T) {
	cases := []struct {
		name string
		op   layout.DisplayOp
	}{
		{
			name: "uniform",
			op: layout.DisplayOp{
				Kind: layout.DisplayOpStrokeRect,
				X:    0, Y: 0, W: 20, H: 10, Width: 1,
				Radius: 4, RadiusY: 2,
			},
		},
		{
			name: "one corner",
			op: layout.DisplayOp{
				Kind: layout.DisplayOpStrokeRect,
				X:    0, Y: 0, W: 20, H: 10, Width: 1,
				RadiusTopLeft: 4, RadiusTopLeftY: 2,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if _, _, ok := strokeGeometry(&tc.op, 0, 0); ok {
				t.Fatal("elliptical stroke accepted")
			}
		})
	}
}

func TestStrokeGeometryRejectsStrokeMask(t *testing.T) {
	op := &layout.DisplayOp{
		Kind: layout.DisplayOpStrokeRect,
		X:    0, Y: 0, W: 20, H: 10, Width: 1, Radius: 4,
		StrokeMask: 1,
	}

	if _, _, ok := strokeGeometry(op, 0, 0); ok {
		t.Fatal("masked stroke accepted")
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
