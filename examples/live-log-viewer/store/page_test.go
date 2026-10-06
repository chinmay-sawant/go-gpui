package store

import (
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

func TestPageKeysetAndFreeze(t *testing.T) {
	st := newStore(t)

	setup, err := st.EnsureDummy(bg(), DummyOptions{Count: 500})
	if err != nil {
		t.Fatal(err)
	}

	sess := setup.Session.ID
	q := Query{Session: &sess}

	p1, err := st.Page(bg(), q, PageOptions{Limit: 200})
	if err != nil {
		t.Fatal(err)
	}

	if len(p1.Entries) != 200 || !p1.More {
		t.Fatalf("page 1 = %d entries, more=%v", len(p1.Entries), p1.More)
	}

	if p1.HighWater != p1.Newest {
		t.Fatalf("high water %d, newest %d", p1.HighWater, p1.Newest)
	}

	for i := 1; i < len(p1.Entries); i++ {
		if p1.Entries[i].ID >= p1.Entries[i-1].ID {
			t.Fatal("page is not ordered by descending ID")
		}
	}

	q2 := q
	q2.MaxID = p1.HighWater
	q2.Cursor = p1.Entries[len(p1.Entries)-1].ID

	p2, err := st.Page(bg(), q2, PageOptions{Limit: 200})
	if err != nil {
		t.Fatal(err)
	}

	if len(p2.Entries) != 200 {
		t.Fatalf("page 2 = %d entries", len(p2.Entries))
	}

	if p2.Entries[0].ID >= q2.Cursor {
		t.Fatal("page 2 overlaps page 1")
	}

	// Freeze: a new entry must not appear in the pinned view.
	src := setup.Sources[0]

	_, err = st.Commit(bg(), src.ID, CommitMeta{
		Generation: 1, Position: src.Position + 1, State: entry.StateLive,
	}, []entry.Entry{{
		Session: sess, Source: src.ID, Position: src.Position, Generation: 1,
		Message: "fresh", Severity: entry.Info, Received: time.Now(),
	}})
	if err != nil {
		t.Fatal(err)
	}

	live, err := st.Page(bg(), q, PageOptions{Limit: 1})
	if err != nil {
		t.Fatal(err)
	}

	if live.HighWater <= p1.HighWater {
		t.Fatalf("live high water %d did not move", live.HighWater)
	}

	frozen, err := st.Page(bg(), Query{Session: &sess, MaxID: p1.HighWater}, PageOptions{Limit: 1})
	if err != nil {
		t.Fatal(err)
	}

	if frozen.Entries[0].ID != p1.Entries[0].ID {
		t.Fatalf("frozen view moved to %d", frozen.Entries[0].ID)
	}
}
