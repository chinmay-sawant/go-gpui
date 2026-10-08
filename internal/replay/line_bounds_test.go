package replay

import (
	"image"
	"testing"

	"github.com/chinmay-sawant/blinkless/layout"
)

// TestLineBoundsCoversDiagonal pins the stroked box of a diagonal segment:
// the union of both endpoints, each grown by half the stroke, so the far
// x+w and y+h extents stay inside. The old axis-aligned branch kept only
// the near endpoint's column.
func TestLineBoundsCoversDiagonal(t *testing.T) {
	t.Parallel()

	op := &layout.DisplayOp{
		Kind: layout.DisplayOpLine,
		X:    10, Y: 20, W: 60, H: 40, Width: 4,
	}

	padded, ok := opBounds(op, 1)
	if !ok || padded != image.Rect(7, 17, 73, 63) {
		t.Fatalf("padded = %v, %v", padded, ok)
	}

	got, ok := PaintBounds(op, 1)
	if !ok {
		t.Fatal("diagonal unbounded")
	}

	if got.Min.X > 8 || got.Min.Y > 18 || got.Max.X < 72 || got.Max.Y < 62 {
		t.Fatalf("diagonal bounds = %v, want far extents 72, 62", got)
	}
}
