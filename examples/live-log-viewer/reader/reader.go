// Package reader tails log sources. A File follows one file through append,
// truncation, rotation, rename, and delete with bounded batches. The Stream
// generator backs the demo. Every reader is cancellable, reports its
// position, and never drops a record silently: a replayable file lags and a
// live stream counts each skipped record.
package reader

import (
	"context"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

// Batch is one bounded read. Position is the source's next unread offset or
// record number, and Lag counts bytes written but not read yet. Lost counts
// records a live stream skipped because the consumer fell behind.
type Batch struct {
	Records    []entry.RawRecord
	Generation int64
	Identity   string
	Position   int64
	Size       int64
	Lag        int64
	Lost       int64
	More       bool
	Rotated    bool
	Missing    bool
	Done       bool
	State      entry.State
}

// Reader yields bounded batches until the context is cancelled. Flush emits
// an unterminated final record without doing IO; the Ingestor calls it once
// on shutdown.
type Reader interface {
	Read(ctx context.Context) (Batch, error)
	Flush() []entry.RawRecord
	Replayable() bool
	Close() error
}
