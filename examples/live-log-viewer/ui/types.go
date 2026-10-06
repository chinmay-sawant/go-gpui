// Package ui is the live log viewer screen. It renders a compact source
// sidebar, a monospaced list of fixed-height summary rows, and an entry
// detail view. It keeps page data, retained display operations, and
// window-related state on the UI loop; a worker goroutine runs Feed calls
// and hands results back through a bounded channel that the tick drains.
// ownframe opens the window; this package does not.
package ui

import "time"

// Entry is one stored log entry in display-ready form.
type Entry struct {
	ID        int64
	Time      time.Time
	TimeRaw   string
	TimeOK    bool
	Source    string
	SourceKey string
	Severity  string
	Text      string
	Lines     int
	Bytes     int
	Truncated bool
	Partial   bool
	Lost      bool
}

// Query is one keyset page request. Limit caps the page at the store.
// FromID and ToID are inclusive export bounds; zero means unbounded.
// Oldest asks for the first page of the result set instead of the newest.
// MinSeverity is "" for every level or a severity name; the store returns
// entries at or above it.
type Query struct {
	Sources  []string
	MinSev   string
	Text     string
	BeforeID int64
	AfterID  int64
	MaxID    int64
	FromID   int64
	ToID     int64
	Limit    int
	Oldest   bool
}

// PageResult is one keyset page. Entries are ordered oldest to newest.
// Expired reports that the requested cursor fell below retained history.
type PageResult struct {
	Entries  []Entry
	HasOlder bool
	HasNewer bool
	Total    int
	Expired  bool
}

// Detail is one entry with its full bounded message.
type Detail struct {
	Entry
	Lines []string
}

// Settings is the persistent slice of the view state.
type Settings struct {
	Dark     bool
	Follow   bool
	Severity string
	Source   string
}
