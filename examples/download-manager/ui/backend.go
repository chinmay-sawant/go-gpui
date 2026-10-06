// Package ui draws the download manager page: the queue table, the job
// details, the history pager, and the light/dark theme. It owns the model
// the template prints and the retained display operations a tick repaints,
// but no transfer or storage code. wire.go adapts the core packages to the
// Backend interface.
package ui

import "context"

// Backend is the data and transfer surface the UI drives. Every method
// returns without touching SQL, the disk, or the network; workers answer
// through Poll instead.
type Backend interface {
	// Start starts workers. It returns at once.
	Start(ctx context.Context)
	// Info describes the backend mode and storage location.
	Info() Info
	// Add queues one job. It does not wait on SQL, the disk, or the network.
	Add(req AddRequest) error
	// Control applies pause, resume, cancel, retry, or remove.
	Control(cmd Control) error
	// Active asks for a fresh active set; it arrives from Poll.
	Active()
	// Page asks for one history page; the answer arrives from Poll.
	Page(req PageRequest)
	// Summary asks for the counts; the answer arrives from Poll.
	Summary(gen uint64)
	// Poll returns up to budget worker results without waiting.
	Poll(budget int) []Update
	// Dark returns the persisted theme and SetDark stores it.
	Dark() (bool, error)
	SetDark(dark bool) error
	// Close stops workers within the documented shutdown budget.
	Close() error
}

// Info is the backend mode and storage location the header shows.
type Info struct {
	Dummy   bool
	DataDir string
}

// HistoryPage is the number of durable history rows per page.
const HistoryPage = 50

// PageRequest asks for one history page. Before is the opaque keyset cursor
// the previous response returned; empty means the first page. The response
// echoes Gen so the UI can discard a stale answer.
type PageRequest struct {
	Gen    uint64
	Filter Filter
	Before string
	Limit  int
}
