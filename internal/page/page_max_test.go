package page_test

import (
	"testing"

	"github.com/chinmay-sawant/ownframe/internal/page"
)

// TestNewLeavesTheMaxUnbounded checks a page with no max in Config asks the
// host for no cap, so the window keeps its OS maximize control.
func TestNewLeavesTheMaxUnbounded(t *testing.T) {
	t.Parallel()

	screen, err := page.New(page.Config{
		HTML:   "<p>Hi</p>",
		Width:  480,
		Height: 640,
	})
	if err != nil {
		t.Fatal(err)
	}

	maxW, maxH := screen.MaxSize()
	if maxW != 0 || maxH != 0 {
		t.Fatalf("max = %d x %d, want no cap", maxW, maxH)
	}
}
