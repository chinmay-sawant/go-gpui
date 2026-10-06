package scheduler

import (
	"context"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
	"github.com/chinmay-sawant/ownframe/examples/download-manager/transfer"
)

// run performs one job and reports the result.
func (e *Engine) run(parent context.Context, l *live) {
	ctx, cancel := context.WithCancel(parent)

	e.mu.Lock()
	l.cancel = cancel
	l.job.State = domain.StateRunning
	l.job.Attempts++
	l.job.Error = ""
	l.job.UpdatedAt = e.opts.Now()
	job := l.job
	e.mu.Unlock()

	e.persistAndPublish(job)

	var last time.Time

	report := func(p transfer.Progress) {
		e.onProgress(l, p, &last)
	}

	out, err := e.opts.Transport.Download(ctx, e.request(job), report)
	cancel()

	e.finish(l, out, err)
}

// onProgress updates memory, publishes one coalesced event, and checkpoints
// the bytes at most once per ProgressEvery.
func (e *Engine) onProgress(l *live, p transfer.Progress, last *time.Time) {
	e.mu.Lock()
	l.job.Done = p.Done
	l.job.Total = p.Total
	job := l.job
	now := e.opts.Now()
	due := now.Sub(*last) >= e.opts.ProgressEvery

	if due {
		*last = now
	}
	e.mu.Unlock()

	e.out.publish(Event{Kind: EventProgress, Job: job, At: now})

	if due {
		_ = e.opts.Store.Checkpoint(context.Background(), job.ID, p.Done, p.Total, now)
	}
}
