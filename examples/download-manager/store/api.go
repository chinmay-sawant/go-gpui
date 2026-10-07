package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
)

// ErrNotFound reports a missing job row.
var ErrNotFound = errors.New("store: job not found")

// ErrLocked reports that another process holds the instance lock.
var ErrLocked = errors.New("store: another instance holds the lock")

// Counts summarizes every state without scanning the table per progress
// event.
type Counts struct {
	ByState map[domain.State]int
	Total   int
}

// Report summarizes one Reconcile pass.
type Report struct {
	Recovered int // running became paused
	Completed int // a finalized or complete file was adopted
	Missing   int // completed lost its file, marked failed
	Reset     int // a partial longer than expected restarts
	Renamed   int // a reused destination got a new name
}

// Add merges another report into r.
func (r *Report) Add(other Report) {
	r.Recovered += other.Recovered
	r.Completed += other.Completed
	r.Missing += other.Missing
	r.Reset += other.Reset
	r.Renamed += other.Renamed
}

// meta reads one meta value. Empty and false when the key is absent.
func (s *Store) meta(ctx context.Context, key string) (string, bool, error) {
	var value string

	err := s.db.QueryRowContext(ctx,
		`SELECT value FROM meta WHERE key = ?`, key).Scan(&value)

	if errors.Is(err, sql.ErrNoRows) {
		return "", false, nil
	}

	if err != nil {
		return "", false, err
	}

	return value, true, nil
}
