package store

import (
	"context"
	"errors"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
)

// ErrNotFound reports a missing job row.
var ErrNotFound = errors.New("store: job not found")

// ErrLocked reports that another process holds the instance lock.
var ErrLocked = errors.New("store: another instance holds the lock")

// errTODO marks the seam stubs still to be implemented.
var errTODO = errors.New("store: not implemented yet")

// SaveJob inserts or replaces a whole job row. The caller owns transition
// validation; the store writes what it is given.
func (s *Store) SaveJob(ctx context.Context, j domain.Job) error { return errTODO }

// Checkpoint writes only Done and Total, plus UpdatedAt. It never changes
// state, names, or validators, so it cannot clobber a concurrent edit.
func (s *Store) Checkpoint(ctx context.Context, id string, done, total int64, at time.Time) error {
	return errTODO
}

// Job reads one row.
func (s *Store) Job(ctx context.Context, id string) (domain.Job, error) {
	return domain.Job{}, errTODO
}

// ActiveJobs returns queued, running, and paused jobs, oldest first.
func (s *Store) ActiveJobs(ctx context.Context) ([]domain.Job, error) { return nil, errTODO }

// Cursor points at the last job of a page: its UpdatedAt plus its stable ID.
type Cursor struct {
	UpdatedAt time.Time
	ID        string
}

// Page is one chunk of history, newest first.
type Page struct {
	Jobs    []domain.Job
	Next    Cursor
	HasMore bool
}

// History pages terminal history by (UpdatedAt, ID) descending.
func (s *Store) History(ctx context.Context, after Cursor, limit int) (Page, error) {
	return Page{}, errTODO
}

// Counts summarizes every state without scanning for each progress event.
type Counts struct {
	ByState map[domain.State]int
	Total   int
}

// Summary returns the per-state counts. The UI calls it at a bounded rate.
func (s *Store) Summary(ctx context.Context) (Counts, error) { return Counts{}, errTODO }

// DeleteJob removes one row by ID.
func (s *Store) DeleteJob(ctx context.Context, id string) error { return errTODO }

// Cleanup deletes at most batch terminal history rows beyond keep, newest
// kept. It returns how many rows it removed.
func (s *Store) Cleanup(ctx context.Context, keep, batch int) (int, error) { return 0, errTODO }

// Report summarizes one Reconcile pass.
type Report struct {
	Recovered int // running became paused
	Completed int // a finalized file was adopted as completed
	Missing   int // completed lost its file, marked failed
	Reset     int // a partial longer than expected restarts
}

// Reconcile repairs the gap between rows and files after a crash. Running
// rows become paused or completed; completed rows with a missing file
// become failed.
func (s *Store) Reconcile(ctx context.Context) (Report, error) { return Report{}, errTODO }

// SeedDummy inserts jobCount deterministic rows once per seed version.
// Reopening never duplicates them and never rewrites later user edits.
func (s *Store) SeedDummy(ctx context.Context, jobCount int) error { return errTODO }

// Backup writes a consistent copy with SQLite VACUUM INTO, never a raw
// file copy of an open database.
func (s *Store) Backup(ctx context.Context, destPath string) error { return errTODO }

// Lock takes the single-instance lock beside the database. The returned
// release is idempotent. A stale lock from a dead process is reclaimed.
func (s *Store) Lock() (func(), error) { return nil, errTODO }
