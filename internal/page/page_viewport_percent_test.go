package page_test

import (
	"context"
	"testing"
)

func TestPercentAndFullHeightFollowTheFrame(t *testing.T) {
	t.Parallel()

	source := `<html><head><style>body{margin:0}` +
		`#outer{width:50%;height:300px}` +
		`#inner{width:50%;height:100%}` +
		`</style></head><body><div id="outer"><div id="inner"></div></div></body></html>`

	p := newViewportPage(t, source, 800, 600)
	if box := viewportBox(t, p, "inner"); box.W != 200 || box.H != 300 {
		t.Fatalf("inner = %.0f x %.0f, want 200 x 300", box.W, box.H)
	}

	p.SetSize(400, 600)
	if err := p.Redraw(context.Background()); err != nil {
		t.Fatal(err)
	}

	if box := viewportBox(t, p, "inner"); box.W != 100 || box.H != 300 {
		t.Fatalf("inner after relayout = %.0f x %.0f, want 100 x 300", box.W, box.H)
	}
}
