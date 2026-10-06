package store

import (
	"context"
	"database/sql"
	"time"
)

type job struct {
	ctx  context.Context
	fn   func(context.Context, *sql.DB) (any, error)
	resp chan jobResult
}

type jobResult struct {
	val any
	err error
}

// work is the only goroutine that touches db, so transactions never race.
func (s *Store) work() {
	defer s.wg.Done()

	for {
		select {
		case <-s.stop:
			return
		case j := <-s.jobs:
			v, err := j.fn(j.ctx, s.db)
			j.resp <- jobResult{val: v, err: err}
		}
	}
}

// runJob sends fn to the worker and waits for its result. Every public
// method goes through here, so a cancelled context, a query deadline, or a
// closed store never leaves a call blocked.
func runJob[T any](s *Store, ctx context.Context, fn func(context.Context, *sql.DB) (T, error)) (T, error) {
	var zero T

	if s == nil || s.closed.Load() {
		return zero, ErrClosed
	}

	d := s.timeout
	if d <= 0 {
		d = 5 * time.Second
	}

	ctx, cancel := context.WithTimeout(ctx, d)
	defer cancel()

	j := &job{
		ctx:  ctx,
		resp: make(chan jobResult, 1),
		fn: func(c context.Context, db *sql.DB) (any, error) {
			return fn(c, db)
		},
	}

	select {
	case s.jobs <- j:
	case <-ctx.Done():
		return zero, ctx.Err()
	case <-s.stop:
		return zero, ErrClosed
	}

	select {
	case r := <-j.resp:
		if r.err != nil {
			return zero, r.err
		}

		v, ok := r.val.(T)
		if !ok {
			return zero, nil
		}

		return v, nil
	case <-ctx.Done():
		return zero, ctx.Err()
	case <-s.stop:
		return zero, ErrClosed
	}
}
