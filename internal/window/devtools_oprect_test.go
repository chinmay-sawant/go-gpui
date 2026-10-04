package window

import (
	"testing"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

// TestDevOpRectTextLineBox checks a text op's outline starts above its
// baseline, one line box tall.
func TestDevOpRectTextLineBox(t *testing.T) {
	t.Parallel()

	op := layout.DisplayOp{Kind: layout.DisplayOpText, X: 10, Y: 30, W: 80, H: 16, InkDescent: 4}

	got := devOpRect(&op, 1)
	if got.Y != 18 || got.H != 16 {
		t.Fatalf("text rect = %+v, want the line box at y 18", got)
	}
}
