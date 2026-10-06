package scene

import (
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
)

func TestFocusLossClearsInputAndPauses(t *testing.T) {
	m := &fakeModel{f: Frame{Phase: game.PhaseRunning}}
	focus := false
	step := &fakeStepper{}
	s, clock := newTestScene(t, m, nil, Options{
		Stepper: step,
		Focused: func() bool { return focus },
	})

	tickAt(t, s, clock, time.Second/60)

	if m.clears != 1 || m.pauses != 1 {
		t.Fatalf("clears = %d pauses = %d, want 1 and 1", m.clears, m.pauses)
	}

	if step.resets != 1 {
		t.Fatalf("resets = %d, want 1", step.resets)
	}

	focus = true
	tickAt(t, s, clock, time.Second/60)

	if m.resumes != 1 {
		t.Fatalf("resumes = %d, want 1", m.resumes)
	}
}

func TestDrainBudget(t *testing.T) {
	m := &fakeModel{}
	st := &fakeStore{}
	s, clock := newTestScene(t, m, st, Options{Stepper: &fakeStepper{}})

	for i := 0; i < 10; i++ {
		st.push(Result{ID: st.settingsID, Kind: ResultSettings})
	}

	tickAt(t, s, clock, time.Second/60)
	if len(m.configs) != 4 {
		t.Fatalf("first drain applied %d, want 4", len(m.configs))
	}

	tickAt(t, s, clock, time.Second/60)
	if len(m.configs) != 8 {
		t.Fatalf("second drain applied %d, want 8", len(m.configs))
	}

	tickAt(t, s, clock, time.Second/60)
	if len(m.configs) != 10 {
		t.Fatalf("third drain applied %d, want 10", len(m.configs))
	}
}
