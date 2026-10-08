package page_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/blinkless/layout"
	"github.com/chinmay-sawant/ownframe/internal/page"
)

const themePage = `<html><head><style>` +
	`#swatch{width:100px;height:40px;background:#ffffff}` +
	`</style></head><body><div id="swatch"></div></body></html>`

// swatchRGB returns the swatch fill channels, 0..1. The swatch is 100x40 CSS
// pixels, 75x30 points in the display list.
func swatchRGB(t *testing.T, p *page.Page) (float64, float64, float64) {
	t.Helper()

	d := p.Display()
	if d == nil {
		t.Fatal("theme page fell back to a bitmap")
	}

	for i := range d.Ops {
		op := &d.Ops[i]
		if op.Kind == layout.DisplayOpFillRect && op.W == 75 && op.H == 30 {
			return op.R, op.G, op.B
		}
	}

	t.Fatal("no swatch fill")

	return 0, 0, 0
}

// newThemePage builds the swatch page with theme CSS.
func newThemePage(t *testing.T, theme string) *page.Page {
	t.Helper()

	p, err := page.New(page.Config{HTML: themePage, Theme: theme, Width: 320, Height: 200})
	if err != nil {
		t.Fatal(err)
	}

	return p
}

func TestConfigThemeWinsTie(t *testing.T) {
	t.Parallel()

	p := newThemePage(t, `#swatch{background:#ff0000}`)
	if err := p.Redraw(context.Background()); err != nil {
		t.Fatal(err)
	}

	if r, g, b := swatchRGB(t, p); r < 1 || g > 0 || b > 0 {
		t.Fatalf("swatch = %v,%v,%v, want red", r, g, b)
	}
}

func TestSetThemeSwitchesAndClears(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	p := newThemePage(t, "")
	if err := p.SetTheme(`#swatch{background:#00ff00}`); err != nil {
		t.Fatal(err)
	}

	if err := p.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	if r, g, b := swatchRGB(t, p); g < 1 || r > 0 || b > 0 {
		t.Fatalf("green swatch = %v,%v,%v", r, g, b)
	}

	if err := p.SetTheme(""); err != nil {
		t.Fatal(err)
	}

	if err := p.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	if r, g, b := swatchRGB(t, p); r < 1 || g < 1 || b < 1 {
		t.Fatalf("cleared swatch = %v,%v,%v, want template white", r, g, b)
	}
}
