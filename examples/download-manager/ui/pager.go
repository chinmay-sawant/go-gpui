package ui

// Pager walks the durable history one keyset page at a time. The cursor is
// opaque: the backend returns the one that follows. A filter change resets
// the walk, and every request carries a generation, so an answer for an
// older filter or page is discarded.
type Pager struct {
	gen     uint64
	filter  Filter
	limit   int
	cursors []string // before-cursor of each visited page
	pos     int
	next    string // cursor for the page after the current one
	rows    []Row
	total   int
	sel     string
	loading bool
}

// NewPager returns a pager for pages of limit rows. A limit below 1 uses
// HistoryPage.
func NewPager(limit int) *Pager {
	if limit < 1 {
		limit = HistoryPage
	}

	return &Pager{limit: limit, cursors: []string{""}, loading: true}
}

// Filter returns the active filter.
func (p *Pager) Filter() Filter { return p.filter }

// SetFilter switches the filter. It reports whether anything changed; a
// change drops the visited cursors and the rows, because the old walk does
// not describe the new result set. The selected job ID survives.
func (p *Pager) SetFilter(f Filter) bool {
	if f == p.filter {
		return false
	}

	p.filter = f
	p.cursors = []string{""}
	p.pos = 0
	p.next = ""
	p.rows = nil
	p.total = 0
	p.loading = true

	return true
}
