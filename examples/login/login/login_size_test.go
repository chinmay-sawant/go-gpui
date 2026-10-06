package login_test

import (
	"bytes"
	"context"
	"image/png"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/login/login"
)

func TestSetSizeClampsAndRedraws(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)
	before := app.Generation()

	app.SetSize(10, 10)
	width, height := app.Size()
	if width != login.MinWidth || height != login.MinHeight {
		t.Fatalf("clamped size = %d x %d", width, height)
	}

	if err := app.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	cfg, err := png.DecodeConfig(bytes.NewReader(app.PNG()))
	if err != nil {
		t.Fatal(err)
	}

	if cfg.Width != login.MinWidth || cfg.Height != login.MinHeight {
		t.Fatalf("png = %d x %d", cfg.Width, cfg.Height)
	}

	if app.Generation() <= before {
		t.Fatalf("generation = %d, before = %d", app.Generation(), before)
	}
}
