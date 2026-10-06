package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

// Retention bounds stored history. A zero field means unlimited.
type Retention struct {
	MaxRows  int64
	MaxBytes int64
	MaxAge   time.Duration
}

// DefaultRetention keeps a week of history, capped at 200,000 rows and
// 256 MiB of raw record bytes.
func DefaultRetention() Retention {
	return Retention{
		MaxRows:  200_000,
		MaxBytes: 256 << 20,
		MaxAge:   7 * 24 * time.Hour,
	}
}

// Pruned reports what one Prune call removed.
type Pruned struct {
	Entries int64
	Batches int
	Newest  entry.EntryID
}

// pruneBatches bounds one Prune call so maintenance never stalls an
// interaction. The next call continues where this one stopped.
const pruneBatches = 64

const pruneBatch = int64(1000)

// Prune deletes history in bounded batches, oldest first.
func (s *Store) Prune(ctx context.Context) (Pruned, error) {
	r := s.retention

	return runJob(s, ctx, func(ctx context.Context, db *sql.DB) (Pruned, error) {
		return pruneTx(ctx, db, r)
	})
}

func pruneTx(ctx context.Context, db *sql.DB, r Retention) (Pruned, error) {
	var out Pruned

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

		over, err := overLimit(ctx, db, r)
		if err != nil {
			return out, err
		}

		if !over {
			break
		}

		n, err := deleteOldest(ctx, db)
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
