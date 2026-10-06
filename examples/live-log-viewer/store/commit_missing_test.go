package store

import (
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

func TestCommitMissingSourceFails(t *testing.T) {
	st := newStore(t)

	_, err := st.Commit(bg(), entry.SourceID(999), CommitMeta{
		Generation: 1, Position: 0,
	}, []entry.Entry{{
		Session: 1, Source: 999, Position: 0, Message: "x", Received: time.Now(),
	}})
	if err == nil {
		t.Fatal("commit for a missing source returned nil")
	}
}

func TestCommitDeletedSourceFails(t *testing.T) {
	st := newStore(t)

	setup, err := st.EnsureDummy(bg(), DummyOptions{Count: 3})
	if err != nil {
		t.Fatal(err)
	}

	src := setup.Sources[0]

	if err := st.DeleteSource(bg(), src.ID); err != nil {
		t.Fatal(err)
	}

	_, err = st.Commit(bg(), src.ID, CommitMeta{Generation: 1, Position: 0},
		[]entry.Entry{{
			Session: setup.Session.ID, Source: src.ID, Position: 0,
			Message: "x", Received: time.Now(),
		}})
	if err == nil {
		t.Fatal("commit for a deleted source returned nil")
	}
}
