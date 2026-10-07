package parser

import (
	"time"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

// Due emits the pending entry when it has waited longer than FlushAfter.
// The Ingestor calls it after every read so an idle source still shows its
// newest line instead of holding it until the next record.
func (g *Grouper) Due(now time.Time) []entry.Entry {
	if g.pending == nil || g.opts.FlushAfter <= 0 {
		return nil
	}

	if now.Sub(g.pendingAt) < g.opts.FlushAfter {
		return nil
	}

	return []entry.Entry{g.emit()}
}
