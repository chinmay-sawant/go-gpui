// Package store keeps the tetris scores, settings, replays, and resumable
// snapshots in one SQLite database under the user config directory. One
// serialized worker owns one connection; the UI loop never touches SQL.
package store

import (
	"errors"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/tetris/input"
)

// Storage layout and versions.
const (
	AppDir        = "tetris"
	DBName        = "tetris.db"
	SchemaVersion = 1
	SeedVersion   = 1
	DefaultLimit  = 20
	QueueSize     = 64
	QueryTimeout  = 2 * time.Second
)

// Errors callers switch on.
var (
	ErrNewerSchema = errors.New("store: database schema is newer than this build")
	ErrClosed      = errors.New("store: closed")
	ErrBusy        = errors.New("store: queue is full")
	ErrInvalid     = errors.New("store: invalid record")
)

// Settings is the durable control configuration.
type Settings struct {
	Keymap input.Keymap `json:"keymap"`
	Theme  string       `json:"theme"`
	Ghost  bool         `json:"ghost"`
}

// DefaultSettings returns the starting configuration.
func DefaultSettings() Settings {
	return Settings{Keymap: input.DefaultKeymap()}
}

// Store is one open database with one serialized worker.
type Store struct{}

// DefaultDir returns os.UserConfigDir()/ownframe/tetris.
func DefaultDir() (string, error) { return "", nil }

// Open opens the database in dir, creating it when needed. An empty dir
// uses DefaultDir.
func Open(dir string) (*Store, error) { return &Store{}, nil }

// OpenMemory opens a private in-memory database for tests.
func OpenMemory() (*Store, error) { return &Store{}, nil }

// Close drains queued work, checkpoints, and closes the database.
func (s *Store) Close() error { return nil }
