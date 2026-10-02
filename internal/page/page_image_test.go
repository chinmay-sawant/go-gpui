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

	if screen.Image() != nil || screen.Display() != nil {
		t.Fatal("page drawn before redraw")
	}

	if screen.PNG() != nil {
		t.Fatal("png before redraw")
	}

	if err := screen.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	display := screen.Display()
	if display == nil {
		t.Fatal("display is nil")
	}

	if display.Width < 320 || display.Height < 400 {
		t.Fatalf("display = %d x %d", display.Width, display.Height)
	}

	if screen.Image() != nil {
		t.Fatal("replayable page kept a bitmap")
	}

	cfg, err := png.DecodeConfig(bytes.NewReader(screen.PNG()))
	if err != nil {
		t.Fatal(err)
	}

	if cfg.Width != display.Width || cfg.Height < 400 {
		t.Fatalf("png = %d x %d, display = %d x %d", cfg.Width, cfg.Height, display.Width, display.Height)
	}
}
