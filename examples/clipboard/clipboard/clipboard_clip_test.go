package clipboard_test

import (
	"context"
	"strings"
	"testing"
)

func TestCopyCutPasteUndoRedo(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "left")
	if err := app.Type(ctx, "hello"); err != nil {
		t.Fatal(err)
	}

	if err := app.SelectAll(ctx); err != nil {
		t.Fatal(err)
	}

	text, ok, err := app.CopyText(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if !ok || text != "hello" {
		t.Fatalf("copy = %q ok=%v, want hello", text, ok)
	}

	text, ok, err = app.CutText(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if !ok || text != "hello" {
		t.Fatalf("cut = %q ok=%v, want hello", text, ok)
	}

	if got := app.FormValue("left"); got != "" {
		t.Fatalf("after cut = %q, want empty", got)
	}

	prev := app.FormValue("left")
	if err := app.PasteText(ctx, "clip"); err != nil {
		t.Fatal(err)
	}

	if got := app.FormValue("left"); !strings.Contains(got, "clip") {
		t.Fatalf("pasted = %q, want clip", got)
	}

	if err := app.Undo(ctx); err != nil {
		t.Fatal(err)
	}

	if got := app.FormValue("left"); got != prev {
		t.Fatalf("undo = %q, want %q", got, prev)
	}

	if err := app.Redo(ctx); err != nil {
		t.Fatal(err)
	}

	if got := app.FormValue("left"); !strings.Contains(got, "clip") {
		t.Fatalf("redo = %q, want clip", got)
	}
}
