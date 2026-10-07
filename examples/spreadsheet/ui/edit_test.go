package ui

import (
	"context"
	"testing"
)

func TestTypingEscapeAndCommit(t *testing.T) {
	b := newFake()
	app := newTestApp(t, b)
	settle(t, app)

	ctx := context.Background()
	if err := app.onType(ctx, "hi"); err != nil {
		t.Fatal(err)
	}

	if err := app.onType(ctx, " there"); err != nil {
		t.Fatal(err)
	}

	if app.edit == nil || app.edit.Text != "hi there" || app.edit.Caret != 8 {
		t.Fatalf("editor = %+v", app.edit)
	}

	app.refresh()
	if v := app.View(); v.Editor == nil || v.Editor.Before != "hi there" {
		t.Fatalf("editor view = %+v", v.Editor)
	}

	// Escape cancels without touching the workbook.
	if err := app.onKeyDown(ctx, "escape"); err != nil {
		t.Fatal(err)
	}

	if app.edit != nil {
		t.Fatal("escape kept the editor")
	}

	settle(t, app)
	if len(b.applied) != 0 {
		t.Fatalf("cancel applied edits: %v", b.applied)
	}

	// Enter commits and moves down.
	if err := app.onType(ctx, "42"); err != nil {
		t.Fatal(err)
	}

	if err := app.onSubmit(ctx); err != nil {
		t.Fatal(err)
	}

	flush(t, app)
	if len(b.applied) != 1 || len(b.applied[0]) != 1 || b.applied[0][0].Raw != "42" {
		t.Fatalf("applied = %v", b.applied)
	}

	if s := app.selection(); s.ActiveR != 1 || s.ActiveC != 0 {
		t.Fatalf("selection after enter = %+v", s)
	}

	settle(t, app)
	if cell, ok := app.cellAt(0, 0); !ok || cell.Raw != "42" {
		t.Fatalf("cell after commit = %+v ok=%v", cell, ok)
	}
}

func TestBackspaceStartsEmptyEdit(t *testing.T) {
	app := newTestApp(t, newFake())
	settle(t, app)

	ctx := context.Background()
	if err := app.onBackspace(ctx); err != nil {
		t.Fatal(err)
	}

	if app.edit == nil || app.edit.Text != "" {
		t.Fatalf("editor = %+v", app.edit)
	}
}
