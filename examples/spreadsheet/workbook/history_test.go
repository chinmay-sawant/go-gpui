package workbook

import "testing"

func TestUndoRedo(t *testing.T) {
	w := New(1, "test")
	s := w.AddSheet("S")

	edit := func(v string) {
		t.Helper()

		if _, err := w.Apply(Command{Sheet: s.ID(), Edits: []CellEdit{
			{Pos: Pos{0, 0}, Cell: ParseInput(v)},
		}}); err != nil {
			t.Fatal(err)
		}
	}

	edit("1")
	edit("2")
	edit("3")

	if ok, _ := w.Undo(); !ok || s.Display(Pos{0, 0}) != "2" {
		t.Fatalf("first undo = %q", s.Display(Pos{0, 0}))
	}

	if ok, _ := w.Undo(); !ok || s.Display(Pos{0, 0}) != "1" {
		t.Fatalf("second undo = %q", s.Display(Pos{0, 0}))
	}

	if ok, _ := w.Redo(); !ok || s.Display(Pos{0, 0}) != "2" {
		t.Fatalf("redo = %q", s.Display(Pos{0, 0}))
	}

	edit("4")

	if w.CanRedo() {
		t.Fatal("a new edit kept the redo stack")
	}

	if ok, _ := w.Undo(); !ok || s.Display(Pos{0, 0}) != "2" {
		t.Fatalf("undo after branch = %q", s.Display(Pos{0, 0}))
	}
}

func mustApply(t *testing.T, w *Workbook, s *Sheet, p Pos, v string) {
	t.Helper()

	if _, err := w.Apply(Command{Sheet: s.ID(), Edits: []CellEdit{
		{Pos: p, Cell: ParseInput(v)},
	}}); err != nil {
		t.Fatal(err)
	}
}
