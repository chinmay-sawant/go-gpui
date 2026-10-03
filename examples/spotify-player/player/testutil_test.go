package player

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/go-gpui"
)

func mustApp(t *testing.T) *App {
	t.Helper()

	app, err := New()
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	return app
}

func redraw(t *testing.T, app *App) {
	t.Helper()

	if err := app.Redraw(context.Background()); err != nil {
		t.Fatalf("Redraw: %v", err)
	}

	if app.Page().Display() == nil {
		t.Fatal("fell back to the bitmap path")
	}

	if len(app.PNG()) == 0 {
		t.Fatal("no PNG after Redraw")
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

func click(t *testing.T, app *App, id string) {
	t.Helper()

	box, ok := findBox(app.Boxes(), id)
	if !ok {
		t.Fatalf("no box %s", id)
	}

	if err := app.Click(context.Background(), box.X+box.W/2, box.Y+box.H/2); err != nil {
		t.Fatalf("click %s: %v", id, err)
	}
}
