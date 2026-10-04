package window

import (
	"bytes"
	"context"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/chinmay-sawant/go-gpui/internal/page"
)

// TestDevToolsLeavesPageAlone is the plan's guard: the same page with the
// overlay off and on keeps its PNG bytes and generation, and the overlay
// never enters the display list or the box list.
func TestDevToolsLeavesPageAlone(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	p, err := page.New(page.Config{
		HTML:   `<div id="a" style="background:#1a56db;border-radius:8px">hello</div>`,
		Width:  320,
		Height: 200,
	})
	if err != nil {
		t.Fatal(err)
	}

	if err := p.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	display := p.Display()
	if display == nil {
		t.Fatal("the page fell back to a bitmap")
	}

	before := p.PNG()
	if len(before) == 0 {
		t.Fatal("no PNG before the overlay")
	}

	gen := p.Generation()
	ops := len(display.Ops)
	boxes := len(p.Boxes())

	s := &shell{app: p, ctx: ctx, screenW: 320, screenH: 200}

	off := ebiten.NewImage(320, 200)
	s.Draw(off)

	p.SetDevTools(true)
	if err := s.devSync(); err != nil {
		t.Fatal(err)
	}

	s.devRefresh()
	s.display = display
	s.dev.ops = true

	if boxes > 0 {
		s.dev.haveHov, s.dev.hovered = true, p.Boxes()[0]
		s.dev.havePin, s.dev.pinned = true, p.Boxes()[0]
	}

	on := ebiten.NewImage(320, 200)
	s.Draw(on)

	if p.Generation() != gen {
		t.Fatalf("generation = %d, want %d", p.Generation(), gen)
	}

	if !bytes.Equal(before, p.PNG()) {
		t.Fatal("the overlay changed the page PNG")
	}

	if len(p.Display().Ops) != ops || len(p.Boxes()) != boxes {
		t.Fatalf("ops = %d, boxes = %d, want %d and %d",
			len(p.Display().Ops), len(p.Boxes()), ops, boxes)
	}
}
