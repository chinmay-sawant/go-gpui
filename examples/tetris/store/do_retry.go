package store

import (
	"context"
	"database/sql"
)

// doRetry runs fn on the worker, retrying lock failures within the
// request deadline.
func (s *Store) doRetry(ctx context.Context, fn func(context.Context, *sql.DB) error) error {
	return s.do(ctx, func(ctx context.Context, db *sql.DB) error {
		return retryBusy(ctx, func() error { return fn(ctx, db) })
	})
}
