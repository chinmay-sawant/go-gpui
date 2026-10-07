package workbook

import (
	"testing"
)

func TestSeedStress(t *testing.T) {
	w := SeedStress(1000)

	if len(w.Sheets()) != 1 {
		t.Fatalf("sheets = %d", len(w.Sheets()))
	}

	s := w.Sheets()[0]

	if r, _ := s.UsedRange(); r.Max.Row != 999 {
		t.Fatalf("used range = %+v", r)
	}

	if s.Count() >= 1000*2 {
		t.Fatalf("sheet is not sparse: %d cells", s.Count())
	}

	if got := s.Display(Pos{50, 2}); got != "50" {
		t.Fatalf("C51 = %q, want 50", got)
	}

	if got := s.Display(Pos{500, 3}); got != "24225" {
		t.Fatalf("D501 = %q, want 24225", got)
	}
}

func TestSeedDeterministic(t *testing.T) {
	a := SeedDummy()[0]
	b := SeedDummy()[0]

	if a.SheetByName("Numbers").Count() != b.SheetByName("Numbers").Count() {
		t.Fatal("seed counts differ")
	}

	if a.SheetByName("Mixed").Display(Pos{40, 5}) != b.SheetByName("Mixed").Display(Pos{40, 5}) {
		t.Fatal("seed values differ")
	}
}
