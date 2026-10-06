package ui

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/store"
)

// newStoreApp opens a temporary real store with a small dummy fixture and
// wires it through the adapter.
func newStoreApp(t *testing.T) (*App, *StoreFeed, int64) {
	t.Helper()

	st, err := store.OpenWithOptions(store.Options{Temp: true})
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	if _, err := st.EnsureDummy(ctx, store.DummyOptions{Count: 500}); err != nil {
		t.Fatal(err)
	}

	stats, err := st.Stats(ctx)
	if err != nil {
		t.Fatal(err)
	}

	feed := NewStoreFeed(st)
	a, err := New(Options{
		Feed: feed, ExportDir: filepath.Join(st.Dir(), "exports"),
		Poll: 20 * time.Millisecond,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		_ = a.Close()
		_ = st.Close()
	})

	return a, feed, int64(stats.Newest)
}
