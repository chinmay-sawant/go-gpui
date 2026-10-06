package scheduler

import (
	"context"
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
	"github.com/chinmay-sawant/ownframe/examples/download-manager/store"
	"github.com/chinmay-sawant/ownframe/examples/download-manager/transfer"
)

// newTestEngine builds an engine over a temporary in-memory store.
func newTestEngine(t *testing.T, tr transfer.Transport, opts Options) (*Engine, *store.Store) {
	t.Helper()

	st, err := store.OpenMemory()
	if err != nil {
		t.Fatal(err)
	}

	opts.Transport = tr
	opts.Store = st
	opts.Now = time.Now

	if opts.Workers == 0 {
		opts.Workers = 2
	}

	if opts.ShutdownBudget == 0 {
		opts.ShutdownBudget = time.Second
	}

	eng, err := New(opts)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		_ = eng.Close(context.Background())
		_ = st.Close()
	})

	return eng, st
}

// add starts one job against dir.
func add(t *testing.T, eng *Engine, dir, url string) domain.Job {
	t.Helper()

	job, err := eng.Add(context.Background(), AddRequest{URL: url, Dir: dir})
	if err != nil {
		t.Fatal(err)
	}

	return job
}
