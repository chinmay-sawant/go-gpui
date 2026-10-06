package ui

// applyTail counts new entries and appends them when following the tail.
func (a *App) applyTail(o out) bool {
	if o.err != nil || o.after != a.tailAfter.Load() {
		return false
	}

	before := a.follow.Unread
	a.follow.Note(o.entries, o.total)
	added := 0

	if a.follow.Live() {
		added = a.appendTail(o.entries)

		if added > 0 {
			a.scrollBottom()
		}
	}

	a.tailAfter.Store(a.follow.LastSeen)

	return added > 0 || a.follow.Unread != before
}

// appendTail appends entries newer than the loaded page, trims the page to
// the page limit, and reports how many rows were added.
func (a *App) appendTail(entries []Entry) int {
	last := int64(0)
	if n := a.pager.Len(); n > 0 {
		last = a.pager.Entries[n-1].ID
	}

	added := 0
	for _, e := range entries {
		if e.ID <= last {
			continue
		}

		a.pager.Entries = append(a.pager.Entries, e)
		last = e.ID
		added++
	}

	if added == 0 {
		return 0
	}

	if a.pager.Total > 0 {
		a.pager.Total += added
	}

	if n := a.pager.Len(); n > PageLimit {
		a.pager.Entries = a.pager.Entries[n-PageLimit:]
		a.pager.HasOlder = true
	}

	a.pager.HasNewer = false

	return added
}
