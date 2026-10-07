package store

import (
	"context"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

// shutdownBudget bounds the final flush and commit when a Run context is
// cancelled.
const shutdownBudget = 2 * time.Second

// finish stores whatever was buffered when the run stopped: an unterminated
// final line becomes a partial entry, and its checkpoint is committed so a
// restart replaces it with the completed line. Best effort within the
// shutdown budget.
func (in *Ingestor) finish() {
	ctx, cancel := context.WithTimeout(context.Background(), shutdownBudget)
	defer cancel()

	var entries []entry.Entry

	for _, rec := range in.r.Flush() {
		entries = append(entries, in.g.Add(rec)...)
	}

	entries = append(entries, in.g.Flush()...)
	if len(entries) == 0 {
		return
	}

	for i := range entries {
		entries[i].Session = in.src.Session
		entries[i].Source = in.src.ID
		entries[i].Received = time.Now()
	}

	gen, pos, ok := in.g.Safe()
	if !ok {
		return
	}

	_, _ = in.st.Commit(ctx, in.src.ID, CommitMeta{
		Generation: gen, Position: pos, Identity: in.src.Identity,
		Size: in.src.Size, State: in.src.State,
	}, entries)
}
