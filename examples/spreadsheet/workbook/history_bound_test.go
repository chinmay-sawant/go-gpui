package workbook

import (
	"strconv"
	"testing"
)

func TestUndoRecalculates(t *testing.T) {
	w := New(1, "test")
	s := w.AddSheet("S")

	mustApply(t, w, s, Pos{0, 0}, "1")
	mustApply(t, w, s, Pos{0, 1}, "=A1*2")

	if got := s.Display(Pos{0, 1}); got != "2" {
		t.Fatalf("B1 = %q", got)
	}

	mustApply(t, w, s, Pos{0, 0}, "5")

	if got := s.Display(Pos{0, 1}); got != "10" {
		t.Fatalf("B1 after edit = %q", got)
	}

	if ok, _ := w.Undo(); !ok || s.Display(Pos{0, 1}) != "2" {
		t.Fatalf("B1 after undo = %q", s.Display(Pos{0, 1}))
	}
}

func TestHistoryBound(t *testing.T) {
	w := New(1, "test")
	s := w.AddSheet("S")
	w.SetHistoryLimit(2)

	for i := 0; i < 5; i++ {
		mustApply(t, w, s, Pos{0, 0}, strconv.Itoa(i+1))
	}

	if w.UndoDepth() != 2 {
		t.Fatalf("UndoDepth = %d, want 2", w.UndoDepth())
	}

	if ok, _ := w.Undo(); !ok {
		t.Fatal("undo 1 failed")
	}

	if ok, _ := w.Undo(); !ok {
		t.Fatal("undo 2 failed")
	}

	if w.CanUndo() {
		t.Fatal("undo went past the bound")
	}

	if got := s.Display(Pos{0, 0}); got != "3" {
		t.Fatalf("A1 after bounded undo = %q, want 3", got)
	}
}
