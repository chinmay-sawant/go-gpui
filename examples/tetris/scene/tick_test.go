package scene

import (
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
)

func TestTickRunsFixedSteps(t *testing.T) {
	m := &fakeModel{}
	step := &fakeStepper{}
	s, clock := newTestScene(t, m, nil, Options{Stepper: step})

	tickAt(t, s, clock, time.Second/60)

	step.advance = 3
	tickAt(t, s, clock, time.Second/60)

	if m.steps != 3 {
		t.Fatalf("steps = %d, want 3", m.steps)
	}

	if step.resets != 0 {
		t.Fatalf("resets = %d, want 0", step.resets)
	}
}

func TestStallPausesInsteadOfCatchingUp(t *testing.T) {
	m := &fakeModel{f: Frame{Phase: game.PhaseRunning}}
	step := &fakeStepper{}
	s, clock := newTestScene(t, m, nil, Options{Stepper: step})

	tickAt(t, s, clock, time.Second/60)

	step.advance = 5
	tickAt(t, s, clock, game.StallLimit+time.Second)

	if m.pauses != 1 {
		t.Fatalf("pauses = %d, want 1", m.pauses)
	}

	if !s.auto {
		t.Fatal("a stall pause must be automatic")
	}

	if step.resets != 1 {
		t.Fatalf("resets = %d, want 1", step.resets)
	}

	if m.steps != 0 {
		t.Fatalf("steps after a stall = %d, want 0", m.steps)
	}
}
