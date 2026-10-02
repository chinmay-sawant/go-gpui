package controls_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/chinmay-sawant/go-gpui/examples/controls/controls"
)

// click uses the last matching box so a nested element wins over its parent.
func click(t *testing.T, ctx context.Context, app *controls.App, id string) {
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

	if err := app.Click(ctx, x, y); err != nil {
		t.Fatal(err)
	}
}

func boxText(t *testing.T, app *controls.App, id string) string {
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

func dumpBoxes(app *controls.App) string {
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
