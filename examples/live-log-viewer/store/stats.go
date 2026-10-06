package store

import (
	"context"
	"database/sql"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

// Stats is a cheap health snapshot for the UI.
type Stats struct {
	Dir           string
	Path          string
	Temporary     bool
	Journal       string
	SchemaVersion int
	Sessions      int64
	Sources       int64
	Entries       int64
	Bytes         int64
	Oldest        entry.EntryID
	Newest        entry.EntryID
}

// Stats reports row counts, retained byte total, and storage mode.
func (s *Store) Stats(ctx context.Context) (Stats, error) {
	return runJob(s, ctx, func(ctx context.Context, db *sql.DB) (Stats, error) {
		st := Stats{
			Dir: s.dir, Path: s.path, Temporary: s.temp,
			Journal: s.journal, SchemaVersion: schemaVersion,
		}

		err := db.QueryRowContext(ctx,
			`SELECT (SELECT count(*) FROM sessions),
			        (SELECT count(*) FROM sources),
			        count(*), COALESCE(SUM(bytes), 0),
			        COALESCE(MIN(id), 0), COALESCE(MAX(id), 0)
			 FROM entries`).
			Scan(&st.Sessions, &st.Sources, &st.Entries, &st.Bytes,
				&st.Oldest, &st.Newest)

		return st, err
	})
}

// Checkpoint folds the write-ahead log back into the database file. It runs
// on the worker, never on the UI loop, and is a no-op outside WAL mode.
func (s *Store) Checkpoint(ctx context.Context) error {
	_, err := runJob(s, ctx, func(ctx context.Context, db *sql.DB) (int, error) {
		if s.journal != "wal" {
			return 0, nil
		}

		var busy, log, moved int64

		err := db.QueryRowContext(ctx, `PRAGMA wal_checkpoint(PASSIVE)`).
			Scan(&busy, &log, &moved)

		return int(moved), err
	})

	return err
}
