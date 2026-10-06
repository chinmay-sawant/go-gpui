package game

import "testing"

func TestBlockedSpawnTopsOut(t *testing.T) {
	g := New(1)
	g.Board = solidBoard(0, PieceJ)

	ev := g.Start()

	if !haveEvent(ev, EventTopOut) {
		t.Fatal("a blocked spawn reported no top-out")
	}

	expectPhase(t, g, PhaseOver)
}

func TestSpawnOnATwoRowStackIsFine(t *testing.T) {
	g := New(1)
	g.Board = solidBoard(2, PieceJ)

	ev := g.Start()

	if !haveEvent(ev, EventSpawn) {
		t.Fatal("a clear spawn reported no spawn event")
	}

	expectPhase(t, g, PhaseRunning)
}

func TestLockAboveTheBoardTopsOut(t *testing.T) {
	g := running(PieceI, Rot0, 3, -2)

	for x := 3; x <= 6; x++ {
		g.Board[0][x] = Cell(PieceJ)
	}

	ev := g.Apply(ActionHardDrop)

	if !haveEvent(ev, EventTopOut) {
		t.Fatal("a lock above the board reported no top-out")
	}

	expectPhase(t, g, PhaseOver)
}

func TestTopOutKeepsTheBoardForDisplay(t *testing.T) {
	g := New(1)
	g.Board = solidBoard(0, PieceJ)
	g.Start()

	if g.Board[19][0] == 0 {
		t.Fatal("the board was cleared on top-out")
	}
}
