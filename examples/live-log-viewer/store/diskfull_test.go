package store

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

func TestDiskFullCommitFails(t *testing.T) {
	st := newStore(t)

	setup, err := st.EnsureDummy(bg(), DummyOptions{Count: 3})
	if err != nil {
		t.Fatal(err)
	}

	src := setup.Sources[0]

	var pages int64

	if err := st.db.QueryRowContext(bg(), "PRAGMA page_count").Scan(&pages); err != nil {
		t.Fatal(err)
	}

	if _, err := st.db.ExecContext(bg(),
		fmt.Sprintf("PRAGMA max_page_count = %d", pages+1)); err != nil {
		t.Fatal(err)
	}

	big := strings.Repeat("x", 1<<20)

	_, err = st.Commit(bg(), src.ID, CommitMeta{
		Generation: 1, Position: src.Position + 1, State: entry.StateLive,
	}, []entry.Entry{{
		Session: setup.Session.ID, Source: src.ID, Position: src.Position,
		Generation: 1, Message: big, Bytes: len(big), Received: time.Now(),
	}})
	if err == nil {
		t.Fatal("commit succeeded on a full database")
	}

	got, err := st.Source(bg(), src.ID)
	if err != nil {
		t.Fatal(err)
	}

	if got.Position != src.Position {
		t.Fatalf("checkpoint moved to %d after a failed commit", got.Position)
	}

	if _, err := st.Count(bg(), Query{Source: &src.ID}); err != nil {
		t.Fatal(err)
	}
}
