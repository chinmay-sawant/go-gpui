package app

import (
	"github.com/chinmay-sawant/go-gpui/examples/teams/store"
)

const (
	// DefaultWidth and DefaultHeight are the size of a newly opened window.
	DefaultWidth  = 1400
	DefaultHeight = 900

	// MinWidth and MinHeight are the smallest frame the screen will draw.
	MinWidth  = 820
	MinHeight = 600
)

// Option configures New.
type Option func(*config)

// config holds the settings an option can change.
type config struct {
	dbPath string
}

// WithDB keeps the state in the SQLite database at path. The default is an
// in-memory database, so a caller that passes no option never touches disk.
func WithDB(path string) Option {
	if path == "" {
		path = store.Memory
	}

	return func(c *config) {
		c.dbPath = path
	}
}
