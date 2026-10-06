package storage

import (
	"database/sql"
	"sync/atomic"
	"time"
)

// Options configures OpenWithOptions. Zero fields take the defaults.
type Options struct {
	// Dir is the data directory. Empty uses DefaultDir.
	Dir string
	// BusyTimeout bounds a wait on a locked database.
	BusyTimeout time.Duration
	// OpTimeout bounds each statement when the caller's context has no
	// closer deadline.
	OpTimeout time.Duration
	// ForceRollback skips WAL and keeps the rollback journal, which is the
	// documented fallback when WAL is unavailable.
	ForceRollback bool
	// Temp opens an in-memory database for a clearly labeled temporary
	// mode. Nothing persists.
	Temp bool
}

const (
	defaultBusyTimeout = 5 * time.Second
	defaultOpTimeout   = 5 * time.Second
)

// Store is one open database.
type Store struct {
	db        *sql.DB
	path      string
	journal   string
	opTimeout time.Duration
	temp      bool
	closed    atomic.Bool
}

// DefaultDir is defined in dir.go.

// Open opens or creates the database in dir, running migrations.
func Open(dir string) (*Store, error) {
	return OpenWithOptions(Options{Dir: dir})
}

// OpenWithOptions opens the store described by opts.
func OpenWithOptions(opts Options) (*Store, error) {
	return nil, errNotImplemented
}

// Close flushes WAL and closes the database. It is safe to call twice.
func (s *Store) Close() error { return errNotImplemented }

// Path returns the database file path, or ":memory:" in temporary mode.
func (s *Store) Path() string { return s.path }

// JournalMode returns the journal mode the connection actually uses, "wal"
// or "delete". A WAL request on storage that cannot support it falls back and
// reports the fallback here.
func (s *Store) JournalMode() string { return s.journal }
