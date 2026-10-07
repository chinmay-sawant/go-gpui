package workbook

import (
	"testing"
)

func TestRecalcOnlyAffected(t *testing.T) {
	w := New(1, "test")
	s := w.AddSheet("S")

	mustApply(t, w, s, Pos{0, 0}, "1")
	mustApply(t, w, s, Pos{0, 1}, "=A1+1")
	mustApply(t, w, s, Pos{0, 2}, "100")
	mustApply(t, w, s, Pos{0, 3}, "=C1+1")

	res, err := w.Apply(Command{Sheet: s.ID(), Edits: []CellEdit{
		{Pos: Pos{0, 0}, Cell: ParseInput("5")},
	}})
	if err != nil {
		t.Fatal(err)
	}

	if res.Evaluated != 1 {
		t.Fatalf("editing A1 evaluated %d formulas, want 1", res.Evaluated)
	}

	if got := s.Display(Pos{0, 1}); got != "6" {
		t.Fatalf("B1 = %q, want 6", got)
	}

	if got := s.Display(Pos{0, 3}); got != "101" {
		t.Fatalf("D1 = %q, want 101", got)
	}
}

func TestRecalcChain(t *testing.T) {
	w := New(1, "test")
	s := w.AddSheet("S")

	mustApply(t, w, s, Pos{0, 0}, "1")
	mustApply(t, w, s, Pos{0, 1}, "=A1+1")
	mustApply(t, w, s, Pos{0, 2}, "=B1+1")

	res, err := w.Apply(Command{Sheet: s.ID(), Edits: []CellEdit{
		{Pos: Pos{0, 0}, Cell: ParseInput("10")},
	}})
	if err != nil {
		t.Fatal(err)
	}

	if res.Evaluated != 2 {
		t.Fatalf("chain evaluated %d, want 2", res.Evaluated)
	}

	if got := s.Display(Pos{0, 2}); got != "12" {
		t.Fatalf("C1 = %q, want 12", got)
	}
}

func TestRecalcRange(t *testing.T) {
	w := New(1, "test")
	s := w.AddSheet("S")

	mustApply(t, w, s, Pos{0, 0}, "1")
	mustApply(t, w, s, Pos{1, 0}, "2")
	mustApply(t, w, s, Pos{2, 0}, "3")
	mustApply(t, w, s, Pos{0, 1}, "=SUM(A1:A3)")

	if got := s.Display(Pos{0, 1}); got != "6" {
		t.Fatalf("B1 = %q, want 6", got)
	}

	res, err := w.Apply(Command{Sheet: s.ID(), Edits: []CellEdit{
		{Pos: Pos{1, 0}, Cell: ParseInput("20")},
	}})
	if err != nil {
		t.Fatal(err)
	}

	if res.Evaluated != 1 {
		t.Fatalf("range recalc evaluated %d, want 1", res.Evaluated)
	}

	if got := s.Display(Pos{0, 1}); got != "24" {
		t.Fatalf("B1 after edit = %q, want 24", got)
	}
}
