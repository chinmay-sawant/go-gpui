package page_test

import (
	"bytes"
	"context"
	"image/png"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/page"
)

func TestImageBeforeRedrawIsNil(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	screen, err := page.New(page.Config{
		HTML:   "<p>Hi</p>",
		Width:  320,
		Height: 400,
	})
	if err != nil {
		t.Fatal(err)
	}

	if screen.Image() != nil {
		t.Fatal("image before redraw")
	}

	if screen.PNG() != nil {
		t.Fatal("png before redraw")
	}

	if err := screen.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	img := screen.Image()
	if img == nil {
		t.Fatal("image is nil")
	}

	bounds := img.Bounds()
	if bounds.Dx() < 320 || bounds.Dy() < 400 {
		t.Fatalf("image = %d x %d", bounds.Dx(), bounds.Dy())
	}

	cfg, err := png.DecodeConfig(bytes.NewReader(screen.PNG()))
	if err != nil {
		t.Fatal(err)
	}

	if cfg.Width != bounds.Dx() || cfg.Height != bounds.Dy() {
		t.Fatalf("png = %d x %d, image = %d x %d", cfg.Width, cfg.Height, bounds.Dx(), bounds.Dy())
	}
}
