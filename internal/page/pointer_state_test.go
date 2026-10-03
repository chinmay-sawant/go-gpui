package page

import (
	"context"
	"testing"
)

func hoverWidth(t *testing.T, p *Page) float64 {
	t.Helper()

	for _, b := range p.Boxes() {
		if b.ID == "e" {
			return b.W
		}
	}

	t.Fatal("no box e")

	return 0
}

func TestHoverPressRelease(t *testing.T) {
	t.Parallel()

	source := `<html><head><style>` +
		`#e{width:100px;height:20px}` +
		`#e:hover{width:120px}` +
		`#e:active{width:140px}` +
		`</style></head><body><div id="e"></div></body></html>`

	p, err := New(Config{HTML: source, Width: 320, Height: 200})
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	if err := p.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	if got := hoverWidth(t, p); got != 100 {
		t.Fatalf("plain width = %v", got)
	}

	if err := p.Hover(ctx, 10, 10); err != nil {
		t.Fatal(err)
	}

	if got := hoverWidth(t, p); got != 120 {
		t.Fatalf("hover width = %v", got)
	}

	if err := p.Press(ctx, 10, 10); err != nil {
		t.Fatal(err)
	}

	if got := hoverWidth(t, p); got != 140 {
		t.Fatalf("active width = %v", got)
	}

	if err := p.Release(ctx); err != nil {
		t.Fatal(err)
	}

	if got := hoverWidth(t, p); got != 120 {
		t.Fatalf("released width = %v", got)
	}

	if err := p.Hover(ctx, 300, 150); err != nil {
		t.Fatal(err)
	}

	if got := hoverWidth(t, p); got != 100 {
		t.Fatalf("unhovered width = %v", got)
	}
}
