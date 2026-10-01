package form_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/chinmay-sawant/go-gpui/examples/forms/form"
)

func newApp(t *testing.T, ctx context.Context) *form.App {
	t.Helper()

	app, err := form.New()
	if err != nil {
		t.Fatal(err)
	}

	app.SetSize(640, 640)
	if err := app.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	return app
}

func click(t *testing.T, ctx context.Context, app *form.App, id string) {
	t.Helper()

	x, y := center(t, app, id)
	if err := app.Click(ctx, x, y); err != nil {
		t.Fatal(err)
	}
}

func center(t *testing.T, app *form.App, id string) (float64, float64) {
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
		t.Fatalf("no box id=%q:%s", id, dumpBoxes(app))
	}

	return x, y
}

func boxText(t *testing.T, app *form.App, id string) string {
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
		t.Fatalf("no box id=%q:%s", id, dumpBoxes(app))
	}

	return text
}

func dumpBoxes(app *form.App) string {
	var b strings.Builder

	for _, box := range app.Boxes() {
		fmt.Fprintf(
			&b,
			"\n%s id=%q action=%q text=%q xywh=%.1f,%.1f,%.1f,%.1f",
			box.Tag,
			box.ID,
			box.Action,
			box.Text,
			box.X,
			box.Y,
			box.W,
			box.H,
		)
	}

	if b.Len() == 0 {
		return "\n(no boxes)"
	}

	return b.String()
}
