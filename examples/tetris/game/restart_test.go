package game

import (
	"testing"
	"time"
)

func TestRestartClearsEveryTransientValue(t *testing.T) {
	g := New(1)
	g.Board = solidBoard(0, PieceJ)
	g.Start()
	g.Score, g.Lines, g.Level = 500, 12, 2
	g.Elapsed = time.Minute
	g.Steps = 99
	oldID := g.ID

	ev := g.Restart()

	if !haveEvent(ev, EventSpawn) {
		t.Fatal("restart did not spawn")
	}

	expectPhase(t, g, PhaseRunning)

	if g.ID == oldID {
		t.Fatal("restart kept the old game ID")
	}

	if g.Score != 0 || g.Lines != 0 || g.Level != 1 {
		t.Fatalf("score %d lines %d level %d", g.Score, g.Lines, g.Level)
	}

	if g.Pieces != 0 || g.Elapsed != 0 || g.Steps != 0 {
		t.Fatalf("pieces %d elapsed %v steps %d", g.Pieces, g.Elapsed, g.Steps)
	}

	if g.fall != 0 || g.lock != 0 || g.resets != 0 || g.grounded {
		t.Fatal("timers or the grounded flag survived the restart")
	}

	for y := 0; y < Rows; y++ {
		for x := 0; x < Cols; x++ {
			if g.Board[y][x] != 0 {
				t.Fatalf("cell %d,%d survived the restart", y, x)
			}
		}
	}

	if len(g.Next) != Preview || !g.Piece.Valid() {
		t.Fatalf("next %v piece %v", g.Next, g.Piece)
	}
}

func TestRestartFromPauseRunsAgain(t *testing.T) {
	g := New(1)
	g.Start()
	g.Apply(ActionPause)
	expectPhase(t, g, PhasePaused)

	ev := g.Apply(ActionRestart)

	expectPhase(t, g, PhaseRunning)

	if !haveEvent(ev, EventSpawn) {
		t.Fatal("restart from pause did not spawn")
	}
}

func TestStartAfterTopOutRestarts(t *testing.T) {
	g := New(1)
	g.Board = solidBoard(0, PieceJ)
	g.Start()
	expectPhase(t, g, PhaseOver)

	g.Apply(ActionStart)

	expectPhase(t, g, PhaseRunning)
}
