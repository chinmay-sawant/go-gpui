package scheduler

import (
	"context"
	"errors"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
	"github.com/chinmay-sawant/ownframe/examples/download-manager/transfer"
)

// persistTimeout bounds one state write from a worker.
const persistTimeout = 5 * time.Second

// finish maps one transfer result to a state, persists it, and publishes
// the state event. Completion and failure events never drop.
func (e *Engine) finish(l *live, out transfer.Outcome, err error) {
	e.mu.Lock()
	l.running = false
	l.cancel = nil

	job := l.job
	pausing, cancelling := l.pausing, l.cancelling
	l.pausing, l.cancelling = false, false

	switch {
	case cancelling:
		job.State = domain.StateCancelled
		job.Error = ""
	case pausing:
		job.State = domain.StatePaused
		job.Error = ""
	case err == nil:
		job.State = domain.StateCompleted
		job.Error = ""
		job.Done = out.Bytes
		job.Total = out.Total
		if out.Total >= 0 {
			job.Expected = out.Total
		}
		job.ETag = out.Validators.ETag
		job.LastModified = out.Validators.LastModified
	case errors.Is(err, context.Canceled):
		job.State = domain.StatePaused
		job.Error = "interrupted"
	default:
		job.State = domain.StateFailed
		job.Error = cleanError(job, err)
	}

	job.UpdatedAt = e.opts.Now()
	l.job = job
	e.mu.Unlock()

	e.persistAndPublish(job)
}

// persistAndPublish writes the row and then announces it. A failed write
// is never hidden: the event carries a note and the in-memory job keeps it.
func (e *Engine) persistAndPublish(job domain.Job) {
	ctx, cancel := context.WithTimeout(context.Background(), persistTimeout)
	defer cancel()

	if err := e.opts.Store.SaveJob(ctx, job); err != nil {
		job.Error = joinNote(job.Error, "state not saved: "+err.Error())
	}

	e.out.publish(Event{Kind: EventState, Job: job, At: job.UpdatedAt})
}
