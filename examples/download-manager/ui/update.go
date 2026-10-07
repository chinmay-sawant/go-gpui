package ui

// UpdateKind says which field of an Update is set.
type UpdateKind uint8

// The five worker results the UI applies on the tick.
const (
	// UpdateProgress is one job's observed done count and speed.
	UpdateProgress UpdateKind = iota
	// UpdateActive replaces the bounded in-memory active set.
	UpdateActive
	// UpdateHistory carries one durable history page.
	UpdateHistory
	// UpdateSummary carries the queue counts.
	UpdateSummary
	// UpdateNotice carries a one-line message from the backend.
	UpdateNotice
)

// Update is one immutable worker result. The backend sends only the fields
// its Kind uses.
type Update struct {
	Kind    UpdateKind
	Gen     uint64
	Row     *Row
	Active  []Row
	Page    *PageResponse
	Summary *Summary
	Notice  string
}

// PageResponse is one durable history page. Next is the cursor for the
// following page, or empty on the last page. Total is the filtered count.
type PageResponse struct {
	Gen    uint64
	Filter Filter
	Rows   []Row
	Next   string
	Total  int
}

// Summary is the queue-wide job count by state.
type Summary struct {
	Queued    int
	Running   int
	Paused    int
	Completed int
	Failed    int
	Cancelled int
}

// Active returns the count in the three non-terminal states.
func (s Summary) Active() int {
	return s.Queued + s.Running + s.Paused
}
