package telegram_test

import (
	"context"
	"image"
	"testing"
)

// TestPinnedBarsRender proves the top bar and composer are pixel-identical
// at the viewport edges for offsets 0, 300, and 1200.
func TestPinnedBarsRender(t *testing.T) {
	ctx := context.Background()
	app := newShotApp(t, ctx)

	w, viewH := app.Page().Size()
	view := app.View()

	ref := shot(t, app, ctx, 0, 0)
	refBar := crop(ref, image.Rect(0, view.InsetTop, w, view.InsetTop+60))
	refCompose := crop(ref, image.Rect(0, viewH-view.InsetBottom-63, w, viewH-view.InsetBottom))

	for _, off := range []int{0, 300, 1200} {
		img := shot(t, app, ctx, off, off)
		barY := view.InsetTop
		composeY := viewH - view.InsetBottom - 63

		bar := crop(img, image.Rect(0, barY, w, barY+60))
		compose := crop(img, image.Rect(0, composeY, w, composeY+63))

		if !nearPix(bar, refBar, 40, 1000) {
			t.Errorf("header band differs at offset %d", off)
		}

		if !nearPix(compose, refCompose, 40, 1000) {
			t.Errorf("composer band differs at offset %d", off)
		}
	}
}
