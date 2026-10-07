package ui

// followState separates following the newest page from pause. Pausing stops
// display updates; the tail poll keeps running and counts unread entries.
// Resume never moves the reader: unread stays until an explicit jump.
type followState struct {
	Follow   bool
	Paused   bool
	Unread   int
	LastSeen int64
}

// Pause stops following and freezes the loaded page.
func (f *followState) Pause() {
	f.Paused = true
	f.Follow = false
}

// Resume lifts the pause without changing the page or the reading anchor.
func (f *followState) Resume() { f.Paused = false }

// SetFollow turns tail-following on or off; enabling clears the pause.
func (f *followState) SetFollow(on bool) {
	f.Follow = on
	if on {
		f.Paused = false
	}
}

// Live reports whether tail results should be applied to the page.
func (f *followState) Live() bool { return f.Follow && !f.Paused }

// Note records entries newer than the last applied one and the total count.
// Following applies them, so the unread count stays zero.
func (f *followState) Note(entries []Entry, total int) {
	for _, e := range entries {
		if e.ID > f.LastSeen {
			f.LastSeen = e.ID
		}
	}

	if f.Live() {
		f.Unread = 0
		return
	}

	f.Unread += total
	if f.Unread < 0 {
		f.Unread = 0
	}
}

// Seen raises the last-applied ID, as a page load or jump does.
func (f *followState) Seen(id int64) {
	if id > f.LastSeen {
		f.LastSeen = id
	}
}

// ClearUnread drops the unread count after a jump to the newest page.
func (f *followState) ClearUnread() { f.Unread = 0 }
