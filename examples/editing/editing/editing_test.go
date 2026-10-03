package editing_test

import (
	"context"
	"testing"
)

func TestEditingRender(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	if len(app.PNG()) == 0 {
		t.Fatal("PNG is empty")
	}
}

func TestEditingShortcuts(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "note")
	if err := app.Type(ctx, "abc"); err != nil {
		t.Fatal(err)
	}

	if got := app.FormValue("note"); got != "abc" {
		t.Fatalf("typed = %q, want abc", got)
	}

	if err := app.Undo(ctx); err != nil {
		t.Fatal(err)
	}

	if got := app.FormValue("note"); got != "" {
		t.Fatalf("undo = %q, want empty", got)
	}

	if err := app.Redo(ctx); err != nil {
		t.Fatal(err)
	}

	if got := app.FormValue("note"); got != "abc" {
		t.Fatalf("redo = %q, want abc", got)
	}

	if err := app.SelectAll(ctx); err != nil {
		t.Fatal(err)
	}

	if !app.FormSelected("note") {
		t.Fatal("note = false, want selected")
	}

	if err := app.Type(ctx, "z"); err != nil {
		t.Fatal(err)
	}

	if got := app.FormValue("note"); got != "z" {
		t.Fatalf("selected type = %q, want z", got)
	}

	before := app.View().Stamp
	click(t, ctx, app, "redraw")

	if got := app.View().Stamp; got != before+1 {
		t.Fatalf("Stamp = %d, want %d", got, before+1)
	}

	if got := app.FormValue("note"); got != "z" {
		t.Fatalf("after redraw = %q, want z", got)
	}
}
