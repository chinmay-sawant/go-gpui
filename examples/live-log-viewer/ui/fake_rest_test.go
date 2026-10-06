package ui

import (
	"context"
	"fmt"
	"os"
)

// Detail returns one entry by ID.
func (f *fakeFeed) Detail(_ context.Context, id int64) (Detail, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	for _, e := range f.entries {
		if e.ID == id {
			return Detail{Entry: e, Lines: splitLines(e.Text)}, nil
		}
	}

	return Detail{}, fmt.Errorf("entry %d not found", id)
}

// Sources returns the sidebar snapshot.
func (f *fakeFeed) Sources(context.Context) ([]SourceInfo, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	out := make([]SourceInfo, len(f.sources))
	copy(out, f.sources)

	for i := range out {
		out[i].Count = len(f.entries)
	}

	return out, nil
}

// Export writes the matching range as CSV.
func (f *fakeFeed) Export(_ context.Context, q Query, path string) (int, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.exports++
	all := f.matched(q)
	if q.Limit > 0 && len(all) > q.Limit {
		all = all[:q.Limit]
	}

	var b []byte

	for _, e := range all {
		b = append(b, fmt.Sprintf("%d,%s,%s\n", e.ID, e.Severity, e.Text)...)
	}

	if err := os.WriteFile(path, b, 0o600); err != nil {
		return 0, err
	}

	return len(all), nil
}

// Settings returns the persisted view state.
func (f *fakeFeed) Settings(context.Context) (Settings, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	return f.settings, nil
}

// SaveSettings stores the view state.
func (f *fakeFeed) SaveSettings(_ context.Context, s Settings) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.settings = s

	return nil
}

func firstN(es []Entry, n int) []Entry {
	if n <= 0 || len(es) <= n {
		return es
	}

	return es[:n]
}

func lowerBound(es []Entry, id int64) int {
	for i, e := range es {
		if e.ID >= id {
			return i
		}
	}

	return len(es)
}

func upperBound(es []Entry, id int64) int {
	for i, e := range es {
		if e.ID > id {
			return i
		}
	}

	return len(es)
}
