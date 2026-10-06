package telegram_test

import (
	"context"
	"image"
	"testing"
)

// TestBarsCoverMessages proves the bars replace content rather than floating
// over identical pixels: with the bars scrolled out, the same viewport rows
// differ from the pinned bands.
func TestBarsCoverMessages(t *testing.T) {
	ctx := context.Background()
	app := newShotApp(t, ctx)

	w, viewH := app.Page().Size()
	view := app.View()

	for _, off := range []int{300, 1200} {
		pinned := shot(t, app, ctx, off, off)
		content := shot(t, app, ctx, off, off-viewH)

		barRect := image.Rect(0, view.InsetTop, w, view.InsetTop+60)
		composeRect := image.Rect(0, viewH-view.InsetBottom-63, w, viewH-view.InsetBottom)

		if samePix(crop(pinned, barRect), crop(content, barRect)) {
			t.Errorf("header band at offset %d is not covering content", off)
		}

		if samePix(crop(pinned, composeRect), crop(content, composeRect)) {
			t.Errorf("composer band at offset %d is not covering content", off)
		}
	}
}
