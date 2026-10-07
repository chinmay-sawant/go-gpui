package scheduler

import (
	"context"
	"fmt"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
)

// requeue moves a job from a paused or terminal state back to queued.
func (e *Engine) requeue(ctx context.Context, id string, from domain.State) error {
	e.mu.Lock()
	l, ok := e.jobs[id]
	if !ok {
		e.mu.Unlock()

		return fmt.Errorf("%w: %s", ErrUnknownJob, id)
	}

	if l.job.State != from {
		e.mu.Unlock()

		return domain.Check(l.job.State, domain.StateQueued)
	}

	if len(e.queue) >= e.opts.Queue {
		e.mu.Unlock()

		return ErrQueueFull
	}

	l.job.State = domain.StateQueued
	l.job.Error = ""
	l.job.UpdatedAt = e.opts.Now()
	l.queued = true
	job := l.job
	e.queue = append(e.queue, id)
	e.mu.Unlock()

	e.signal()

	return e.saveAndPublish(ctx, job)
}
