package bind_test

import (
	"context"
	"strings"
	"testing"

	"github.com/chinmay-sawant/go-gpui/examples/bind/bind"
)

func TestBindingWritesView(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "name")
	if err := app.Type(ctx, "ada"); err != nil {
		t.Fatal(err)
	}

	if got := boxText(t, app, "state"); !strings.Contains(got, "ada") {
		t.Fatalf("state = %q, want the typed name", got)
	}

	click(t, ctx, app, "agree")
	if !app.View().Agree {
		t.Fatal("Agree = false, want true")
	}

	click(t, ctx, app, "color")
	if got := app.View().Color; got != "Blue" {
		t.Fatalf("Color = %q, want Blue", got)
	}

	if got := app.View().Status; !strings.Contains(got, "changed") {
		t.Fatalf("Status = %q, want a change note", got)
	}
}

func newApp(t *testing.T, ctx context.Context) *bind.App {
	t.Helper()

	app, err := bind.New()
	if err != nil {
		t.Fatal(err)
	}

	app.SetSize(bind.DefaultWidth, bind.DefaultHeight)
	if err := app.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	return app
}

func click(t *testing.T, ctx context.Context, app *bind.App, id string) {
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

func boxText(t *testing.T, app *bind.App, id string) string {
	t.Helper()

	var text string
	found := false

	for _, b := range app.Boxes() {
		if b.ID != id {
			continue
		}

		text = b.Text
		found = true
	}

	if !found {
		t.Fatalf("no box id=%q", id)
	}

	return text
}
