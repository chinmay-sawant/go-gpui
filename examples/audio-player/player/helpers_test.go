package player

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/go-gpui"
)

func newApp(t *testing.T) *App {
	t.Helper()

	app, err := New()
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	if err := app.Redraw(context.Background()); err != nil {
		t.Fatalf("Redraw: %v", err)
	}

	return app
}

func clickBox(t *testing.T, app *App, id string) {
	t.Helper()

	box, ok := findBox(app.Boxes(), id)
	if !ok {
		t.Fatalf("no %s box", id)
	}

	if err := app.Click(context.Background(), box.X+box.W/2, box.Y+box.H/2); err != nil {
		t.Fatalf("click %s: %v", id, err)
	}
}

// assertReplayable pins the display-list path after a click redraw.
func assertReplayable(t *testing.T, app *App) {
	t.Helper()

	if app.Page().Display() == nil {
		t.Fatal("the frame fell back to the bitmap path")
	}

	if len(app.PNG()) == 0 {
		t.Fatal("no PNG after the redraw")
	}
}

func findBox(boxes []gpui.Box, id string) (gpui.Box, bool) {
	for _, box := range boxes {
		if box.ID == id {
			return box, true
		}
	}

	return gpui.Box{}, false
}
