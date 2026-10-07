package storage

import "context"

// foldBatch folds one bounded batch of old raw rows into aggregates and
// deletes exactly the rows it folded, in one transaction. The delete is keyed
// by the batch's highest ID, so a row inserted while the batch runs is never
// removed by it.
func (s *Store) foldBatch(ctx context.Context, cutoff int64) (int64, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer func() { _ = tx.Rollback() }()

	var (
		maxID int64
		count int64
	)

	if err := tx.QueryRowContext(ctx, `
SELECT COALESCE(MAX(id), 0), COUNT(*) FROM
	(SELECT id FROM metric_history WHERE ts_ns < ? ORDER BY id LIMIT ?)`,
		cutoff, BatchRows).Scan(&maxID, &count); err != nil {
		return 0, err
	}

	if count == 0 {
		return 0, nil
	}

	bucket := AggregateBucket.Nanoseconds()

	if _, err := tx.ExecContext(ctx, `
INSERT INTO metric_aggregate(session_id, metric, device, bucket_ns, min_value, max_value, avg_value, count)
SELECT session_id, metric, device, (ts_ns / ?) * ? AS bucket_ns, MIN(value), MAX(value), AVG(value), COUNT(*)
FROM metric_history
WHERE id <= ? AND ts_ns < ?
GROUP BY session_id, metric, device, bucket_ns
ON CONFLICT(session_id, metric, device, bucket_ns) DO UPDATE SET
	min_value = MIN(metric_aggregate.min_value, excluded.min_value),
	max_value = MAX(metric_aggregate.max_value, excluded.max_value),
	avg_value = (metric_aggregate.avg_value * metric_aggregate.count + excluded.avg_value * excluded.count)
		/ (metric_aggregate.count + excluded.count),
	count = metric_aggregate.count + excluded.count`,
		bucket, bucket, maxID, cutoff); err != nil {
		return 0, err
	}

	res, err := tx.ExecContext(ctx, `DELETE FROM metric_history WHERE id <= ? AND ts_ns < ?`, maxID, cutoff)
	if err != nil {
		return 0, err
	}

	deleted, _ := res.RowsAffected()

	return deleted, tx.Commit()
}
