package page_test

import (
	"context"
	"testing"
)

func TestMediaOrientationFollowsTheFrame(t *testing.T) {
	t.Parallel()

	source := `<html><head><style>body{margin:0}` +
		`#a{width:40px}` +
		`@media (orientation: landscape){#a{width:120px}}` +
		`</style></head><body><div id="a"></div></body></html>`

	p := newViewportPage(t, source, 800, 600)
	if got := viewportBox(t, p, "a").W; got != 120 {
		t.Fatalf("landscape width = %.0f, want 120", got)
	}

	p.SetSize(400, 600)
	if err := p.Redraw(context.Background()); err != nil {
		t.Fatal(err)
	}

	if got := viewportBox(t, p, "a").W; got != 40 {
		t.Fatalf("portrait width = %.0f, want 40", got)
	}
}
