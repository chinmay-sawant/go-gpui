package workbook

import (
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/formula"
)

func TestSeedDummy(t *testing.T) {
	books := SeedDummy()

	if len(books) != 2 {
		t.Fatalf("books = %d, want 2", len(books))
	}

	demo := books[0]
	if demo.Name() != "Demo workbook" || len(demo.Sheets()) != 3 {
		t.Fatalf("demo = %q with %d sheets", demo.Name(), len(demo.Sheets()))
	}

	numbers := demo.SheetByName("Numbers")
	text := demo.SheetByName("Text")
	mixed := demo.SheetByName("Mixed")

	r, ok := numbers.UsedRange()
	if !ok || r.Max.Row != DummyRows-1 || r.Max.Col != DummyCols-1 {
		t.Fatalf("numbers used range = %+v, want %d rows x %d cols", r, DummyRows, DummyCols)
	}

	if got := text.Display(Pos{0, 0}); got != "Türkçe metin" {
		t.Fatalf("Text!A1 = %q", got)
	}

	if got := text.Display(Pos{0, 5}); got != "0" {
		t.Fatalf("Text!F1 = %q", got)
	}

	if c, _ := text.Cell(Pos{0, 5}); c.Kind != Text {
		t.Fatalf("Text!F1 kind = %d, want text", c.Kind)
	}

	if got := numbers.Display(Pos{0, 5}); got != "0" {
		t.Fatalf("Numbers!F1 = %q", got)
	}

	if _, ok := numbers.Cell(Pos{1, 5}); ok {
		t.Fatal("Numbers!F2 is stored, want blank")
	}

	checks := map[Pos]string{
		{5, 2}:  formula.ErrDiv,
		{10, 4}: formula.ErrName,
		{12, 4}: formula.ErrRef,
		{0, 6}:  formula.ErrValue,
		{40, 5}: formula.ErrCycle,
		{41, 5}: formula.ErrCycle,
	}

	for p, want := range checks {
		if got := mixed.Display(p); got != want {
			t.Errorf("Mixed!%s = %q, want %s", p, got, want)
		}
	}

	if got := numbers.Display(Pos{0, 2}); got != "0" {
		t.Errorf("Numbers!C1 = %q", got)
	}

	if got := numbers.Display(Pos{1, 2}); got != "-20" {
		t.Errorf("Numbers!C2 = %q", got)
	}
}
