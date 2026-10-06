package store

import (
	"context"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/reader"
)

func (in *Ingestor) commit(ctx context.Context, b reader.Batch) error {
	var entries []entry.Entry

	for _, rec := range b.Records {
		entries = append(entries, in.g.Add(rec)...)
	}

	entries = append(entries, in.g.Due(time.Now())...)

	meta := in.meta(b)
	if len(entries) == 0 && !in.needCommit(b) {
		in.note(b, entries)
		in.merge(b, meta)

		return nil
	}

	for i := range entries {
		entries[i].Session = in.src.Session
		entries[i].Source = in.src.ID
		entries[i].Received = time.Now()
	}

	if _, err := in.st.Commit(ctx, in.src.ID, meta, entries); err != nil {
		return err
	}

	in.note(b, entries)
	in.merge(b, meta)

	return nil
}

// meta picks the checkpoint that may be committed safely: the end of the
// last emitted entry, or the start of a new generation after rotation.
func (in *Ingestor) meta(b reader.Batch) CommitMeta {
	gen, pos, ok := in.g.Safe()

	switch {
	case ok && gen == b.Generation:
	case b.Rotated:
		gen, pos = b.Generation, 0
	default:
		gen, pos = in.src.Generation, in.src.Position
	}

	return CommitMeta{
		Generation: gen, Identity: b.Identity, Position: pos,
		Size: b.Size, State: b.State, Lost: b.Lost,
	}
}

// needCommit reports whether the batch changed something worth storing even
// with no new entries: a rotation, counted loss, or a state change.
func (in *Ingestor) needCommit(b reader.Batch) bool {
	return b.Rotated || b.Lost > 0 || b.State != in.src.State
}

func (in *Ingestor) merge(b reader.Batch, meta CommitMeta) {
	in.src.State = b.State
	in.src.Generation = meta.Generation
	in.src.Position = meta.Position
	in.src.Size = meta.Size
}
