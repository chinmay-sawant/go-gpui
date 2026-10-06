package store

import (
	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

// Source returns the stored source row this Ingestor follows.
func (in *Ingestor) Source() entry.Source { return in.src }

// SetPaused pauses ingestion. A replayable file simply lags; a live stream
// keeps producing and counts what it cannot hand over.
func (in *Ingestor) SetPaused(v bool) { in.paused.Store(v) }

// Paused reports the current pause state.
func (in *Ingestor) Paused() bool { return in.paused.Load() }

// Stats returns a snapshot of this source's counters.
func (in *Ingestor) Stats() IngestStats {
	in.mu.Lock()
	defer in.mu.Unlock()

	return in.stats
}

// Close releases the reader. Cancel the Run context and wait for Run to
// return before closing.
func (in *Ingestor) Close() error {
	if !in.closed.CompareAndSwap(false, true) {
		return nil
	}

	if in.r == nil {
		return nil
	}

	return in.r.Close()
}
