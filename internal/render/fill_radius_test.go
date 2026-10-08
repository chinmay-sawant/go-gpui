package render

import (
	"testing"

	"github.com/chinmay-sawant/blinkless/layout"
)

func TestFillRadiiResolvesCorners(t *testing.T) {
	t.Parallel()

	radii, ok := FillRadii(&layout.DisplayOp{Radius: 4})
	if !ok || radii != [4]float64{4, 4, 4, 4} {
		t.Fatalf("uniform radii = %v, ok = %v", radii, ok)
	}

	radii, ok = FillRadii(&layout.DisplayOp{
		RadiusTopLeft: 2, RadiusTopRight: 4,
		RadiusBottomRight: 6, RadiusBottomLeft: 8,
	})
	want := [4]float64{2, 4, 6, 8}

	if !ok || radii != want {
		t.Fatalf("corner radii = %v, ok = %v, want %v", radii, ok, want)
	}
}

func TestFillRadiiRejectsEllipticalCorners(t *testing.T) {
	t.Parallel()

	if _, ok := FillRadii(&layout.DisplayOp{Radius: 4, RadiusY: 2}); ok {
		t.Fatal("uniform ellipse accepted")
	}

	radii, ok := FillRadii(&layout.DisplayOp{Radius: 4, RadiusTopRightY: 2})
	if ok {
		t.Fatalf("corner ellipse accepted: %v", radii)
	}
}
