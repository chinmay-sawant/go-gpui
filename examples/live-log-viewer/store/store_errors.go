package store

import (
	"errors"
	"time"
)

// ErrClosed is returned when a call races a closed store.
var ErrClosed = errors.New("store: closed")

// ErrSchemaNewer reports a database written by a newer build. Open leaves
// that file untouched.
var ErrSchemaNewer = errors.New("store: database schema is newer than this build")

// ErrForeign reports a database file that is not this application's.
var ErrForeign = errors.New("store: not a live-log-viewer database")

// ErrCorrupt reports a file that SQLite cannot read.
var ErrCorrupt = errors.New("store: database is not readable")

// ErrReadOnly reports a database location that cannot be written.
var ErrReadOnly = errors.New("store: database directory is not writable")

const (
	dbName       = "live-log-viewer.db"
	appID        = "live-log-viewer"
	openTimeout  = 10 * time.Second
	sqliteExtra  = "?_pragma=busy_timeout(5000)&_pragma=foreign_keys(1)&_pragma=synchronous(NORMAL)"
	sqliteMemory = "file::memory:" + sqliteExtra
)
