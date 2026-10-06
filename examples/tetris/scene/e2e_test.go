package scene

import (
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
	"github.com/chinmay-sawant/ownframe/examples/tetris/store"
)

// pump ticks until ok holds, or fails after five seconds.
func pump(t *testing.T, s *Scene, clock *testClock, step *fakeStepper, ok func() bool) {
	t.Helper()

	deadline := time.Now().Add(5 * time.Second)
	for !ok() {
		if time.Now().After(deadline) {
			t.Fatal("the condition did not hold in time")
		}

		step.advance = 1
		tickAt(t, s, clock, time.Second/60)
		time.Sleep(time.Millisecond)
	}
}

func TestEndToEndWithMemoryStore(t *testing.T) {
	st, err := store.OpenMemory()
	if err != nil {
		t.Fatal(err)
	}

	step := &fakeStepper{}
	m := NewCore(11)
	s, clock := newTestScene(t, m, NewWorker(st), Options{Stepper: step})
	defer func() {
		if err := s.Close(); err != nil {
			t.Error(err)
		}
	}()

	click(t, s, "btn-scores")
	pump(t, s, clock, step, func() bool { return s.hist.msg == "NO SCORES YET" })

	click(t, s, "btn-demo")
	pump(t, s, clock, step, func() bool { return len(s.hist.entries) > 0 })

	if !hasText(s, "DEMO SCORES") {
		t.Fatal("the demo page is not labelled")
	}

	click(t, s, "btn-close")
	press(t, s, "enter")
	pump(t, s, clock, step, func() bool { return s.frame.Phase == game.PhaseRunning })

	press(t, s, "arrowleft")
	press(t, s, "space")
	pump(t, s, clock, step, func() bool { return s.frame.Score > 0 })

	m.g.Phase = game.PhaseOver
	m.watch()
	pump(t, s, clock, step, func() bool { return s.status == "SCORE SAVED" })

	click(t, s, "btn-scores")
	pump(t, s, clock, step, func() bool { return len(s.hist.entries) >= 1 })

	click(t, s, "btn-close")
	click(t, s, "btn-theme")
	pump(t, s, clock, step, func() bool { return s.status == "THEME SAVED" })

	if !s.dark {
		t.Fatal("the theme did not switch")
	}
}
