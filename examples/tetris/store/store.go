// Package store keeps the tetris scores, settings, replays, and resumable
// snapshots in one SQLite database under the user config directory. One
// serialized worker owns one connection; the UI loop never touches SQL.
// Scores, settings, and snapshots are durable. There is no disposable
// telemetry table.
package store

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
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
	busyTimeoutMS = 5000
)

// Memory is the in-memory database path used with OpenMemory.
const Memory = ":memory:"

// Errors callers switch on.
var (
	ErrNewerSchema = errors.New("store: database schema is newer than this build")
	ErrClosed      = errors.New("store: closed")
	ErrBusy        = errors.New("store: queue is full")
	ErrInvalid     = errors.New("store: invalid record")
	ErrNotFound    = errors.New("store: not found")
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

// DefaultDir returns os.UserConfigDir()/ownframe/tetris.
func DefaultDir() (string, error) {
	base, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("store: config dir: %w", err)
	}

	return filepath.Join(base, "ownframe", AppDir), nil
}
