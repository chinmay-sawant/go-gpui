package login_test

import (
	"testing"

	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/examples/login/login"
)

func boxByTag(t *testing.T, app *login.App, tag string) gpui.Box {
	t.Helper()

	for _, box := range app.Boxes() {
		if box.Tag == tag {
			return box
		}
	}

	t.Fatalf("no box tag=%q:%s", tag, dumpBoxes(app))

	return gpui.Box{}
}

func boxText(t *testing.T, app *login.App, id string) string {
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
