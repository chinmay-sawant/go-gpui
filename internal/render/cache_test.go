package render_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe/internal/render"
)

const cacheSource = `<style>` +
	`#b{width:100px;height:20px;background:#3366cc}` +
	`@media (min-width:600px){#b{width:300px}}` +
	`</style><div id="b"></div>`

// TestCacheRelayoutFollowsViewport relayouts the cached document at a wider
// viewport and checks the media query and the fill width follow it.
func TestCacheRelayoutFollowsViewport(t *testing.T) {
	ctx := context.Background()

	cache, err := render.NewCache(ctx, cacheSource, 320, 200, render.State{})
	if err != nil {
		t.Fatal(err)
	}

	if cache.Source() != cacheSource {
		t.Fatalf("source = %q", cache.Source())
	}

	display, err := render.DisplayListDocument(ctx, cache.Styled(), nil)
	if err != nil {
		t.Fatal(err)
	}

	if got := fillWidth(display); got != 75 {
		t.Fatalf("narrow width = %v, want 75", got)
	}

	styled, err := cache.Relayout(ctx, 800, 200, render.State{})
	if err != nil {
		t.Fatal(err)
	}

	display, err = render.DisplayListDocument(ctx, styled, nil)
	if err != nil {
		t.Fatal(err)
	}

	if got := fillWidth(display); got != 225 {
		t.Fatalf("wide width = %v, want 225", got)
	}
}

// fillWidth returns the width of the first fill operation.
func fillWidth(display *render.Display) float64 {
	for i := range display.Ops {
		if display.Ops[i].Kind == render.OpFillRect {
			return display.Ops[i].W
		}
	}

	return 0
}
