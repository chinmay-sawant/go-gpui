package ui

import "strings"

// State is the printable lifecycle state of one job. The values match the
// order of the states the core domain defines.
type State uint8

// The six states, in lifecycle order.
const (
	StateQueued State = iota
	StateRunning
	StatePaused
	StateCompleted
	StateFailed
	StateCancelled
)

// Symbol returns a shape that reads without color.
func (s State) Symbol() string {
	switch s {
	case StateQueued:
		return "…"
	case StateRunning:
		return "▼"
	case StatePaused:
		return "‖"
	case StateCompleted:
		return "✓"
	case StateFailed:
		return "✗"
	case StateCancelled:
		return "⊘"
	default:
		return "?"
	}
}

// Label returns the state name.
func (s State) Label() string {
	switch s {
	case StateQueued:
		return "Queued"
	case StateRunning:
		return "Running"
	case StatePaused:
		return "Paused"
	case StateCompleted:
		return "Completed"
	case StateFailed:
		return "Failed"
	case StateCancelled:
		return "Cancelled"
	default:
		return "Unknown"
	}
}

// Key returns the lowercase class name the template uses.
func (s State) Key() string {
	return strings.ToLower(s.Label())
}

// Active reports whether the job holds or waits for a worker slot.
func (s State) Active() bool {
	return s == StateQueued || s == StateRunning || s == StatePaused
}

// Terminal reports whether no transition leaves the state.
func (s State) Terminal() bool {
	return s == StateCompleted || s == StateFailed || s == StateCancelled
}
