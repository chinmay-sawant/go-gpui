package store

import (
	"context"
	"database/sql"
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
