package editing_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/editing/editing"
)

func newApp(t *testing.T, ctx context.Context) *editing.App {
	t.Helper()

	app, err := editing.New()
	if err != nil {
		t.Fatal(err)
	}

	app.SetSize(editing.DefaultWidth, editing.DefaultHeight)
	if err := app.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	return app
}

// click uses the last matching box so a nested element wins over its parent.
func click(t *testing.T, ctx context.Context, app *editing.App, id string) {
	t.Helper()

	var x, y float64
	found := false

	for _, b := range app.Boxes() {
		if b.ID != id || b.W <= 0 || b.H <= 0 {
			continue
		}

		x = b.X + b.W/2
		y = b.Y + b.H/2
		found = true
	}

	if !found {
		t.Fatalf("no box id=%q", id)
	}

	if err := app.Click(ctx, x, y); err != nil {
		t.Fatal(err)
	}
}
