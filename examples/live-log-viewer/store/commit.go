package store

import (
	"context"
	"database/sql"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

// CommitMeta carries the reader state that rides with a batch. Position is
// the next unread offset or record number; an empty Identity, a zero
// HeadHash, or an empty State keeps the stored value.
type CommitMeta struct {
	Generation int64
	Identity   string
	HeadHash   uint64
	HeadLen    int64
	Position   int64
	Size       int64
	State      entry.State
	Lost       int64
}

// Position describes the last stored entry of a source, so an Ingestor can
// rewind over a partial record after a restart.
type Position struct {
	Generation int64
	Offset     int64
	Bytes      int
	Partial    bool
	OK         bool
}

// Commit inserts a batch and advances the source checkpoint in one
// transaction, so a restart can never replay past a gap. A replayed row is
// ignored; a stored partial record is replaced by its completed version.
func (s *Store) Commit(ctx context.Context, src entry.SourceID, m CommitMeta, entries []entry.Entry) (int, error) {
	n, err := runJob(s, ctx, func(ctx context.Context, db *sql.DB) (int, error) {
		return commitTx(ctx, db, src, m, entries)
	})
	if err != nil {
		return n, err
	}

	if s.checkpointEvery > 0 && s.commits.Add(1)%int64(s.checkpointEvery) == 0 {
		_ = s.Checkpoint(context.WithoutCancel(ctx))
	}

	return n, nil
}
