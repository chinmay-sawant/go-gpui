package store

import (
	"context"
	"database/sql"
	"time"
)

func pruneTx(ctx context.Context, db *sql.DB, r Retention) (Pruned, error) {
	var out Pruned

loop:
	for out.Batches < pruneBatches {
		if r.MaxAge > 0 {
			n, err := deleteWhere(ctx, db, `received_ns < ?`,
				time.Now().Add(-r.MaxAge).UnixNano())
			if err != nil {
				return out, err
			}

			if n > 0 {
				out.Entries += n
				out.Batches++

				continue
			}
		}

		if r.MaxRows <= 0 && r.MaxBytes <= 0 {
			break
		}

		count, bytes, err := entryStats(ctx, db)
		if err != nil {
			return out, err
		}

		var n int64

		switch {
		case r.MaxRows > 0 && count > r.MaxRows:
			n, err = deleteOldest(ctx, db, min64(pruneBatch, count-r.MaxRows))
		case r.MaxBytes > 0 && bytes > r.MaxBytes:
			n, err = deleteOldest(ctx, db, pruneBatch)
		default:
			break loop
		}

		if err != nil {
			return out, err
		}

		if n == 0 {
			break
		}

		out.Entries += n
		out.Batches++
	}

	if err := db.QueryRowContext(ctx,
		`SELECT COALESCE(MAX(id), 0) FROM entries`).Scan(&out.Newest); err != nil {
		return out, err
	}

	return out, nil
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	}

	return b
}
