package crash_test

import (
	"bytes"
	"context"
	"image/png"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/crash/crash"
)

func newApp(t *testing.T, ctx context.Context) *crash.App {
	t.Helper()

	app, err := crash.New()
	if err != nil {
		t.Fatal(err)
	}

	app.SetSize(crash.DefaultWidth, crash.DefaultHeight)
	if err := app.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	if len(app.PNG()) == 0 {
		t.Fatal("png is empty")
	}

	return app
}

func TestSmokeFramePng(t *testing.T) {
	ctx := context.Background()
	app := newApp(t, ctx)

	img, err := png.Decode(bytes.NewReader(app.PNG()))
	if err != nil {
		t.Fatal(err)
	}

	if img.Bounds().Dx() != crash.DefaultWidth || img.Bounds().Dy() != crash.DefaultHeight {
		t.Fatalf("png = %v, want %dx%d", img.Bounds(), crash.DefaultWidth, crash.DefaultHeight)
	}
}
