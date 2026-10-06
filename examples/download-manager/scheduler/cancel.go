package scheduler

import (
	"context"
	"fmt"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
)

// Cancel stops an active job and marks it cancelled. A running transfer is
// cancelled and its partial file is kept for a later retry.
func (e *Engine) Cancel(ctx context.Context, id string) error {
	e.mu.Lock()
	l, ok := e.jobs[id]
	if !ok {
		e.mu.Unlock()

		return fmt.Errorf("%w: %s", ErrUnknownJob, id)
	}

	job := l.job
	switch job.State {
	case domain.StateRunning:
		l.cancelling = true
		if l.cancel != nil {
			l.cancel()
		}
		e.mu.Unlock()

		return nil
	case domain.StateQueued, domain.StatePaused:
		job.State = domain.StateCancelled
		job.Error = ""
		job.UpdatedAt = e.opts.Now()
		l.job = job
		e.removeQueuedLocked(id)
		e.mu.Unlock()

		return e.saveAndPublish(ctx, job)
	default:
		e.mu.Unlock()

		return domain.Check(job.State, domain.StateCancelled)
	}
}

// removeQueuedLocked drops id from the waiting list.
func (e *Engine) removeQueuedLocked(id string) {
	for i, queued := range e.queue {
		if queued == id {
			e.queue = append(e.queue[:i], e.queue[i+1:]...)

			if l := e.jobs[id]; l != nil {
				l.queued = false
			}

			return
		}
	}
}

// saveAndPublish persists one state change and announces it.
func (e *Engine) saveAndPublish(ctx context.Context, job domain.Job) error {
	err := e.opts.Store.SaveJob(ctx, job)
	e.out.publish(Event{Kind: EventState, Job: job, At: job.UpdatedAt})

	return err
}
