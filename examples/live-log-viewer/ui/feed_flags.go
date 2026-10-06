package ui

import (
	"context"
	"fmt"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/store"
)

// pageFlags decides the boundary flags from the query direction.
func pageFlags(q Query, p store.Page, entries []Entry) (bool, bool) {
	switch {
	case q.Oldest:
		return false, p.More
	case q.AfterID > 0:
		return true, p.More
	case q.BeforeID > 0:
		newer := p.HighWater > 0

		if len(entries) > 0 {
			newer = entry.EntryID(entries[len(entries)-1].ID) < p.HighWater
		}

		return p.More, newer
	}

	return p.More, false
}

// count totals the matching entries, from a cursor when one is given. The
// query's MaxID cap stays, so a frozen browse totals only its own range.
func (f *StoreFeed) count(ctx context.Context, cq store.Query, cursor entry.EntryID) (int, error) {
	cq.Cursor = cursor

	n, err := f.st.Count(ctx, cq)

	return int(n), err
}

// Detail loads one entry with its full message.
func (f *StoreFeed) Detail(ctx context.Context, id int64) (Detail, error) {
	if id <= 0 {
		return Detail{}, fmt.Errorf("ui: bad entry id %d", id)
	}

	cq := store.Query{
		Cursor: entry.EntryID(id - 1),
		MaxID:  entry.EntryID(id),
	}

	p, err := f.st.Page(ctx, cq, store.PageOptions{Limit: 1, Ascending: true})
	if err != nil {
		return Detail{}, err
	}

	if len(p.Entries) == 0 {
		return Detail{}, fmt.Errorf("entry %d is no longer retained", id)
	}

	e := p.Entries[0]

	return Detail{Entry: f.toEntry(e), Lines: splitLines(e.Message)}, nil
}
