package ui

import (
	"context"
	"testing"
)

func TestPasteMultiCellRange(t *testing.T) {
	b := newFake()
	app := newTestApp(t, b)
	settle(t, app)

	ctx := context.Background()
	if err := app.onPaste(ctx, "1\t2\n3\t4\n"); err != nil {
		t.Fatal(err)
	}

	flush(t, app)
	if len(b.applied) != 1 || len(b.applied[0]) != 4 {
		t.Fatalf("applied = %v", b.applied)
	}

	r0, c0, r1, c1 := app.selection().Rect()
	if r0 != 0 || c0 != 0 || r1 != 1 || c1 != 1 {
		t.Fatalf("pasted range = %d,%d,%d,%d", r0, c0, r1, c1)
	}

	settle(t, app)
	if cell, _ := app.cellAt(1, 1); cell.Raw != "4" {
		t.Fatalf("D2 = %+v", cell)
	}
}

func TestPasteBoundedBySheet(t *testing.T) {
	app := newTestApp(t, newFake())
	settle(t, app)

	// The active cell is the last column of a 20-column sheet.
	app.setSelection(Selection{AnchorR: 0, AnchorC: 19, ActiveR: 0, ActiveC: 19})
	ctx := context.Background()
	if err := app.onPaste(ctx, "a\tb\tc"); err != nil {
		t.Fatal(err)
	}

	flush(t, app)
	if s := app.selection(); s.ActiveC != 19 {
		t.Fatalf("paste escaped the sheet: %+v", s)
	}
}

func TestCopySelectionTSV(t *testing.T) {
	app := newTestApp(t, newFake())
	settle(t, app)

	app.selectRect(0, 0, 1, 1)
	text, ok, err := app.onCopy(context.Background())
	if err != nil || !ok {
		t.Fatalf("copy ok=%v err=%v", ok, err)
	}

	if text != "name\tcount\nalpha\t12\n" {
		t.Fatalf("copy = %q", text)
	}
}

func TestSaveFailureRestoresEditor(t *testing.T) {
	b := newFake()
	app := newTestApp(t, b)
	settle(t, app)

	b.failSet = true
	ctx := context.Background()
	if err := app.onType(ctx, "lost"); err != nil {
		t.Fatal(err)
	}

	if err := app.onSubmit(ctx); err != nil {
		t.Fatal(err)
	}

	flush(t, app)
	if app.edit == nil || app.edit.Text != "lost" {
		t.Fatalf("editor after failure = %+v", app.edit)
	}

	if app.status == "" {
		t.Fatal("no failure status")
	}
}
