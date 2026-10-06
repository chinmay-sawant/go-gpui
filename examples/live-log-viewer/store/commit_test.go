package store

import (
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

func TestCommitAdvancesCheckpoint(t *testing.T) {
	st := newStore(t)

	setup, err := st.EnsureDummy(bg(), DummyOptions{Count: 30})
	if err != nil {
		t.Fatal(err)
	}

	src := setup.Sources[0]
	start := src.Position

	before, err := st.Count(bg(), Query{Source: &src.ID})
	if err != nil {
		t.Fatal(err)
	}

	e := entry.Entry{
		Session: setup.Session.ID, Source: src.ID, Position: start,
		Generation: 1, Severity: entry.Info, Message: "hello world",
		Bytes: 12, Received: time.Now(),
	}

	n, err := st.Commit(bg(), src.ID, CommitMeta{
		Generation: 1, Position: start + 12, Size: 4096, State: entry.StateLive,
	}, []entry.Entry{e})
	if err != nil {
		t.Fatal(err)
	}

	if n != 1 {
		t.Fatalf("inserted %d rows", n)
	}

	got, err := st.Source(bg(), src.ID)
	if err != nil {
		t.Fatal(err)
	}

	if got.Position != start+12 || got.State != entry.StateLive || got.Size != 4096 {
		t.Fatalf("source = %+v", got)
	}

	// A replayed batch must not duplicate the row.
	n2, err := st.Commit(bg(), src.ID, CommitMeta{
		Generation: 1, Position: start + 12, State: entry.StateLive,
	}, []entry.Entry{e})
	if err != nil {
		t.Fatal(err)
	}

	if n2 != 0 {
		t.Fatalf("replay inserted %d rows", n2)
	}

	after, err := st.Count(bg(), Query{Source: &src.ID})
	if err != nil {
		t.Fatal(err)
	}

	if after != before+1 {
		t.Fatalf("count %d -> %d, want %d", before, after, before+1)
	}
}
