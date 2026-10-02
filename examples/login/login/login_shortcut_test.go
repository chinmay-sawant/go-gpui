package login_test

import (
	"context"
	"testing"
)

func TestCopyPasteCutAndUndo(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "email", "")
	if err := app.Type(ctx, "ab"); err != nil {
		t.Fatal(err)
	}

	text, ok, err := app.Page().Copy(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if !ok || text != "ab" {
		t.Fatalf("copy = %q ok=%v", text, ok)
	}

	if err := app.Page().Paste(ctx, "xy"); err != nil {
		t.Fatal(err)
	}

	if got := app.Page().FormValue("email"); got != "abxy" {
		t.Fatalf("pasted = %q", got)
	}

	text, ok, err = app.Page().Cut(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if !ok || text != "abxy" {
		t.Fatalf("cut = %q ok=%v", text, ok)
	}

	if got := app.Page().FormValue("email"); got != "" {
		t.Fatalf("after cut = %q", got)
	}

	if err := app.Page().Undo(ctx); err != nil {
		t.Fatal(err)
	}

	if got := app.Page().FormValue("email"); got != "abxy" {
		t.Fatalf("undo = %q", got)
	}
}

func TestSelectAllReplacesTheField(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "email", "")
	if err := app.Type(ctx, "abxy"); err != nil {
		t.Fatal(err)
	}

	if err := app.Page().SelectAll(ctx); err != nil {
		t.Fatal(err)
	}

	if err := app.Type(ctx, "z"); err != nil {
		t.Fatal(err)
	}

	if got := app.Page().FormValue("email"); got != "z" {
		t.Fatalf("replaced = %q", got)
	}
}
