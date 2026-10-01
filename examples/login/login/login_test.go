package login_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/examples/login/login"
)

// center uses the last matching box so a nested element wins over its parent.
func center(t *testing.T, app *login.App, id, action string) (float64, float64) {
	t.Helper()

	var x, y float64
	found := false

	for _, b := range app.Boxes() {
		if id != "" && b.ID != id {
			continue
		}

		if action != "" && b.Action != action {
			continue
		}

		if b.W <= 0 || b.H <= 0 {
			continue
		}

		x = b.X + b.W/2
		y = b.Y + b.H/2
		found = true
	}

	if !found {
		t.Fatalf("no box id=%q action=%q:%s", id, action, dumpBoxes(app))
	}

	return x, y
}

func dumpBoxes(app *login.App) string {
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

func boxByID(t *testing.T, app *login.App, id string) gpui.Box {
	t.Helper()

	for _, box := range app.Boxes() {
		if box.ID == id {
			return box
		}
	}

	t.Fatalf("no box id=%q:%s", id, dumpBoxes(app))

	return gpui.Box{}
}
