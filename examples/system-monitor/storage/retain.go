package storage

import (
	"context"
	"time"
)

// maxFoldBatches bounds one Retain call. A backlog larger than this waits for
// the next call, so retention never holds the connection for long.
const maxFoldBatches = 20

// Retain folds raw rows older than RawRetention into five minute aggregates,
// deletes the folded rows in bounded batches, drops aggregates older than
// AggregateRetention, and trims both tables to their caps. It is safe to run
// while sessions record and while the UI reads history, because every step is
// one short transaction on the single connection.
//
// A wall clock step backwards moves the cutoff with it, which pauses age
// based deletion; the table caps still bound the file.
func (s *Store) Retain(ctx context.Context, now time.Time) (RetainResult, error) {
	if err := s.ready(); err != nil {
		return RetainResult{}, err
	}

	ctx, cancel := s.opCtx(ctx)
	defer cancel()

	var out RetainResult
	cutoff := now.Add(-RawRetention).UnixNano()

	for range maxFoldBatches {
		n, err := s.foldBatch(ctx, cutoff)
		if err != nil {
			return out, err
		}

		out.Folded += n
		if n < BatchRows {
			break
		}
	}

	aggCutoff := now.Add(-AggregateRetention).UnixNano()

	for range maxFoldBatches {
		n, err := s.dropAggregateBatch(ctx, aggCutoff)
		if err != nil {
			return out, err
		}

		out.Aggregates += n
		if n < BatchRows {
			break
		}
	}

	trimmed, err := s.trimCaps(ctx)
	if err != nil {
		return out, err
	}
	out.Trimmed = trimmed

	return out, nil
}
