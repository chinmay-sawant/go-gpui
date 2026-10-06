package game

import "testing"

func TestFixturesAreRunnable(t *testing.T) {
	names := FixtureNames()

	if len(names) != 6 {
		t.Fatalf("fixture count %d, want 6", len(names))
	}

	for _, name := range names {
		g, err := NewFromFixture(name, 3)
		if err != nil {
			t.Fatalf("%s: %v", name, err)
		}

		expectPhase(t, g, PhaseRunning)

		if !g.Piece.Valid() {
			t.Fatalf("%s started without a piece", name)
		}

		if len(g.Next) != Preview {
			t.Fatalf("%s preview has %d pieces", name, len(g.Next))
		}
	}

	if _, err := NewFromFixture("nope", 1); err == nil {
		t.Fatal("an unknown fixture was accepted")
	}
}

func TestNearTopOutFixtureIsTwoRowsFromTheTop(t *testing.T) {
	g, err := NewFromFixture("near-top-out", 1)
	if err != nil {
		t.Fatal(err)
	}

	if g.Board[0][0] != 0 || g.Board[1][0] != 0 {
		t.Fatal("the top two rows must be clear")
	}

	if g.Board[2][0] == 0 {
		t.Fatal("row 2 should hold the stack")
	}

	if g.Board[2][4] != 0 {
		t.Fatal("the middle well should be open")
	}
}

func TestEmptyFixtureStartsWithAnEmptyBoard(t *testing.T) {
	g, err := NewFromFixture("empty", 1)
	if err != nil {
		t.Fatal(err)
	}

	for y := 0; y < Rows; y++ {
		for x := 0; x < Cols; x++ {
			if g.Board[y][x] != 0 {
				t.Fatalf("cell %d,%d is not empty", y, x)
			}
		}
	}
}
