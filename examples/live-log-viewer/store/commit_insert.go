package store

import (
	"context"
	"database/sql"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

// entryChunk bounds a multi-row insert. 200 rows times 16 columns stays far
// below the SQLite variable limit.
const entryChunk = 200

// insertEntries writes the rows with fresh source sequence numbers. It
// returns how many rows were inserted or replaced.
func insertEntries(ctx context.Context, tx *sql.Tx, entries []entry.Entry, startSeq int64) (int, error) {
	inserted, seq := 0, startSeq

	for start := 0; start < len(entries); start += entryChunk {
		end := start + entryChunk
		if end > len(entries) {
			end = len(entries)
		}

		query, args := insertSQL(entries[start:end], &seq)

		res, err := tx.ExecContext(ctx, query, args...)
		if err != nil {
			return inserted, err
		}

		n, _ := res.RowsAffected()
		inserted += int(n)
	}

	return inserted, nil
}
