package page_test

import (
	"context"
	"testing"
)

// TestSetImageKeepsCache checks that a new image resolver does not force a
// parse or a cascade: only the layout that reads the resolver changes.
func TestSetImageKeepsCache(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	p := newCachePage(t, backgroundHTML)

	if err := p.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	p.SetImage("bg", redPNG(t))
	if err := p.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	stats := p.Stats()
	if stats.Parses != 1 || stats.Cascades != 1 {
		t.Fatalf("parses = %d, cascades = %d, want 1 and 1", stats.Parses, stats.Cascades)
	}
}
