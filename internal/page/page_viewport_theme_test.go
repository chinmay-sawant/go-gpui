package page_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/page"
)

func TestThemeMediaQueryFollowsTheFrame(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	// #a is the template sheet's first rule, so the theme rule ties its
	// index and wins. A theme rule behind the template rule's index would
	// lose the tie whatever the media query says.
	source := `<html><head><style>` +
		`#a{width:60px;height:10px}` +
		`</style></head><body><div id="a"></div></body></html>`

	p, err := page.New(page.Config{
		HTML:  source,
		Theme: `@media (min-width: 700px){#a{width:150px}}`,
		Width: 800, Height: 600,
	})
	if err != nil {
		t.Fatal(err)
	}

	if err := p.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	if got := viewportBox(t, p, "a").W; got != 150 {
		t.Fatalf("wide theme width = %.0f, want 150", got)
	}

	p.SetSize(400, 600)
	if err := p.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	if got := viewportBox(t, p, "a").W; got != 60 {
		t.Fatalf("narrow theme width = %.0f, want 60", got)
	}
}
