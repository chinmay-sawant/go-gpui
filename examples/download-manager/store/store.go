// Package store keeps download-manager jobs in SQLite through database/sql
// and modernc.org/sqlite. The database lives under
// os.UserConfigDir()/ownframe/download-manager unless a caller passes an
// explicit directory.
//
// One worker goroutine and one physical connection serialize every
// statement, so callers never hold a transaction open across calls. Jobs
// and meta rows are durable; a progress checkpoint is a hint that is
// re-checked against the partial file after a restart.
package store

import (
	"database/sql"
	"errors"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// ErrClosed reports a use after Close.
var ErrClosed = errors.New("store: closed")

// queryTimeout bounds one statement when the caller sets no deadline.
const queryTimeout = 5 * time.Second

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
	journal string

	ops  chan func()
	quit chan struct{}
	done chan struct{}
	once sync.Once
}

// Open opens or creates the database file in dir, creates dir when needed,
// and runs the pending migrations. It does not seed data and takes no
// instance lock; call Lock for that.
func Open(dir string) (*Store, error) { return open(dir, false) }

// OpenMemory opens a temporary in-memory database for temporary mode. The
// data lasts until Close and takes no lock.
func OpenMemory() (*Store, error) { return open("", true) }
