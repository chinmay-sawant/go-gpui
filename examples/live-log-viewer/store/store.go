// Package store keeps the live log viewer's SQLite state: sessions, sources,
// ordered entries, and committed ingestion checkpoints. One worker goroutine
// serializes every database call through one connection. Open creates and
// migrates the database, EnsureDummy fills the demo fixture, Page and Export
// read bounded windows, and Ingestor drives readers into Commit.
//
// Durability: entries, sessions, sources, and checkpoints are durable; Commit
// inserts rows and advances the source checkpoint in one transaction. The
// Ingestor counters (Batches, Lag) are disposable telemetry and reset on
// restart; a source's Lost total is durable because it records real loss.
package store

import (
	"database/sql"
	"sync"
	"sync/atomic"
	"time"

	_ "modernc.org/sqlite"
)

// Memory is a database that lasts as long as the process. Tests use it, and
// two in-memory stores share nothing.
const Memory = ":memory:"

// Options configures OpenWithOptions. Dir is the data directory; Temp
// overrides it with a fresh directory under the OS temp root that Close
// removes. Retention limits pruning, and CheckpointEvery counts commits
// between passive WAL checkpoints.
type Options struct {
	Dir             string
	Temp            bool
	QueryTimeout    time.Duration
	Retention       Retention
	CheckpointEvery int
}

// Store is one open SQLite database with its worker.
type Store struct {
	db              *sql.DB
	dir             string
	path            string
	temp            bool
	journal         string
	timeout         time.Duration
	retention       Retention
	checkpointEvery int
	commits         atomic.Int64
	jobs            chan *job
	stop            chan struct{}
	wg              sync.WaitGroup
	closed          atomic.Bool
}
