package store

import (
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

func TestCommitPartialReplacement(t *testing.T) {
	st := newStore(t)

	setup, err := st.EnsureDummy(bg(), DummyOptions{Count: 3})
	if err != nil {
		t.Fatal(err)
	}

	src := setup.Sources[0]

	part := entry.Entry{
		Session: setup.Session.ID, Source: src.ID, Position: 100,
		Generation: 1, Severity: entry.Info, Message: "INFO par",
		Bytes: 8, Partial: true, Received: time.Now(),
	}

	if _, err := st.Commit(bg(), src.ID, CommitMeta{
		Generation: 1, Position: 108, State: entry.StateLive,
	}, []entry.Entry{part}); err != nil {
		t.Fatal(err)
	}

	full := part
	full.Partial = false
	full.Bytes = 9
	full.Message = "INFO part"

	if _, err := st.Commit(bg(), src.ID, CommitMeta{
		Generation: 1, Position: 109, State: entry.StateLive,
	}, []entry.Entry{full}); err != nil {
		t.Fatal(err)
	}

	pos, err := st.LastPosition(bg(), src.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !pos.OK || pos.Partial || pos.Bytes != 9 {
		t.Fatalf("last position = %+v", pos)
	}

	// Replaying the old partial must not undo the completed row.
	if _, err := st.Commit(bg(), src.ID, CommitMeta{
		Generation: 1, Position: 108, State: entry.StateLive,
	}, []entry.Entry{part}); err != nil {
		t.Fatal(err)
	}

	pos2, err := st.LastPosition(bg(), src.ID)
	if err != nil {
		t.Fatal(err)
	}

	if pos2.Partial || pos2.Bytes != 9 {
		t.Fatalf("partial overwrote complete row: %+v", pos2)
	}
}
