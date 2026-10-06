package workbook

import (
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/formula"
)

func TestRecalcLimits(t *testing.T) {
	w := New(1, "test")
	w.SetLimits(formula.Limits{MaxEvals: 2})
	s := w.AddSheet("S")

	mustApply(t, w, s, Pos{0, 0}, "1")
	mustApply(t, w, s, Pos{0, 1}, "2")
	mustApply(t, w, s, Pos{0, 2}, "3")
	mustApply(t, w, s, Pos{0, 3}, "=A1+B1+C1")

	res, err := w.Apply(Command{Sheet: s.ID(), Edits: []CellEdit{
		{Pos: Pos{0, 0}, Cell: ParseInput("10")},
	}})
	if err != nil {
		t.Fatal(err)
	}

	if !res.Truncated {
		t.Fatalf("expected truncation: %+v", res)
	}

	if got := s.Display(Pos{0, 3}); got != formula.ErrLimit {
		t.Fatalf("D1 = %q, want %s", got, formula.ErrLimit)
	}
}

func TestGraphDepthLimit(t *testing.T) {
	w := New(1, "test")
	w.SetLimits(formula.Limits{MaxDepth: 1})
	s := w.AddSheet("S")

	mustApply(t, w, s, Pos{0, 0}, "1")
	mustApply(t, w, s, Pos{0, 1}, "=A1+1")
	mustApply(t, w, s, Pos{0, 2}, "=B1+1")

	res, err := w.Apply(Command{Sheet: s.ID(), Edits: []CellEdit{
		{Pos: Pos{0, 0}, Cell: ParseInput("2")},
	}})
	if err != nil {
		t.Fatal(err)
	}

	if !res.Truncated {
		t.Fatalf("depth limit not reported: %+v", res)
	}

	if got := s.Display(Pos{0, 2}); got != formula.ErrLimit {
		t.Fatalf("C1 = %q, want %s", got, formula.ErrLimit)
	}
}
