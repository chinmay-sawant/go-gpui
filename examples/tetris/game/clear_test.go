package game

import (
	"fmt"
	"testing"
)

func TestFixturesClearOneToFourRowsAtOnce(t *testing.T) {
	for n := 1; n <= 4; n++ {
		g, err := NewFromFixture(fmt.Sprintf("clear-%d", n), 1)
		if err != nil {
			t.Fatal(err)
		}

		ev := g.Apply(ActionHardDrop)

		cleared, ok := clearEvent(ev)
		if !ok {
			t.Fatalf("clear-%d returned no clear event", n)
		}

		if cleared.Lines != n {
			t.Fatalf("clear-%d cleared %d rows", n, cleared.Lines)
		}

		if g.Lines != n {
			t.Fatalf("clear-%d lines = %d", n, g.Lines)
		}

		if g.Score != LineScores[n] {
			t.Fatalf("clear-%d score = %d, want %d", n, g.Score, LineScores[n])
		}
	}
}

func TestFourRowClearEmptiesTheBoard(t *testing.T) {
	g, err := NewFromFixture("clear-4", 1)
	if err != nil {
		t.Fatal(err)
	}

	g.Apply(ActionHardDrop)

	for y := 0; y < Rows; y++ {
		for x := 0; x < Cols; x++ {
			if g.Board[y][x] != 0 {
				t.Fatalf("cell %d,%d survived the clear", y, x)
			}
		}
	}
}

func TestOneRowClearShiftsTheStackDown(t *testing.T) {
	g, err := NewFromFixture("clear-1", 1)
	if err != nil {
		t.Fatal(err)
	}

	g.Apply(ActionHardDrop)

	for y := 17; y <= 19; y++ {
		if g.Board[y][0] == 0 {
			t.Fatalf("row %d lost the I column", y)
		}
	}

	if g.Board[16][0] != 0 {
		t.Fatal("row 16 should be empty after the shift")
	}
}

func TestClearRowsRemovesNonAdjacentRowsTogether(t *testing.T) {
	var b Board
	for x := 0; x < Cols; x++ {
		b[19][x] = Cell(PieceJ)
		b[17][x] = Cell(PieceJ)
	}

	b[16][0] = Cell(PieceL)

	full := b.fullRows()
	if n := b.clearRows(full); n != 2 {
		t.Fatalf("cleared %d rows, want 2", n)
	}

	if b[18][0] == 0 {
		t.Fatal("the marked cell should shift down two rows")
	}

	filled := 0
	for y := range b {
		for x := range b[y] {
			if b[y][x] != 0 {
				filled++
			}
		}
	}

	if filled != 1 {
		t.Fatalf("%d cells left, want 1", filled)
	}
}
