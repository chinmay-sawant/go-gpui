package storage

import "context"

// dropAggregateBatch deletes one bounded batch of aggregates older than the
// cutoff.
func (s *Store) dropAggregateBatch(ctx context.Context, cutoff int64) (int64, error) {
	res, err := s.db.ExecContext(ctx, `
DELETE FROM metric_aggregate WHERE rowid IN
	(SELECT rowid FROM metric_aggregate WHERE bucket_ns < ? LIMIT ?)`,
		cutoff, BatchRows)
	if err != nil {
		return 0, err
	}

	n, _ := res.RowsAffected()

	return n, nil
}

// trimCaps removes the oldest rows when a table is over its cap. The cap wins
// over age retention: the file stays bounded even when the clock jumps or a
// recorder runs for weeks.
func (s *Store) trimCaps(ctx context.Context) (int64, error) {
	removed, err := s.trimTable(ctx,
		`SELECT COUNT(*) FROM metric_history`,
		`DELETE FROM metric_history WHERE id IN (SELECT id FROM metric_history ORDER BY id LIMIT ?)`,
		MaxRawRows)
	if err != nil {
		return removed, err
	}

	n, err := s.trimTable(ctx,
		`SELECT COUNT(*) FROM metric_aggregate`,
		`DELETE FROM metric_aggregate WHERE rowid IN (SELECT rowid FROM metric_aggregate ORDER BY bucket_ns LIMIT ?)`,
		MaxAggregateRows)

	return removed + n, err
}

// trimTable counts a table and deletes the oldest rows beyond a limit in
// bounded batches.
func (s *Store) trimTable(ctx context.Context, countSQL, deleteSQL string, limit int64) (int64, error) {
	var total int64
	if err := s.db.QueryRowContext(ctx, countSQL).Scan(&total); err != nil {
		return 0, err
	}

	removed := int64(0)

	for total-removed > limit {
		n := total - removed - limit
		if n > BatchRows {
			n = BatchRows
		}

		res, err := s.db.ExecContext(ctx, deleteSQL, n)
		if err != nil {
			return removed, err
		}

		affected, _ := res.RowsAffected()
		if affected == 0 {
			break
		}
		removed += affected
	}

	return removed, nil
}
