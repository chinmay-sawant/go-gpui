package scrolling_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe"
	"github.com/chinmay-sawant/ownframe/examples/scrolling/scrolling"
)

func TestTallPageAndResize(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	first := findBox(t, app, "row-1")
	if first.Y >= float64(scrolling.DefaultHeight) {
		t.Fatalf("row-1 y = %v, want above the fold", first.Y)
	}

	// The full page is laid out even though row-40 sits below the frame.
	last := findBox(t, app, "row-40")
	if last.Y <= float64(scrolling.DefaultHeight) {
		t.Fatalf("row-40 y = %v, want below the %d fold", last.Y, scrolling.DefaultHeight)
	}

	if len(app.PNG()) == 0 {
		t.Fatal("png is empty at the default size")
	}

	before := app.Generation()
	app.SetSize(240, 200)

	if err := app.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	if app.Generation() <= before {
		t.Fatalf("resize generation = %d, want past %d", app.Generation(), before)
	}

	findBox(t, app, "row-40")

	if len(app.PNG()) == 0 {
		t.Fatal("png is empty after the resize")
	}
}

func newApp(t *testing.T, ctx context.Context) *scrolling.App {
	t.Helper()

	app, err := scrolling.New()
	if err != nil {
		t.Fatal(err)
	}

	app.SetSize(scrolling.DefaultWidth, scrolling.DefaultHeight)

	if err := app.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	return app
}

func findBox(t *testing.T, app *scrolling.App, id string) ownframe.Box {
	t.Helper()

	for _, b := range app.Boxes() {
		if b.ID == id {
			return b
		}
	}

	t.Fatalf("no box id=%q", id)

	return ownframe.Box{}
}
