package input_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe"
	"github.com/chinmay-sawant/ownframe/examples/input/input"
)

func TestInputFieldsAndLongPage(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	email := findBox(t, app, "email")
	if err := app.Click(ctx, email.X+email.W/2, email.Y+email.H/2); err != nil {
		t.Fatal(err)
	}

	if got := app.FocusedField(); got != "email" {
		t.Fatalf("focused = %q", got)
	}

	if err := app.Type(ctx, "a@b.c"); err != nil {
		t.Fatal(err)
	}

	if got := app.FormValue("email"); got != "a@b.c" {
		t.Fatalf("email = %q", got)
	}

	remember := findBox(t, app, "remember")
	if err := app.Click(ctx, remember.X+remember.W/2, remember.Y+remember.H/2); err != nil {
		t.Fatal(err)
	}

	if !app.FormChecked("remember") {
		t.Fatal("checkbox did not toggle")
	}

	last := findBox(t, app, "row-30")
	if last.Y <= float64(input.DefaultHeight) {
		t.Fatalf("row-30 y = %v, want below the fold", last.Y)
	}
}

func newApp(t *testing.T, ctx context.Context) *input.App {
	t.Helper()

	app, err := input.New()
	if err != nil {
		t.Fatal(err)
	}

	app.SetSize(input.DefaultWidth, input.DefaultHeight)

	if err := app.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	return app
}

func findBox(t *testing.T, app *input.App, id string) ownframe.Box {
	t.Helper()

	for _, b := range app.Boxes() {
		if b.ID == id {
			return b
		}
	}

	t.Fatalf("no box id=%q", id)

	return ownframe.Box{}
}
