package page_test

import (
	"context"
	"testing"
)

func TestViewportUnitsFollowTheFrame(t *testing.T) {
	t.Parallel()

	source := `<html><head><style>body{margin:0}` +
		`#w{width:50vw;height:1px}` +
		`#vh{height:20vh} #dvh{height:20dvh}` +
		`#svh{height:20svh} #lvh{height:20lvh}` +
		`</style></head><body>` +
		`<div id="w"></div><div id="vh"></div><div id="dvh"></div>` +
		`<div id="svh"></div><div id="lvh"></div></body></html>`

	p := newViewportPage(t, source, 800, 600)
	if got := viewportBox(t, p, "w").W; got != 400 {
		t.Fatalf("50vw = %.0f, want 400", got)
	}

	for _, id := range []string{"vh", "dvh", "svh", "lvh"} {
		if got := viewportBox(t, p, id).H; got != 120 {
			t.Fatalf("%s = %.0f, want 120", id, got)
		}
	}
}

func TestPercentAndFullHeightFollowTheFrame(t *testing.T) {
	t.Parallel()

	source := `<html><head><style>body{margin:0}` +
		`#outer{width:200px;height:300px}` +
		`#inner{width:50%;height:100%}` +
		`</style></head><body><div id="outer"><div id="inner"></div></div></body></html>`

	p := newViewportPage(t, source, 800, 600)
	if box := viewportBox(t, p, "inner"); box.W != 100 || box.H != 300 {
		t.Fatalf("inner = %.0f x %.0f, want 100 x 300", box.W, box.H)
	}
}

func TestVarChainReadingAViewportUnitRecomputes(t *testing.T) {
	t.Parallel()

	source := `<html><head><style>body{margin:0}` +
		`:root{--a:var(--b);--b:30vw}` +
		`#a{width:var(--a);height:10px}` +
		`</style></head><body><div id="a"></div></body></html>`

	p := newViewportPage(t, source, 800, 600)
	if got := viewportBox(t, p, "a").W; got != 240 {
		t.Fatalf("wide var width = %.0f, want 240", got)
	}

	p.SetSize(400, 600)
	if err := p.Redraw(context.Background()); err != nil {
		t.Fatal(err)
	}

	if got := viewportBox(t, p, "a").W; got != 120 {
		t.Fatalf("narrow var width = %.0f, want 120", got)
	}
}
