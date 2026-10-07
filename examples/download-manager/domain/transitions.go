package domain

// Terminal reports whether no transition leaves s.
func (s State) Terminal() bool {
	return s == StateCompleted || s == StateFailed || s == StateCancelled
}

// Active reports whether s claims the destination file.
func (s State) Active() bool {
	return s == StateQueued || s == StateRunning || s == StatePaused
}

// transitions is the allowed state graph. Paused goes back to queued, never
// straight to running, so one path re-enters the worker pool.
var transitions = map[State][]State{
	StateQueued:    {StateRunning, StateCancelled},
	StateRunning:   {StatePaused, StateCompleted, StateFailed, StateCancelled},
	StatePaused:    {StateQueued, StateCancelled, StateFailed},
	StateFailed:    {StateQueued, StateCancelled},
	StateCancelled: {StateQueued},
	StateCompleted: {StateQueued},
}

// CanTransition reports whether from may move to to. A state never moves to
// itself.
func CanTransition(from, to State) bool {
	for _, next := range transitions[from] {
		if next == to {
			return true
		}
	}

	return false
}

// Check reports a transition error when from may not move to to.
func Check(from, to State) error {
	if CanTransition(from, to) {
		return nil
	}

	return &TransitionError{From: from, To: to}
}
