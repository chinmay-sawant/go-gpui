package workbook

import (
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/formula"
)

func TestCycles(t *testing.T) {
	w := New(1, "test")
	s := w.AddSheet("S")

	mustApply(t, w, s, Pos{0, 0}, "=B1")
	mustApply(t, w, s, Pos{0, 1}, "=A1")
	mustApply(t, w, s, Pos{2, 0}, "=A3+1")

	if got := s.Display(Pos{0, 0}); got != formula.ErrCycle {
		t.Fatalf("A1 = %q, want %s", got, formula.ErrCycle)
	}

	if got := s.Display(Pos{0, 1}); got != formula.ErrCycle {
		t.Fatalf("B1 = %q, want %s", got, formula.ErrCycle)
	}

	if got := s.Display(Pos{2, 0}); got != formula.ErrCycle {
		t.Fatalf("A3 = %q, want %s", got, formula.ErrCycle)
	}
}

func TestStaleRecalc(t *testing.T) {
	w := New(1, "test")
	s := w.AddSheet("S")

	mustApply(t, w, s, Pos{0, 0}, "1")
	mustApply(t, w, s, Pos{0, 1}, "=A1+1")

	rev := w.Rev()
	mustApply(t, w, s, Pos{0, 0}, "9")

	res := w.Recalc(RecalcRequest{Rev: rev, Sheet: s.ID(), Changed: []Pos{{0, 1}}})
	if !res.Stale {
		t.Fatalf("old revision was not stale: %+v", res)
	}

	res = w.Recalc(RecalcRequest{Rev: w.Rev(), Sheet: s.ID(), Changed: []Pos{{0, 0}}})
	if res.Stale || res.Evaluated != 1 {
		t.Fatalf("fresh recalc = %+v", res)
	}
}
