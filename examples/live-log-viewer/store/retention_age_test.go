package store

import (
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

func TestPruneAge(t *testing.T) {
	st := newStoreOpts(t, Options{Retention: Retention{MaxAge: 24 * time.Hour}})

	sess, err := st.EnsureSession(bg(), "age", "burst")
	if err != nil {
		t.Fatal(err)
	}

	src, err := st.AddSource(bg(), SourceSpec{
		Session: sess.ID, Kind: entry.KindBurst, Path: "age", Total: 20,
	})
	if err != nil {
		t.Fatal(err)
	}

	old := time.Now().Add(-48 * time.Hour)

	for i := 0; i < 8; i++ {
		received := old
		if i >= 5 {
			received = time.Now()
		}

		if _, err := st.Commit(bg(), src.ID, CommitMeta{
			Generation: 1, Position: int64(i + 1),
		}, []entry.Entry{{
			Session: sess.ID, Source: src.ID, Position: int64(i),
			Generation: 1, Message: "e", Received: received,
		}}); err != nil {
			t.Fatal(err)
		}
	}

	out, err := st.Prune(bg())
	if err != nil {
		t.Fatal(err)
	}

	if out.Entries != 5 {
		t.Fatalf("pruned %d entries, want 5", out.Entries)
	}

	s, err := st.Stats(bg())
	if err != nil {
		t.Fatal(err)
	}

	if s.Entries != 3 {
		t.Fatalf("entries = %d, want 3", s.Entries)
	}
}
