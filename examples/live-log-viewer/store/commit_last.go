package store

import (
	"context"
	"database/sql"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

// LastPosition returns the last stored entry of a source.
func (s *Store) LastPosition(ctx context.Context, src entry.SourceID) (Position, error) {
	return runJob(s, ctx, func(ctx context.Context, db *sql.DB) (Position, error) {
		var (
			p     Position
			bytes int64
			part  int64
		)

		err := db.QueryRowContext(ctx,
			`SELECT generation, position, bytes, partial
			 FROM entries WHERE source_id = ? ORDER BY id DESC LIMIT 1`,
			int64(src)).Scan(&p.Generation, &p.Offset, &bytes, &part)
		if err == sql.ErrNoRows {
			return p, nil
		}

		if err != nil {
			return p, err
		}

		p.OK = true
		p.Bytes = int(bytes)
		p.Partial = part != 0

		return p, nil
	})
}
