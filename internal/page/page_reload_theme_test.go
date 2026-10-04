package page_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/page"
)

func newThemedFilePage(t *testing.T, path string) *page.Page {
	t.Helper()

	p := newFilePage(t, page.Config{
		HTML:      themePage,
		ThemeFile: path,
		Width:     320,
		Height:    200,
	})
	if err := p.Redraw(context.Background()); err != nil {
		t.Fatal(err)
	}

	return p
}

func TestPollReloadThemeFile(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	path := writeSource(t, "theme.css", `#swatch{background:#ff0000}`)
	p := newThemedFilePage(t, path)

	if r, g, b := swatchRGB(t, p); r < 1 || g > 0 || b > 0 {
		t.Fatalf("swatch = %v,%v,%v, want red", r, g, b)
	}

	gen := p.Generation()
	rewriteSource(t, path, `#swatch{background:#00ff00}`)

	changed, err := p.PollReload(ctx)
	if !changed || err != nil {
		t.Fatalf("poll = %v, %v", changed, err)
	}

	if r, g, b := swatchRGB(t, p); g < 1 || r > 0 || b > 0 {
		t.Fatalf("swatch = %v,%v,%v, want green", r, g, b)
	}

	if p.Generation() != gen+1 {
		t.Fatalf("generation moved by %d", p.Generation()-gen)
	}
}

func TestPollReloadThemeErrorKeepsOld(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	path := writeSource(t, "theme.css", `#swatch{background:#ff0000}`)
	p := newThemedFilePage(t, path)
	gen := p.Generation()

	rewriteSource(t, path, `}`)

	changed, err := p.PollReload(ctx)
	if changed || err == nil {
		t.Fatalf("bad theme poll = %v, %v, want a parse error", changed, err)
	}

	if r, g, b := swatchRGB(t, p); r < 1 || g > 0 || b > 0 {
		t.Fatalf("swatch = %v,%v,%v, want the old red", r, g, b)
	}

	if p.Generation() != gen {
		t.Fatal("a broken theme redrew the page")
	}
}
