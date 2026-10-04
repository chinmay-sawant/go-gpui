package drop_test

import (
	"bytes"
	"context"
	"image"
	"image/png"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/chinmay-sawant/go-gpui/examples/drop/drop"
)

func TestDroppedPNGShowsThroughSetImage(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)
	before := app.PNG()

	if hasRed(before) {
		t.Fatal("the page was red before the drop")
	}

	fsys := fstest.MapFS{"photo.png": {Data: redPNG(t)}}

	if err := app.Drop(ctx, dropsFrom(t, fsys)); err != nil {
		t.Fatal(err)
	}

	if got := app.View().Status; !strings.Contains(got, "photo.png") {
		t.Fatalf("Status = %q", got)
	}

	after := app.PNG()
	if bytes.Equal(before, after) {
		t.Fatal("the picture did not change")
	}

	if !hasRed(after) {
		t.Fatal("no red pixel in the page")
	}
}

func newApp(t *testing.T, ctx context.Context) *drop.App {
	t.Helper()

	app, err := drop.New()
	if err != nil {
		t.Fatal(err)
	}

	app.SetSize(drop.DefaultWidth, drop.DefaultHeight)

	if err := app.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	return app
}

// redPNG encodes a 16x16 opaque red image.
func redPNG(t *testing.T) []byte {
	t.Helper()

	img := image.NewNRGBA(image.Rect(0, 0, 16, 16))
	for i := 0; i < len(img.Pix); i += 4 {
		img.Pix[i] = 255
		img.Pix[i+3] = 255
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		t.Fatal(err)
	}

	return buf.Bytes()
}

// hasRed reports whether the decoded page has a red pixel.
func hasRed(data []byte) bool {
	img, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		return false
	}

	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			if r > 0xc000 && g < 0x4000 && b < 0x4000 {
				return true
			}
		}
	}

	return false
}
