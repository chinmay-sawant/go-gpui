package inputlab

import (
	"testing"
)

// TestChordUndoRestoresNote types twice, then undoes and redoes through
// the chord handlers. Before the fix onUndo called page.Undo, which
// re-invoked onUndo until the stack overflowed.
func TestChordUndoRestoresNote(t *testing.T) {
	app, ctx := newTest(t)
	clickID(t, app, ctx, "e-note")
	if err := app.Type(ctx, "hello"); err != nil {
		t.Fatal(err)
	}
	if err := app.Type(ctx, " world"); err != nil {
		t.Fatal(err)
	}
	if err := app.onUndo(ctx); err != nil {
		t.Fatal(err)
	}
	if got := app.FormValue("e-note"); got != "hello" {
		t.Fatalf("after undo e-note=%q", got)
	}
	if err := app.onRedo(ctx); err != nil {
		t.Fatal(err)
	}
	if got := app.FormValue("e-note"); got != "hello world" {
		t.Fatalf("after redo e-note=%q", got)
	}
}

// TestChordUndoRestoresClip checks the clipboard pair the same way.
func TestChordUndoRestoresClip(t *testing.T) {
	app, ctx := newTest(t)
	clickID(t, app, ctx, "c-left")
	if err := app.Type(ctx, "left"); err != nil {
		t.Fatal(err)
	}
	if err := app.Type(ctx, "2"); err != nil {
		t.Fatal(err)
	}
	if err := app.onUndo(ctx); err != nil {
		t.Fatal(err)
	}
	if got := app.FormValue("c-left"); got != "left" {
		t.Fatalf("after undo c-left=%q", got)
	}
	if err := app.onRedo(ctx); err != nil {
		t.Fatal(err)
	}
	if got := app.FormValue("c-left"); got != "left2" {
		t.Fatalf("after redo c-left=%q", got)
	}
}
