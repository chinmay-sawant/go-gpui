package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/parser"
	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/reader"
)

// seedSource generates count entries for one dummy source and stores them
// with the source checkpoint. The records come from the deterministic
// generator, so the same seed always yields the same fixture.
func seedSource(ctx context.Context, tx *sql.Tx, src entry.Source, count int, seed int64, pol entry.Policy) error {
	g := parser.NewGrouper(parser.Options{
		MaxLines: pol.MultilineLines,
		MaxBytes: pol.MultilineBytes,
	})

	var entries []entry.Entry

	seq := int64(1)

	for len(entries) < count {
		need := count - len(entries)
		recs := reader.DummyRecords(seed, src.Path, seq, need, pol.MaxRecord)

		for _, rec := range recs {
			seq = rec.Offset + 1
			entries = append(entries, g.Add(rec)...)
		}
	}

	entries = append(entries, g.Flush()...)

	for i := range entries {
		entries[i].Session = src.Session
		entries[i].Source = src.ID
		entries[i].Received = reader.DummyBase.Add(
			time.Duration(entries[i].Position) * 250 * time.Millisecond)
	}

	if _, err := insertEntries(ctx, tx, entries, 0); err != nil {
		return err
	}

	_, err := tx.ExecContext(ctx,
		`UPDATE sources SET position = ?, generation = 1, state = 'idle', updated_ns = ?
		 WHERE id = ?`, seq, time.Now().UnixNano(), int64(src.ID))

	return err
}
