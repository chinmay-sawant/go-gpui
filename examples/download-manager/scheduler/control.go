package scheduler

import (
	"context"
	"fmt"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
)

// ErrUnknownJob reports an ID the engine never saw.
var ErrUnknownJob = fmt.Errorf("scheduler: unknown job")

// Pause moves a queued or running job to paused. A running transfer is
// cancelled and its partial file is kept.
func (e *Engine) Pause(ctx context.Context, id string) error {
	e.mu.Lock()
	l, ok := e.jobs[id]
	if !ok {
		e.mu.Unlock()

		return fmt.Errorf("%w: %s", ErrUnknownJob, id)
	}

	job := l.job
	switch job.State {
	case domain.StateQueued:
		job.State = domain.StatePaused
		job.UpdatedAt = e.opts.Now()
		l.job = job
		e.removeQueuedLocked(id)
		e.mu.Unlock()

		return e.saveAndPublish(ctx, job)
	case domain.StateRunning:
		l.pausing = true
		if l.cancel != nil {
			l.cancel()
		}
		e.mu.Unlock()

		return nil
	default:
		e.mu.Unlock()

		return domain.Check(job.State, domain.StatePaused)
	}
}

// Resume returns a paused job to the queue.
func (e *Engine) Resume(ctx context.Context, id string) error {
	return e.requeue(ctx, id, domain.StatePaused)
}

// Retry returns a failed, cancelled, or completed job to the queue. The
// partial file, if any, is resumed when the validators still match.
func (e *Engine) Retry(ctx context.Context, id string) error {
	e.mu.Lock()
	l, ok := e.jobs[id]
	if !ok {
		e.mu.Unlock()

		return fmt.Errorf("%w: %s", ErrUnknownJob, id)
	}

	from := l.job.State
	e.mu.Unlock()

	switch from {
	case domain.StateFailed, domain.StateCancelled, domain.StateCompleted:
		return e.requeue(ctx, id, from)
	default:
		return domain.Check(from, domain.StateQueued)
	}
}
