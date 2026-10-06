package game

import "testing"

func TestIWallKickFromTheLeftWall(t *testing.T) {
	g := running(PieceI, RotR, -2, 5)

	if !g.rotate(true) {
		t.Fatal("the R to 2 kick should fit")
	}

	if g.Rot != Rot2 || g.X != 0 || g.Y != 5 {
		t.Fatalf("after kick rot %v at %d,%d", g.Rot, g.X, g.Y)
	}
}

func TestTFloorKick(t *testing.T) {
	g := running(PieceT, Rot0, 3, 18)

	if !g.rotate(true) {
		t.Fatal("the floor kick should fit")
	}

	if g.Rot != RotR || g.X != 2 || g.Y != 17 {
		t.Fatalf("after kick rot %v at %d,%d", g.Rot, g.X, g.Y)
	}
}

func TestRotationBlockedInAOneWideWell(t *testing.T) {
	g := New(1)
	g.Phase = PhaseRunning

	for y := 0; y < Rows; y++ {
		for x := 0; x < Cols; x++ {
			if x != 4 {
				g.Board[y][x] = Cell(PieceJ)
			}
		}
	}

	g.Piece = PieceI
	g.Rot = RotR
	g.X, g.Y = 2, 10

	if g.rotate(true) {
		t.Fatal("no kick should fit inside the well")
	}

	if g.Rot != RotR || g.X != 2 || g.Y != 10 {
		t.Fatalf("a blocked rotation changed state to %v %d,%d", g.Rot, g.X, g.Y)
	}
}

func TestKickTablesStartAtZero(t *testing.T) {
	for p := PieceI; p <= PieceL; p++ {
		for r := Rotation(0); r < 4; r++ {
			for _, cw := range []bool{true, false} {
				if k := kicks(p, r, cw); k[0] != (kick{}) {
					t.Fatalf("%v %v cw=%v starts at %+v", p, r, cw, k[0])
				}
			}
		}
	}
}
