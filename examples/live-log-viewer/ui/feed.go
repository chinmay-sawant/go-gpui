package ui

import "context"

// Feed is the data source behind the viewer. The example wires it to the
// store package; tests use a fake.
type Feed interface {
	// Sources lists every known source with its state and entry count.
	Sources(ctx context.Context) ([]SourceInfo, error)
	// Page returns one keyset page ordered oldest to newest. BeforeID
	// pages backward, AfterID forward, both zero starts at newest.
	Page(ctx context.Context, q Query) (PageResult, error)
	// Detail returns one entry with its full multiline message.
	Detail(ctx context.Context, id int64) (Detail, error)
	// Tail returns entries newer than afterID, ordered oldest first, and
	// the total number newer than afterID.
	Tail(ctx context.Context, afterID int64, limit int) ([]Entry, int, error)
	// Export writes one bounded result set to path and returns the count.
	Export(ctx context.Context, q Query, path string) (int, error)
	// Settings loads and saves the persistent view settings.
	Settings(ctx context.Context) (Settings, error)
	SaveSettings(ctx context.Context, s Settings) error
}

// SourceInfo is one followed source as the sidebar shows it.
type SourceInfo struct {
	Key   string
	Name  string
	State string
	Count int
}
