// Package store keeps download-manager jobs in SQLite through
// database/sql and modernc.org/sqlite. The database lives under
// os.UserConfigDir()/ownframe/download-manager unless a caller passes an
// explicit directory. One worker goroutine and one connection serialize
// every statement. Jobs and meta rows are durable; nothing here is UI state.
package store

import (
	"database/sql"
	"errors"
	"os"
	"path/filepath"
	"sync"

	_ "modernc.org/sqlite"
)

// DefaultDir returns the per-user data directory for this example.
func DefaultDir() (string, error) {
	base, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(base, "ownframe", "download-manager"), nil
}

// Store is one open database plus its serialized worker.
type Store struct {
	db      *sql.DB
	dir     string
	memory  bool
	ops     chan func()
	done    chan struct{}
	closeMu sync.Mutex
	closed  bool
}

// Open opens or creates the database in dir, creates dir when needed, and
// runs the pending migrations. It does not seed data and takes no instance
// lock; call Lock for that.
func Open(dir string) (*Store, error) { return open(dir, false) }

// OpenMemory opens a temporary in-memory database for temporary mode. The
// data lasts until Close and takes no lock.
func OpenMemory() (*Store, error) { return open("", true) }

// open is the shared constructor.
func open(dir string, memory bool) (*Store, error) {
	if !memory && dir == "" {
		return nil, errors.New("store: empty data directory")
	}

	return nil, errors.New("store: not implemented yet")
}

// Close stops the worker and closes the connection. It is safe twice.
func (s *Store) Close() error { return nil }
