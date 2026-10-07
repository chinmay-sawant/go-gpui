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
