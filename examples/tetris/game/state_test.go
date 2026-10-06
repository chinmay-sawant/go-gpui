package game

import (
	"testing"
	"time"
)

func TestStartPauseResume(t *testing.T) {
	g := New(1)
	expectPhase(t, g, PhaseReady)

	if !haveEvent(g.Start(), EventSpawn) {
		t.Fatal("start did not spawn")
	}

	expectPhase(t, g, PhaseRunning)

	if !haveEvent(g.Apply(ActionPause), EventPause) {
		t.Fatal("pause reported no event")
	}

	expectPhase(t, g, PhasePaused)

	steps := g.Steps
	g.Step(time.Second)

	if g.Steps != steps {
		t.Fatal("a paused game stepped")
	}

	if !haveEvent(g.Apply(ActionStart), EventPause) {
		t.Fatal("resume reported no event")
	}

	expectPhase(t, g, PhaseRunning)
}

func TestPauseOnReadyDoesNothing(t *testing.T) {
	g := New(1)

	if g.Apply(ActionPause) != nil {
		t.Fatal("pause on the ready screen reported an event")
	}

	expectPhase(t, g, PhaseReady)
}

func TestSoftDropMovesAndScoresOne(t *testing.T) {
	g := New(1)
	g.Start()
	y := g.Y

	g.Apply(ActionSoftDrop)

	if g.Y != y+1 || g.Score != 1 {
		t.Fatalf("soft drop moved to %d with score %d", g.Y, g.Score)
	}
}

func TestHardDropScoresTwoPerRow(t *testing.T) {
	g := New(1)
	g.Start()

	ev := g.Apply(ActionHardDrop)

	if !haveEvent(ev, EventHardDrop) || !haveEvent(ev, EventLock) {
		t.Fatal("hard drop missed its events")
	}

	if g.Pieces != 1 {
		t.Fatalf("pieces = %d, want 1", g.Pieces)
	}

	if g.Score != 38 {
		t.Fatalf("hard drop score = %d, want 38", g.Score)
	}
}

func TestGravityFallsOneRowPerSecondAtLevelOne(t *testing.T) {
	g := New(1)
	g.Start()
	y := g.Y

	g.Step(500 * time.Millisecond)

	if g.Y != y {
		t.Fatalf("half a second moved %d rows", g.Y-y)
	}

	g.Step(500 * time.Millisecond)

	if g.Y != y+1 {
		t.Fatalf("one second moved %d rows, want 1", g.Y-y)
	}
}

func TestHighLevelGravityFallsSeveralRowsPerStep(t *testing.T) {
	g := New(1)
	g.Start()
	g.Level = MaxLevel
	y := g.Y

	g.Step(180 * time.Millisecond)

	if g.Y != y+3 {
		t.Fatalf("high gravity moved %d rows, want 3", g.Y-y)
	}
}
