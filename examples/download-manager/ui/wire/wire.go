// Package wire adapts the download-manager core packages (store, scheduler,
// transfer, fixture) to the ui.Backend interface. It owns no page state and
// draws nothing.
package wire

import (
	"context"
	"sync"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/fixture"
	"github.com/chinmay-sawant/ownframe/examples/download-manager/scheduler"
	"github.com/chinmay-sawant/ownframe/examples/download-manager/store"
	"github.com/chinmay-sawant/ownframe/examples/download-manager/ui"
)

// Config configures the backend.
type Config struct {
	// DataDir overrides the per-user storage directory. Empty means
	// store.DefaultDir(), and a directory that will not open falls back to
	// an in-memory database.
	DataDir string
	// Dummy selects the deterministic fake transport and seeds 100 jobs.
	Dummy bool
	// Fixture starts the local HTTP fixture service and logs its URLs.
	Fixture bool
}

// Backend implements ui.Backend over the core packages.
type Backend struct {
	store   *store.Store
	eng     *scheduler.Engine
	release func()
	fixture *fixture.Server
	dir     string
	dummy   bool

	cmds    chan command
	results chan ui.Update
	ctx     context.Context
	cancel  context.CancelFunc
	done    chan struct{}

	mu      sync.Mutex
	dark    bool
	speeds  map[string]speed
	dropped int
	closed  bool
}
