package ui

import (
	"context"
	"fmt"
	"strconv"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/store"
)

// coreQuery maps the UI filters into a store query.
func (f *StoreFeed) coreQuery(q Query) (store.Query, error) {
	var cq store.Query

	cq.Text = q.Text
	if q.MinSev != "" {
		sev := entry.ParseSeverity(q.MinSev)
		cq.MinSeverity = &sev
	}

	if len(q.Sources) > 0 && q.Sources[0] != "" {
		id, err := strconv.ParseInt(q.Sources[0], 10, 64)
		if err != nil {
			return cq, fmt.Errorf("ui: bad source key %q: %w", q.Sources[0], err)
		}

		src := entry.SourceID(id)
		cq.Source = &src
	}

	cq.MaxID = entry.EntryID(q.MaxID)

	return cq, nil
}

// Page runs one keyset page. The store returns newest first; the UI wants
// oldest first, so the adapter reverses the slice.
func (f *StoreFeed) Page(ctx context.Context, q Query) (PageResult, error) {
	cq, err := f.coreQuery(q)
	if err != nil {
		return PageResult{}, err
	}

	opts := store.PageOptions{Limit: q.Limit}
	if opts.Limit <= 0 {
		opts.Limit = PageLimit
	}

	switch {
	case q.Oldest:
		opts.Ascending = true
	case q.AfterID > 0:
		opts.Ascending = true
		cq.Cursor = entry.EntryID(q.AfterID)
	case q.BeforeID > 0:
		cq.Cursor = entry.EntryID(q.BeforeID)
	}

	p, err := f.st.Page(ctx, cq, opts)
	if err != nil {
		return PageResult{}, err
	}

	out := PageResult{Expired: p.Expired}
	out.Entries = make([]Entry, 0, len(p.Entries))

	if opts.Ascending {
		for _, e := range p.Entries {
			out.Entries = append(out.Entries, f.toEntry(e))
		}
	} else {
		for i := len(p.Entries) - 1; i >= 0; i-- {
			out.Entries = append(out.Entries, f.toEntry(p.Entries[i]))
		}
	}

	out.HasOlder, out.HasNewer = pageFlags(q, p, out.Entries)

	if q.AfterID > 0 {
		out.Total, err = f.count(ctx, cq, entry.EntryID(q.AfterID))
	} else {
		out.Total, err = f.count(ctx, cq, 0)
	}

	return out, err
}
