package domain

import "testing"

// TestStateClasses checks the terminal and active groupings.
func TestStateClasses(t *testing.T) {
	if !StateCompleted.Terminal() || !StateFailed.Terminal() || !StateCancelled.Terminal() {
		t.Error("a terminal state reported non-terminal")
	}

	if StateRunning.Terminal() || StateQueued.Terminal() {
		t.Error("an active state reported terminal")
	}

	for _, s := range []State{StateQueued, StateRunning, StatePaused} {
		if !s.Active() {
			t.Errorf("%s should be active", s)
		}
	}

	if StateCompleted.Active() {
		t.Error("completed reported active")
	}
}
