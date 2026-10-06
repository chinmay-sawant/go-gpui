package ui

// Layout constants the window math and the template both rely on.
const (
	// PageLimit is the keyset page size.
	PageLimit = 200
	// RowH is the fixed summary row height in CSS pixels.
	RowH = 22
	// Overscan is the extra rows rendered above and below the viewport.
	Overscan = 8
	// SideW is the source sidebar width.
	SideW = 180
	// HeaderH is the pinned filter bar height.
	HeaderH = 38
	// StatusH is the pinned footer height.
	StatusH = 22
)

// Pager holds one loaded keyset page and the reading anchor. Entries are
// ordered oldest to newest and their IDs rise monotonically. While HWM is
// nonzero the page set is frozen at that high-water mark, so entries that
// arrive during a browse cannot move rows under the reader.
type Pager struct {
	Entries     []Entry
	Total       int
	HasOlder    bool
	HasNewer    bool
	HWM         int64
	AnchorID    int64
	AnchorDelta int
	Skipped     int
	Gen         uint64
}

// Load replaces the page with a worker result.
func (p *Pager) Load(r PageResult) {
	p.Entries = r.Entries
	p.Total = r.Total
	p.HasOlder = r.HasOlder
	p.HasNewer = r.HasNewer
	p.Skipped = r.Skipped
	p.Gen++
}

// Len returns the loaded entry count.
func (p *Pager) Len() int { return len(p.Entries) }

// Empty reports a page with no entries.
func (p *Pager) Empty() bool { return len(p.Entries) == 0 }

// OlderCursor is the BeforeID that pages one step toward older entries.
func (p *Pager) OlderCursor() int64 {
	if !p.HasOlder || len(p.Entries) == 0 {
		return 0
	}

	return p.Entries[0].ID
}

// NewerCursor is the AfterID that pages one step toward newer entries.
func (p *Pager) NewerCursor() int64 {
	if !p.HasNewer || len(p.Entries) == 0 {
		return 0
	}

	return p.Entries[len(p.Entries)-1].ID
}

// MaxID returns the query cap: the frozen high-water mark while browsing,
// zero when the page may follow live entries.
func (p *Pager) MaxID() int64 { return p.HWM }
