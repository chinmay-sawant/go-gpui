package telegram_test

import (
	"context"
	"image"
	"testing"
)

// TestInsetStripsStayPanel proves the strips the system bars cover paint the
// panel, not scrolled messages. The status pad and the composer's nav strip
// fill them while the thread scrolls.
func TestInsetStripsStayPanel(t *testing.T) {
	ctx := context.Background()
	app := newShotApp(t, ctx)

	app.SetInsets(28, 32)
	if err := app.Tick(ctx); err != nil {
		t.Fatal(err)
	}

	w, viewH := app.Page().Size()
	view := app.View()

	for _, off := range []int{300, 1200} {
		img := shot(t, app, ctx, off, off)

		top := crop(img, image.Rect(0, 0, w, view.InsetTop))
		bot := crop(img, image.Rect(0, viewH-view.InsetBottom, w, viewH))

		if !uniform(top) {
			t.Errorf("top inset strip shows content at offset %d", off)
		}

		if !uniform(bot) {
			t.Errorf("bottom inset strip shows content at offset %d", off)
		}
	}
}

// uniform reports whether every pixel equals the top-left one.
func uniform(img *image.RGBA) bool {
	b := img.Bounds()
	c := img.RGBAAt(0, 0)

	for y := 0; y < b.Dy(); y++ {
		for x := 0; x < b.Dx(); x++ {
			if img.RGBAAt(x, y) != c {
				return false
			}
		}
	}

	return true
}
