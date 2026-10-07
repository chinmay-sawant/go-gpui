package store

import (
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

func TestExportBounded(t *testing.T) {
	st := newStore(t)

	setup, err := st.EnsureDummy(bg(), DummyOptions{Count: 50})
	if err != nil {
		t.Fatal(err)
	}

	ex, err := st.Export(bg(), ExportOptions{
		Query: Query{Session: &setup.Session.ID}, MaxEntries: 10,
	})
	if err != nil {
		t.Fatal(err)
	}

	if ex.Entries != 10 || !ex.Truncated {
		t.Fatalf("export = %d entries, truncated=%v", ex.Entries, ex.Truncated)
	}

	if len(ex.Data) == 0 || ex.Oldest == 0 || ex.Newest < ex.Oldest {
		t.Fatalf("export data = %d bytes, oldest=%d newest=%d",
			len(ex.Data), ex.Oldest, ex.Newest)
	}
}

func TestExportIDs(t *testing.T) {
	st := newStore(t)

	setup, err := st.EnsureDummy(bg(), DummyOptions{Count: 50})
	if err != nil {
		t.Fatal(err)
	}

	entries := allEntries(t, st, Query{Session: &setup.Session.ID})
	if len(entries) < 3 {
		t.Fatalf("only %d entries", len(entries))
	}

	ids := []entry.EntryID{entries[0].ID, entries[1].ID}

	ex, err := st.Export(bg(), ExportOptions{IDs: ids})
	if err != nil {
		t.Fatal(err)
	}

	if ex.Entries != 2 || ex.Truncated {
		t.Fatalf("export = %d entries, truncated=%v", ex.Entries, ex.Truncated)
	}

	if ex.Oldest != entries[0].ID || ex.Newest != entries[1].ID {
		t.Fatalf("range %d..%d", ex.Oldest, ex.Newest)
	}
}
