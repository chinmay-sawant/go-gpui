package page

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/host"
)

func TestCursorShapes(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	p := bftPage(t, `<input id="t" type="text"`+bedWide+`>`+
		`<button id="go">Go</button>`+
		`<a id="l" href="http://example.com/">link</a>`+
		`<div id="ord">plain</div>`)
	bftDraw(t, p, nil)

	steps := []struct {
		id   string
		want host.Shape
	}{
		{"t", host.ShapeText},
		{"go", host.ShapePointer},
		{"ord", host.ShapeDefault},
	}

	for _, step := range steps {
		b := p.boxByID(step.id)
		if err := p.Hover(ctx, b.X+b.W/2, b.Y+b.H/2); err != nil {
			t.Fatal(err)
		}

		if got := p.CursorShape(); got != step.want {
			t.Fatalf("%s shape %v want %v", step.id, got, step.want)
		}
	}

	link := runOp(p, "link")
	left, width := runSpan(p, link)
	if err := p.Hover(ctx, left+width/2, runBase(p, link)-4); err != nil {
		t.Fatal(err)
	}

	if got := p.CursorShape(); got != host.ShapePointer {
		t.Fatalf("link shape %v", got)
	}

	if err := p.Hover(ctx, 1, 1); err != nil {
		t.Fatal(err)
	}

	if got := p.CursorShape(); got != host.ShapeDefault {
		t.Fatalf("background shape %v", got)
	}
}
