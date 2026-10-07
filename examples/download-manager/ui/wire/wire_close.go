package wire

import (
	"context"
	"errors"
	"time"
)

// Close stops the worker and the engine within the documented budget: one
// second for the worker, then a six-second context for the engine's own
// five-second transfer shutdown. Later calls return nil.
func (b *Backend) Close() error {
	b.mu.Lock()
	if b.closed {
		b.mu.Unlock()

		return nil
	}

	b.closed = true
	cancel := b.cancel
	b.cancel = nil
	b.mu.Unlock()

	if cancel != nil {
		cancel()

		select {
		case <-b.done:
		case <-time.After(workerBudget):
		}
	}

	ctx, cancelClose := context.WithTimeout(context.Background(), closeBudget)
	defer cancelClose()

	var errs []error

	if b.eng != nil {
		if err := b.eng.Close(ctx); err != nil {
			errs = append(errs, err)
		}
	}

	if b.release != nil {
		b.release()
		b.release = nil
	}

	if b.store != nil {
		if err := b.store.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	if b.fixture != nil {
		b.fixture.Close()
		b.fixture = nil
	}

	return errors.Join(errs...)
}
