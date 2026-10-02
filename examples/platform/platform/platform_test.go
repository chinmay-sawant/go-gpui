package platform_test

import (
	"bytes"
	"context"
	"image/png"
	"testing"

	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/examples/platform/platform"
)

func newApp(t *testing.T, ctx context.Context) *platform.App {
	t.Helper()

	app, err := platform.New()
	if err != nil {
		t.Fatal(err)
	}

	app.SetSize(platform.DefaultWidth, platform.DefaultHeight)
	if err := app.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	if len(app.PNG()) == 0 {
		t.Fatal("png is empty")
	}

	return app
}

func boxByID(t *testing.T, app *platform.App, id string) gpui.Box {
	t.Helper()

	for _, box := range app.Boxes() {
		if box.ID == id {
			return box
		}
	}

	t.Fatalf("no box id=%q", id)

	return gpui.Box{}
}

func click(t *testing.T, ctx context.Context, app *platform.App, id string) {
	t.Helper()

	box := boxByID(t, app, id)
	if err := app.Click(ctx, box.X+box.W/2, box.Y+box.H/2); err != nil {
		t.Fatal(err)
	}
}

func TestSmokeFramePng(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	img, err := png.Decode(bytes.NewReader(app.PNG()))
	if err != nil {
		t.Fatal(err)
	}

	if img.Bounds().Dx() != platform.DefaultWidth || img.Bounds().Dy() != platform.DefaultHeight {
		t.Fatalf("png = %v, want %dx%d", img.Bounds(), platform.DefaultWidth, platform.DefaultHeight)
	}
}
