package ui

import (
	"context"
)

// Page implements Feed with keyset slicing. Total counts every matching
// entry, except for a tail query (AfterID set), where it counts entries
// above the cursor.
func (f *fakeFeed) Page(ctx context.Context, q Query) (PageResult, error) {
	if f.delay > 0 {
		if err := f.wait(ctx); err != nil {
			return PageResult{}, err
		}
	}

	f.mu.Lock()
	defer f.mu.Unlock()

	f.pages++
	all := f.matched(q)
	var out PageResult

	switch {
	case q.Oldest:
		out.Entries = firstN(all, q.Limit)
		out.HasNewer = len(all) > len(out.Entries)
	case q.BeforeID > 0:
		idx := lowerBound(all, q.BeforeID)
		older := all[:idx]
		out.HasNewer = true

		if len(older) > q.Limit {
			out.Entries = older[len(older)-q.Limit:]
			out.HasOlder = true
		} else {
			out.Entries = older
		}
	case q.AfterID > 0:
		idx := upperBound(all, q.AfterID)
		newer := all[idx:]
		out.HasOlder = true
		out.Total = len(newer)

		if len(newer) > q.Limit {
			out.Entries = newer[:q.Limit]
			out.HasNewer = true
		} else {
			out.Entries = newer
		}
	default:
		if len(all) > q.Limit {
			out.Entries = all[len(all)-q.Limit:]
			out.HasOlder = true
		} else {
			out.Entries = all
		}
	}

	if out.Total == 0 {
		out.Total = len(all)
	}

	return out, nil
}
