package game

import "testing"

func TestEveryPieceMeetsEveryWall(t *testing.T) {
	for p := PieceI; p <= PieceL; p++ {
		for r := Rotation(0); r < 4; r++ {
			cells := p.Cells(r)
			minX, maxX, maxY := cells[0].X, cells[0].X, cells[0].Y

			for _, c := range cells {
				minX = min(minX, c.X)
				maxX = max(maxX, c.X)
				maxY = max(maxY, c.Y)
			}

			g := running(p, r, -minX, 0)

			if g.collides(p, r, -minX, 0) {
				t.Fatalf("%v %v collides flush left", p, r)
			}

			if !g.collides(p, r, -minX-1, 0) {
				t.Fatalf("%v %v escapes the left wall", p, r)
			}

			x := Cols - 1 - maxX
			if g.collides(p, r, x, 0) {
				t.Fatalf("%v %v collides flush right", p, r)
			}

			if !g.collides(p, r, x+1, 0) {
				t.Fatalf("%v %v escapes the right wall", p, r)
			}

			y := Rows - 1 - maxY
			if g.collides(p, r, 0, y) {
				t.Fatalf("%v %v collides on the floor", p, r)
			}

			if !g.collides(p, r, 0, y+1) {
				t.Fatalf("%v %v falls through the floor", p, r)
			}
		}
	}
}

func TestCellsAboveTheBoardAreAllowed(t *testing.T) {
	g := New(1)

	if g.collides(PieceT, Rot0, 3, -3) {
		t.Fatal("a piece above the board must not collide")
	}
}

func TestStackBlocksTheFall(t *testing.T) {
	g := New(1)
	g.Board[19][4] = Cell(PieceJ)

	if g.collides(PieceO, Rot0, 3, 17) {
		t.Fatal("O at row 17 must fit above the stack")
	}

	if !g.collides(PieceO, Rot0, 3, 18) {
		t.Fatal("O at row 18 must rest on the stack")
	}
}

func TestLandYFindsTheRestRow(t *testing.T) {
	g := running(PieceT, Rot0, 3, 0)
	g.Board[19][4] = Cell(PieceJ)

	if got := g.landY(); got != 17 {
		t.Fatalf("landY = %d, want 17", got)
	}
}
