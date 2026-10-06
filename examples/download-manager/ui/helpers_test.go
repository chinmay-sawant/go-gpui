package ui

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe"
)

// newTestApp builds an App over the fake backend.
func newTestApp(t *testing.T) (*App, *fakeBackend) {
	t.Helper()

	back := newFake()

	app, err := New(context.Background(), back)
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	return app, back
}

// findBox returns the first box with the id.
func findBox(boxes []ownframe.Box, id string) (ownframe.Box, bool) {
	for _, box := range boxes {
		if box.ID == id {
			return box, true
		}
	}

	return ownframe.Box{}, false
}

// findAction returns the first box with the data-action value.
func findAction(boxes []ownframe.Box, action string) (ownframe.Box, bool) {
	for _, box := range boxes {
		if box.Action == action {
			return box, true
		}
	}

	return ownframe.Box{}, false
}

// clickAction redraws, finds the action box, and clicks its middle.
func clickAction(t *testing.T, app *App, action string) {
	t.Helper()
	ctx := context.Background()

	if err := app.Redraw(ctx); err != nil {
		t.Fatalf("Redraw: %v", err)
	}

	clickVisible(t, app, action)
}

// clickVisible clicks the action box in the current drawing.
func clickVisible(t *testing.T, app *App, action string) {
	t.Helper()
	ctx := context.Background()

	box, ok := findAction(app.Boxes(), action)
	if !ok {
		t.Fatalf("no box with action %q", action)
	}

	if err := app.Click(ctx, box.X+box.W/2, box.Y+box.H/2); err != nil {
		t.Fatalf("Click %s: %v", action, err)
	}
}
