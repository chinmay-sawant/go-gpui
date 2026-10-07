package ui

// next moves one page forward within the frozen snapshot.
func (t *table) next() {
	if t.page < t.pages() {
		t.page++
	}
}

// prev moves one page back within the frozen snapshot.
func (t *table) prev() {
	if t.page > 1 {
		t.page--
	}
}

// pages returns the page count of the filtered snapshot.
func (t *table) pages() int {
	return pageCount(len(t.filtered()), rowsPerPage)
}

// clamp keeps the page inside the filtered snapshot, so a shrinking last
// page never shows empty.
func (t *table) clamp() {
	if p := t.pages(); t.page > p {
		t.page = p
	}

	if t.page < 1 {
		t.page = 1
	}
}

// find returns the frozen row for a process identity.
func (t *table) find(id string) (Process, bool) {
	for _, p := range t.shown.Procs {
		if p.ID == id {
			return p, true
		}
	}

	return Process{}, false
}
