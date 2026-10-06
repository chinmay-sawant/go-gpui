package workbook

import (
	"errors"
	"testing"
)

func TestApply(t *testing.T) {
	w := New(1, "test")
	s := w.AddSheet("S")

	res, err := w.Apply(Command{Sheet: s.ID(), Edits: []CellEdit{
		{Pos: Pos{0, 0}, Cell: ParseInput("2")},
		{Pos: Pos{0, 1}, Cell: ParseInput("=A1*3")},
	}})
	if err != nil {
		t.Fatal(err)
	}

	if res.Evaluated != 1 || w.Rev() != 1 {
		t.Fatalf("res=%+v rev=%d, want 1 evaluation at rev 1", res, w.Rev())
	}

	if got := s.Display(Pos{0, 1}); got != "6" {
		t.Fatalf("B1 = %q, want 6", got)
	}

	if _, err := w.Apply(Command{Sheet: s.ID(), Edits: []CellEdit{
		{Pos: Pos{0, 0}, Cell: Cell{}},
	}}); err != nil {
		t.Fatal(err)
	}

	if got := s.Display(Pos{0, 1}); got != "0" {
		t.Fatalf("B1 after clearing A1 = %q, want 0", got)
	}

	if w.Rev() != 2 {
		t.Fatalf("rev = %d, want 2", w.Rev())
	}
}

func TestApplyErrors(t *testing.T) {
	w := New(1, "test")
	s := w.AddSheet("S")

	if _, err := w.Apply(Command{Sheet: 99, Edits: []CellEdit{
		{Pos: Pos{0, 0}, Cell: ParseInput("1")},
	}}); !errors.Is(err, ErrNoSheet) {
		t.Fatalf("unknown sheet err = %v", err)
	}

	if _, err := w.Apply(Command{Sheet: s.ID(), Edits: []CellEdit{
		{Pos: Pos{Row: -1}, Cell: ParseInput("1")},
	}}); !errors.Is(err, ErrBadRef) {
		t.Fatalf("bad ref err = %v", err)
	}

	if w.Rev() != 0 {
		t.Fatalf("rev = %d after failed applies", w.Rev())
	}

	res, err := w.Apply(Command{Sheet: s.ID()})
	if err != nil || res.Evaluated != 0 {
		t.Fatalf("empty apply = %+v, %v", res, err)
	}
}

func TestApplyDuplicatePosition(t *testing.T) {
	w := New(1, "test")
	s := w.AddSheet("S")

	if _, err := w.Apply(Command{Sheet: s.ID(), Edits: []CellEdit{
		{Pos: Pos{0, 0}, Cell: ParseInput("first")},
		{Pos: Pos{0, 0}, Cell: ParseInput("second")},
	}}); err != nil {
		t.Fatal(err)
	}

	if got := s.Display(Pos{0, 0}); got != "first" {
		t.Fatalf("A1 = %q, want the first edit", got)
	}
}
