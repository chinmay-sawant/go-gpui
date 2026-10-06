package storage

import (
	"context"
	"errors"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/workbook"
)

// TestFailedSaveKeepsMemoryEdit checks that a rejected save leaves the
// in-memory workbook, and the pending command, exactly as they were.
func TestFailedSaveKeepsMemoryEdit(t *testing.T) {
	st := openMemory(t)
	ctx := context.Background()

	w, s := newBook(t, st)
	edit(t, st, w, s, workbook.Pos{Row: 0, Col: 0}, "one")

	pending := workbook.NewPending(4)

	cmd := workbook.Command{Sheet: s.ID(), Edits: []workbook.CellEdit{
		{Pos: workbook.Pos{Row: 0, Col: 0}, Cell: workbook.ParseInput("two")},
	}}

	if _, err := w.Apply(cmd); err != nil {
		t.Fatal(err)
	}

	if err := pending.Push(cmd); err != nil {
		t.Fatal(err)
	}

	// A stale base revision rejects the save.
	_, err := st.SaveCells(ctx, w.ID(), s.ID(), cmd.Edits, w.SavedRev()-1, "edit")
	if !errors.Is(err, ErrConflict) {
		t.Fatalf("save = %v, want ErrConflict", err)
	}

	if v := s.Display(workbook.Pos{Row: 0, Col: 0}); v != "two" {
		t.Fatalf("in-memory A1 = %q, the edit was lost", v)
	}

	if !w.Dirty() {
		t.Fatal("workbook reports clean after a failed save")
	}

	if pending.Len() != 1 {
		t.Fatalf("pending = %d, the command was dropped", pending.Len())
	}

	rev, err := st.SaveCells(ctx, w.ID(), s.ID(), cmd.Edits, w.SavedRev(), "retry")
	if err != nil {
		t.Fatal(err)
	}

	w.SetSavedRev(rev)

	if _, ok := pending.Pop(); !ok {
		t.Fatal("pending command vanished")
	}

	if w.Dirty() {
		t.Fatal("workbook still dirty after the acknowledged retry")
	}
}
