package ui

import (
	"context"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/store"
)

// StoreFeed adapts the core store to the Feed interface. It runs on the
// worker goroutine only; the UI loop never calls the store directly.
type StoreFeed struct {
	st       *store.Store
	settings string

	mu     sync.Mutex
	labels map[entry.SourceID]string
}

// NewStoreFeed wraps an open store. View settings live in a JSON file next
// to the database, because the store schema has no settings table.
func NewStoreFeed(st *store.Store) *StoreFeed {
	dir := st.Dir()
	if dir == "" {
		dir = os.TempDir()
	}

	return &StoreFeed{
		st:       st,
		settings: filepath.Join(dir, "ui-settings.json"),
		labels:   map[entry.SourceID]string{},
	}
}

// Sources lists every session's sources for the sidebar.
func (f *StoreFeed) Sources(ctx context.Context) ([]SourceInfo, error) {
	sessions, err := f.st.Sessions(ctx)
	if err != nil {
		return nil, err
	}

	var out []SourceInfo

	for _, sess := range sessions {
		srcs, err := f.st.Sources(ctx, sess.ID)
		if err != nil {
			return nil, err
		}

		for _, src := range srcs {
			out = append(out, f.sourceInfo(ctx, src))
		}
	}

	return out, nil
}

// sourceInfo maps one stored source and its entry count.
func (f *StoreFeed) sourceInfo(ctx context.Context, src entry.Source) SourceInfo {
	id := src.ID
	count, err := f.st.Count(ctx, store.Query{Source: &id})
	if err != nil {
		count = 0
	}

	name := src.Label
	if name == "" {
		name = filepath.Base(src.Path)
	}

	f.setLabel(id, name)

	return SourceInfo{
		Key:   strconv.FormatInt(int64(id), 10),
		Name:  name,
		State: string(src.State),
		Count: int(count),
	}
}
