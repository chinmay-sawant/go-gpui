package editing_test

import (
	"context"
	"testing"
)

func TestCopyCutPasteAndUndo(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "note")
	if err := app.Type(ctx, "abc"); err != nil {
		t.Fatal(err)
	}

	if err := app.Paste(ctx, "def"); err != nil {
		t.Fatal(err)
	}

	if got := app.FormValue("note"); got != "abcdef" {
		t.Fatalf("paste = %q, want abcdef", got)
	}

	text, ok, err := app.Copy(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if !ok || text != "abcdef" {
		t.Fatalf("copy = %q ok=%v, want abcdef", text, ok)
	}

	text, ok, err = app.Cut(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if !ok || text != "abcdef" {
		t.Fatalf("cut = %q ok=%v, want abcdef", text, ok)
	}

	if got := app.FormValue("note"); got != "" {
		t.Fatalf("after cut = %q, want empty", got)
	}

	if err := app.Undo(ctx); err != nil {
		t.Fatal(err)
	}

	if got := app.FormValue("note"); got != "abcdef" {
		t.Fatalf("undo after cut = %q, want abcdef", got)
	}
}
