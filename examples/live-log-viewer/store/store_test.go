package store

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

func newStore(t *testing.T) *Store {
	t.Helper()

	st, err := Open(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { _ = st.Close() })

	return st
}

func newStoreOpts(t *testing.T, o Options) *Store {
	t.Helper()

	if o.Dir == "" {
		o.Dir = t.TempDir()
	}

	st, err := OpenWithOptions(o)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { _ = st.Close() })

	return st
}

func bg() context.Context { return context.Background() }

func allEntries(t *testing.T, st *Store, q Query) []entry.Entry {
	t.Helper()

	var out []entry.Entry

	cur := entry.EntryID(0)

	for {
		qq := q
		qq.Cursor = cur

		p, err := st.Page(bg(), qq, PageOptions{Limit: 500, Ascending: true})
		if err != nil {
			t.Fatal(err)
		}

		if len(p.Entries) == 0 {
			break
		}

		out = append(out, p.Entries...)
		cur = p.Entries[len(p.Entries)-1].ID

		if !p.More {
			break
		}
	}

	return out
}
