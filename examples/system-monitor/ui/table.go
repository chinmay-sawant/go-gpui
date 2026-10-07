package ui

import "strings"

// rowsPerPage bounds the rendered table rows independently of the total
// process count.
const rowsPerPage = 50

// sortKey names one process table column.
type sortKey string

// The sortable columns.
const (
	sortCPU  sortKey = "cpu"
	sortMem  sortKey = "mem"
	sortPID  sortKey = "pid"
	sortName sortKey = "name"
)

// table keeps a frozen process snapshot plus the query, sort, and page over
// it. A new snapshot waits in pending until refresh adopts it, so paging a
// snapshot never shifts rows underneath the user.
type table struct {
	shown   ProcSnapshot
	pending *ProcSnapshot
	query   string
	key     sortKey
	desc    bool
	page    int
}

// newTable returns a table sorted by CPU, heaviest first.
func newTable() *table {
	return &table{key: sortCPU, desc: true, page: 1}
}

// offer stashes the newest snapshot for an explicit refresh.
func (t *table) offer(s ProcSnapshot) {
	t.pending = &s
}

// hasNew reports whether a snapshot waits for refresh.
func (t *table) hasNew() bool {
	return t.pending != nil
}

// refresh adopts the offered snapshot. It reports false when none waits.
func (t *table) refresh() bool {
	if t.pending == nil {
		return false
	}

	t.shown = *t.pending
	t.pending = nil
	t.clamp()

	return true
}

// setQuery resets to page one when the filter changes.
func (t *table) setQuery(q string) {
	q = strings.TrimSpace(q)
	if q == t.query {
		return
	}

	t.query = q
	t.page = 1
}

// sortBy selects a column; repeating a column flips its direction. Every
// change resets to page one.
func (t *table) sortBy(k sortKey) {
	if k == t.key {
		t.desc = !t.desc
	} else {
		t.key = k
		t.desc = k == sortCPU || k == sortMem
	}

	t.page = 1
}
