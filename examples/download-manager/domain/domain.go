// Package domain holds the download-manager data model: the job states,
// their valid transitions, and the job record. It performs no I/O and
// depends only on the standard library.
package domain

// State is the lifecycle state of a download job.
type State string

// The six job states. A job starts queued and reaches one terminal state:
// completed, failed, or cancelled. Paused is the only state a user resumes
// from, and a failed or cancelled job returns to queued through retry.
const (
	StateQueued    State = "queued"
	StateRunning   State = "running"
	StatePaused    State = "paused"
	StateCompleted State = "completed"
	StateFailed    State = "failed"
	StateCancelled State = "cancelled"
)

// States lists the six states in lifecycle order.
var States = []State{
	StateQueued, StateRunning, StatePaused,
	StateCompleted, StateFailed, StateCancelled,
}

// Valid reports whether s is one of the six states.
func (s State) Valid() bool {
	for _, known := range States {
		if s == known {
			return true
		}
	}

	return false
}

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
