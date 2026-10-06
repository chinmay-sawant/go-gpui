package store

import (
	"sync"
	"sync/atomic"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/parser"
	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/reader"
)

// IngestStats is a snapshot of one Ingestor. Lost counts records a live
// stream skipped because the store could not keep up; it is never silently
// discarded.
type IngestStats struct {
	Source     entry.SourceID
	State      entry.State
	Generation int64
	Position   int64
	Batches    int64
	Records    int64
	Bytes      int64
	Lost       int64
	Errors     int64
	LastError  string
	Lag        int64
	Done       bool
}

// Ingestor drives one reader into the store: read a bounded batch, parse and
// group it, commit the entries and the checkpoint together. Run it on its
// own goroutine; the UI polls Stats.
type Ingestor struct {
	st     *Store
	src    entry.Source
	pol    entry.Policy
	r      reader.Reader
	g      *parser.Grouper
	mu     sync.Mutex
	stats  IngestStats
	paused atomic.Bool
	closed atomic.Bool
}
