package page_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/page"
	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

// viewportBox returns the box with an id after a relayout.
func viewportBox(t *testing.T, p *page.Page, id string) layout.Box {
	t.Helper()

	for _, b := range p.Boxes() {
		if b.ID == id {
			return b
		}
	}

	t.Fatalf("no box id=%q", id)

	return layout.Box{}
}

// newViewportPage draws source at width x height.
func newViewportPage(t *testing.T, source string, width, height int) *page.Page {
	t.Helper()

	p, err := page.New(page.Config{HTML: source, Width: width, Height: height})
	if err != nil {
		t.Fatal(err)
	}

	if err := p.Redraw(context.Background()); err != nil {
		t.Fatal(err)
	}

	return p
}
