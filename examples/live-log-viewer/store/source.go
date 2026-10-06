package store

import (
	"context"
	"database/sql"
	"os"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

// SourceSpec describes a source to follow. Path is unique inside a session:
// a file path, or a key for a generated stream. FromEnd starts a file at its
// current size instead of replaying it from the first byte.
type SourceSpec struct {
	Session entry.SessionID
	Kind    string
	Path    string
	Label   string
	Seed    int64
	Rate    int
	Total   int64
	FromEnd bool
}

// AddSource inserts a source or refreshes the label of an existing one. The
// stored checkpoint survives a re-add.
func (s *Store) AddSource(ctx context.Context, spec SourceSpec) (entry.Source, error) {
	return runJob(s, ctx, func(ctx context.Context, db *sql.DB) (entry.Source, error) {
		var zero entry.Source

		kind := spec.Kind
		if kind == "" {
			kind = entry.KindFile
		}

		if kind == entry.KindFile && spec.Path == "" {
			return zero, sql.ErrNoRows
		}

		position := int64(0)
		if kind == entry.KindFile && spec.FromEnd {
			if fi, err := os.Stat(spec.Path); err == nil {
				position = fi.Size()
			}
		}

		now := time.Now().UnixNano()

		_, err := db.ExecContext(ctx,
			`INSERT INTO sources
			 (session_id, kind, path, label, position, total, state, seed, rate, created_ns, updated_ns)
			 VALUES(?, ?, ?, ?, ?, ?, 'idle', ?, ?, ?, ?)
			 ON CONFLICT(session_id, path) DO UPDATE SET label = excluded.label`,
			int64(spec.Session), kind, spec.Path, spec.Label,
			position, spec.Total, spec.Seed, spec.Rate, now, now)
		if err != nil {
			return zero, err
		}

		return sourceByPath(ctx, db, spec.Session, spec.Path)
	})
}
