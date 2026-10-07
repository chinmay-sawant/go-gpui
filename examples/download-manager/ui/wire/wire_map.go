package wire

import (
	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
	"github.com/chinmay-sawant/ownframe/examples/download-manager/ui"
)

// mapState maps the core state onto the printable one.
func mapState(s domain.State) ui.State {
	switch s {
	case domain.StateQueued:
		return ui.StateQueued
	case domain.StateRunning:
		return ui.StateRunning
	case domain.StatePaused:
		return ui.StatePaused
	case domain.StateCompleted:
		return ui.StateCompleted
	case domain.StateFailed:
		return ui.StateFailed
	case domain.StateCancelled:
		return ui.StateCancelled
	default:
		return ui.StateQueued
	}
}

// totalOf prefers the expected length and falls back to the observed one.
func totalOf(job domain.Job) int64 {
	if job.Expected > 0 {
		return job.Expected
	}

	return job.Total
}
