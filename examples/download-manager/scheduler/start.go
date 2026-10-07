package scheduler

import (
	"context"
	"time"
)

// timeNow is the default clock; Options.Now replaces it.
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
	ctx, e.cancel = context.WithCancel(ctx)
	workers := e.opts.Workers
	e.mu.Unlock()

	for i := 0; i < workers; i++ {
		e.wg.Add(1)

		go e.workerLoop(ctx)
	}
}

// Close cancels every transfer and waits for the workers within
// Options.ShutdownBudget. It is safe twice.
func (e *Engine) Close(ctx context.Context) error {
	e.mu.Lock()
	if e.closed {
		e.mu.Unlock()

		return nil
	}

	e.closed = true
	cancel := e.cancel
	e.mu.Unlock()

	if cancel != nil {
		cancel()
	}

	deadline := timeNow().Add(e.opts.ShutdownBudget)
	if d, ok := ctx.Deadline(); ok && d.Before(deadline) {
		deadline = d
	}

	done := make(chan struct{})

	go func() {
		e.wg.Wait()
		close(done)
	}()

	timer := time.NewTimer(time.Until(deadline))
	defer timer.Stop()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return ErrShutdownTimeout
	}
}
