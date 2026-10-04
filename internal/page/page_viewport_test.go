package page_test

import (
	"context"
	"testing"
)

func TestMediaWidthFollowsTheFrame(t *testing.T) {
	t.Parallel()

	source := `<html><head><style>body{margin:0}` +
		`#a{width:40px;height:30px}` +
		`@media (min-width: 700px){#a{width:120px}}` +
		`</style></head><body><div id="a"></div></body></html>`

	p := newViewportPage(t, source, 800, 600)
	if got := viewportBox(t, p, "a").W; got != 120 {
		t.Fatalf("wide width = %.0f, want 120", got)
	}

	p.SetSize(400, 300)
	if err := p.Redraw(context.Background()); err != nil {
		t.Fatal(err)
	}

	if got := viewportBox(t, p, "a").W; got != 40 {
		t.Fatalf("narrow width = %.0f, want 40", got)
	}
}

func TestMediaHeightFollowsTheFrame(t *testing.T) {
	t.Parallel()

	// Engine gap: the media-feature parser accepts only width and
	// inline-size, so height, min-height, and max-height never match
	// (internal/css/container.go parseSizeFeature). Recorded in
	// plans/v0.0.2/dynamic-resize.md; remove the skip when the engine
	// parses height.
	t.Skip("engine gap: height media features do not parse")

	source := `<html><head><style>body{margin:0}` +
		`#a{height:30px}` +
		`@media (min-height: 500px){#a{height:90px}}` +
		`</style></head><body><div id="a"></div></body></html>`

	p := newViewportPage(t, source, 800, 600)
	if got := viewportBox(t, p, "a").H; got != 90 {
		t.Fatalf("tall height = %.0f, want 90", got)
	}

	p.SetSize(400, 400)
	if err := p.Redraw(context.Background()); err != nil {
		t.Fatal(err)
	}

	if got := viewportBox(t, p, "a").H; got != 30 {
		t.Fatalf("short height = %.0f, want 30", got)
	}
}
