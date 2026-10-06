package storage

import (
	"context"
	"time"
)

// Append writes metric rows for a session. A duplicate row (same session,
// metric, device, and timestamp) is ignored rather than overwriting, so a
// retried batch cannot corrupt history. The second return counts rows the
// store kept.
func (s *Store) Append(ctx context.Context, sessionID int64, rows []Row) (int, error) {
	return 0, errNotImplemented
}

// History returns raw rows in ID order, oldest first.
func (s *Store) History(ctx context.Context, q Query) ([]Point, error) {
	return nil, errNotImplemented
}

// Aggregates returns downsampled buckets in bucket order.
func (s *Store) Aggregates(ctx context.Context, q Query) ([]Aggregate, error) {
	return nil, errNotImplemented
}

// Retain folds raw rows older than RawRetention into aggregates, deletes the
// folded rows in bounded batches, drops aggregates older than
// AggregateRetention, and trims both tables to their caps. It is safe to run
// while sessions are recording and while the UI reads history.
func (s *Store) Retain(ctx context.Context, now time.Time) (RetainResult, error) {
	return RetainResult{}, errNotImplemented
}

// CountRows returns how many raw rows a session holds.
func (s *Store) CountRows(ctx context.Context, sessionID int64) (int64, error) {
	return 0, errNotImplemented
}
