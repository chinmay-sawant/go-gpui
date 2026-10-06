package history_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe"
	"github.com/chinmay-sawant/ownframe/examples/history/history"
)

func newApp(t *testing.T, ctx context.Context) *history.App {
	t.Helper()

	app, err := history.New()
	if err != nil {
		t.Fatal(err)
	}

	if err := app.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	return app
}

func clickAction(t *testing.T, ctx context.Context, app *history.App, action string) {
	t.Helper()

	for _, b := range app.Boxes() {
		if b.Action == action {
			if err := app.Click(ctx, b.X+b.W/2, b.Y+b.H/2); err != nil {
				t.Fatal(err)
			}

			return
		}
	}

	clickBox(t, ctx, app, action)
}

func clickBox(t *testing.T, ctx context.Context, app *history.App, id string) {
	t.Helper()

	b := findBox(t, app, id)
	if err := app.Click(ctx, b.X+b.W/2, b.Y+b.H/2); err != nil {
		t.Fatal(err)
	}
}

func findBox(t *testing.T, app *history.App, id string) ownframe.Box {
	t.Helper()

	for _, b := range app.Boxes() {
		if b.ID == id && b.W > 0 && b.H > 0 {
			return b
		}
	}

	t.Fatalf("no box id=%q", id)

	return ownframe.Box{}
}
