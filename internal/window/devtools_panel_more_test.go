package window

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/page"
)

// TestDevPanelMovesWithARedraw checks the Phase 5 exit: the panel reads the
// page stats, so a redraw moves the redraws row in the same frame.
func TestDevPanelMovesWithARedraw(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	p, err := page.New(page.Config{HTML: `<p id="a">one</p>`, Width: 200, Height: 100})
	if err != nil {
		t.Fatal(err)
	}

	if err = p.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	s := &shell{app: p, ctx: ctx}
	s.dev.stats = p.Stats()

	if got := devFrameValue(s.devFrameRows(320), "Redraws"); got != "1" {
		t.Fatalf("redraws = %q, want 1", got)
	}

	if err = p.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	s.dev.stats = p.Stats()

	if got := devFrameValue(s.devFrameRows(320), "Redraws"); got != "2" {
		t.Fatalf("redraws = %q, want 2", got)
	}
}
