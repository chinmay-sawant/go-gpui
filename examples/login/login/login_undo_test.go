package login_test

import (
	"context"
	"testing"
)

func TestUndoRedoRestoresFieldValues(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "email", "")
	if err := app.Type(ctx, "ab"); err != nil {
		t.Fatal(err)
	}

	if err := app.Type(ctx, "xy"); err != nil {
		t.Fatal(err)
	}

	if got := app.Page().FormValue("email"); got != "abxy" {
		t.Fatalf("typed = %q", got)
	}

	if err := app.Page().Undo(ctx); err != nil {
		t.Fatal(err)
	}

	if got := app.Page().FormValue("email"); got != "ab" {
		t.Fatalf("undo = %q", got)
	}

	if err := app.Page().Undo(ctx); err != nil {
		t.Fatal(err)
	}

	if got := app.Page().FormValue("email"); got != "" {
		t.Fatalf("second undo = %q", got)
	}

	if err := app.Page().Redo(ctx); err != nil {
		t.Fatal(err)
	}

	if got := app.Page().FormValue("email"); got != "ab" {
		t.Fatalf("redo = %q", got)
	}

	if err := app.Page().Redo(ctx); err != nil {
		t.Fatal(err)
	}

	if got := app.Page().FormValue("email"); got != "abxy" {
		t.Fatalf("second redo = %q", got)
	}
}

func TestUndoGraysTheButtonAgain(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "email", "")
	if err := app.Type(ctx, "secret"); err != nil {
		t.Fatal(err)
	}

	click(t, ctx, app, "password", "")
	if err := app.Type(ctx, "secret"); err != nil {
		t.Fatal(err)
	}

	if err := app.Page().Undo(ctx); err != nil {
		t.Fatal(err)
	}

	if app.View().Ready {
		t.Fatal("ready after undoing the password")
	}

	if got := buttonFill(t, app); got != disabledButton {
		t.Fatalf("undo fill = %+v, want %+v", got, disabledButton)
	}
}
