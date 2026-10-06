package storage

import (
	"context"
	"database/sql"
)

func (s *Store) worker() {
	defer close(s.done)

	for req := range s.reqs {
		req.done <- s.run(req)
	}
}

// run applies the default deadline and executes fn.
func (s *Store) run(req request) error {
	ctx := req.ctx
	if ctx == nil {
		ctx = context.Background()
	}

	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc

		ctx, cancel = context.WithTimeout(ctx, s.timeout)
		defer cancel()
	}

	return req.fn(ctx, s.db)
}

// do sends fn to the worker and waits. The queue blocks when full; it never
// drops a request.
func (s *Store) do(ctx context.Context, fn func(context.Context, *sql.DB) error) error {
	if s.closed.Load() {
		return ErrClosed
	}

	if ctx == nil {
		ctx = context.Background()
	}

	req := request{ctx: ctx, fn: fn, done: make(chan error, 1)}

	select {
	case s.reqs <- req:
	case <-ctx.Done():
		return ctx.Err()
	case <-s.done:
		return ErrClosed
	}

	select {
	case err := <-req.done:
		return err
	case <-s.done:
		return ErrClosed
	}
}

// rollback undoes a transaction and returns the original error.
func rollback(tx *sql.Tx, err error) error {
	_ = tx.Rollback()

	return err
}
