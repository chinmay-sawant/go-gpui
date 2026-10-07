package domain

import "testing"

// TestTransitions walks the allowed edges and a few refused ones.
func TestTransitions(t *testing.T) {
	allowed := []struct{ from, to State }{
		{StateQueued, StateRunning},
		{StateQueued, StateCancelled},
		{StateRunning, StatePaused},
		{StateRunning, StateCompleted},
		{StateRunning, StateFailed},
		{StateRunning, StateCancelled},
		{StatePaused, StateQueued},
		{StatePaused, StateCancelled},
		{StateFailed, StateQueued},
		{StateFailed, StateCancelled},
		{StateCancelled, StateQueued},
		{StateCompleted, StateQueued},
	}

	for _, edge := range allowed {
		if !CanTransition(edge.from, edge.to) {
			t.Errorf("want %s -> %s allowed", edge.from, edge.to)
		}

		if err := Check(edge.from, edge.to); err != nil {
			t.Errorf("Check(%s, %s) = %v", edge.from, edge.to, err)
		}
	}

	refused := []struct{ from, to State }{
		{StateQueued, StateCompleted},
		{StateQueued, StatePaused},
		{StateRunning, StateQueued},
		{StatePaused, StateRunning},
		{StateCompleted, StateFailed},
		{StateCancelled, StateCompleted},
	}

	for _, edge := range refused {
		if CanTransition(edge.from, edge.to) {
			t.Errorf("want %s -> %s refused", edge.from, edge.to)
		}

		var te *TransitionError
		if err := Check(edge.from, edge.to); !errorsAs(err, &te) {
			t.Errorf("Check(%s, %s) = %v, want TransitionError", edge.from, edge.to, err)
		}
	}

	for _, s := range States {
		if !s.Valid() {
			t.Errorf("%s not valid", s)
		}
	}

	if State("nope").Valid() {
		t.Error("unknown state reported valid")
	}
}
