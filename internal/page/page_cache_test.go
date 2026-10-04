package page_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/page"
)

// newCachePage builds a page for the cache tests.
func newCachePage(t *testing.T, html string) *page.Page {
	t.Helper()

	p, err := page.New(page.Config{HTML: html, Width: 320, Height: 200})
	if err != nil {
		t.Fatal(err)
	}

	return p
}

func TestResizeReusesParseAndCascade(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	p := newCachePage(t, `<p>hello cache</p>`)

	for i := 0; i < 10; i++ {
		p.SetSize(320+i, 200+i)
		if err := p.Redraw(ctx); err != nil {
			t.Fatal(err)
		}
	}

	stats := p.Stats()
	if stats.Parses != 1 || stats.Cascades != 1 {
		t.Fatalf("parses = %d, cascades = %d, want 1 and 1", stats.Parses, stats.Cascades)
	}
	if stats.Redraws != 10 || stats.Repaints != 10 || stats.Layouts != 10 {
		t.Fatalf("redraws = %d, repaints = %d, layouts = %d, want 10 each",
			stats.Redraws, stats.Repaints, stats.Layouts)
	}
	if stats.LastRedraw <= 0 || stats.LastDraw <= 0 {
		t.Fatalf("last redraw = %v, last draw = %v", stats.LastRedraw, stats.LastDraw)
	}
}

func TestHoverRelayoutKeepsCache(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	p := newCachePage(t, `<style>#a:hover{color:#ff0000}</style><p id="a">hover</p>`)
	if err := p.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	x, y := boxCenter(t, p, "a")
	if err := p.Hover(ctx, x, y); err != nil {
		t.Fatal(err)
	}

	stats := p.Stats()
	if stats.Parses != 1 || stats.Cascades != 1 {
		t.Fatalf("parses = %d, cascades = %d, want 1 and 1", stats.Parses, stats.Cascades)
	}
	if stats.Layouts != 2 {
		t.Fatalf("layouts = %d, want 2", stats.Layouts)
	}
}
