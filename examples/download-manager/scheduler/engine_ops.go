package scheduler

import (
	"context"
	"errors"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
)

var errMissingOption = errors.New("scheduler: Transport and Store are required")

// timeNow is the default clock.
var timeNow = time.Now

// Start launches the worker pool. Calling it twice is a no-op. A context
// that ends cancels every transfer and stops the pool.
func (e *Engine) Start(ctx context.Context) {
	e.mu.Lock()
	if e.started || e.closed {
		e.mu.Unlock()
		return
	}

	e.started = true
	workers := e.opts.Workers
	e.mu.Unlock()

	for i := 0; i < workers; i++ {
		e.wg.Add(1)
		go e.workerLoop(ctx)
	}
}

// Add validates the request, computes a unique destination, persists the
// queued job, and enqueues it. It returns the stored job.
func (e *Engine) Add(ctx context.Context, req AddRequest) (domain.Job, error) {
	return domain.Job{}, errNotImplemented
}

// Pause moves a queued or running job to paused. A running transfer is
// cancelled and its partial file is kept.
func (e *Engine) Pause(ctx context.Context, id string) error {
	return errNotImplemented
}

// Resume returns a paused job to the queue.
func (e *Engine) Resume(ctx context.Context, id string) error {
	return errNotImplemented
}

// Cancel stops an active job and marks it cancelled.
func (e *Engine) Cancel(ctx context.Context, id string) error {
	return errNotImplemented
}

// Retry returns a failed, cancelled, or completed job to the queue.
func (e *Engine) Retry(ctx context.Context, id string) error {
	return errNotImplemented
}

// Job returns one live job.
func (e *Engine) Job(id string) (domain.Job, bool) {
	return domain.Job{}, false
}

// Active returns the queued, running, and paused jobs, newest first.
func (e *Engine) Active() []domain.Job { return nil }

// Drain pops at most max events.
func (e *Engine) Drain(max int) []Event { return e.out.Drain(max) }

// Recover loads active jobs from the store and enqueues the queued ones.
func (e *Engine) Recover(ctx context.Context) error {
	return errNotImplemented
}

// Close cancels every transfer and waits for the workers within
// Options.ShutdownBudget. It is safe to call once; later calls return nil.
func (e *Engine) Close(ctx context.Context) error {
	return errNotImplemented
}

var errNotImplemented = errors.New("scheduler: not implemented yet")

// workerLoop is the pool body; it is replaced by the real implementation.
func (e *Engine) workerLoop(ctx context.Context) {
	defer e.wg.Done()

	<-ctx.Done()
}
