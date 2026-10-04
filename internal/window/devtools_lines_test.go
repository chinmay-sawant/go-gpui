package window

import (
	"strings"
	"testing"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

// TestDevOpsLineCounts checks the ops panel line prints the per-kind counts
// and the total.
func TestDevOpsLineCounts(t *testing.T) {
	t.Parallel()

	d := newDevScreen()
	s := newDevShell(d)
	s.display = devTestDisplay()
	s.dev.ops = true

	line := s.devOpsLine()

	for _, want := range []string{"fill 1", "text 1", "total 2"} {
		if !strings.Contains(line, want) {
			t.Fatalf("ops line = %q, want %q", line, want)
		}
	}
}

// TestDevOpsLineFallback checks a fallback page says so where the counts
// would go.
func TestDevOpsLineFallback(t *testing.T) {
	t.Parallel()

	d := newDevScreen()
	s := newDevShell(d)
	s.dev.ops = true

	if line := s.devOpsLine(); !strings.Contains(line, "bitmap fallback") {
		t.Fatalf("ops line = %q, want bitmap fallback", line)
	}
}

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
