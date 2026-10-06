package store

import (
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

func TestPageExpiredCursor(t *testing.T) {
	st := newStoreOpts(t, Options{Retention: Retention{MaxRows: 2}})

	sess, err := st.EnsureSession(bg(), "expire", "burst")
	if err != nil {
		t.Fatal(err)
	}

	src, err := st.AddSource(bg(), SourceSpec{
		Session: sess.ID, Kind: entry.KindBurst, Path: "expire", Total: 10,
	})
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 5; i++ {
		if _, err := st.Commit(bg(), src.ID, CommitMeta{
			Generation: 1, Position: int64(i + 1),
		}, []entry.Entry{{
			Session: sess.ID, Source: src.ID, Position: int64(i),
			Generation: 1, Message: "e", Received: time.Now(),
		}}); err != nil {
			t.Fatal(err)
		}
	}

	if _, err := st.Prune(bg()); err != nil {
		t.Fatal(err)
	}

	p, err := st.Page(bg(), Query{Session: &sess.ID, Cursor: 1}, PageOptions{Limit: 10})
	if err != nil {
		t.Fatal(err)
	}

	if len(p.Entries) != 0 || !p.Expired {
		t.Fatalf("page = %d entries, expired=%v", len(p.Entries), p.Expired)
	}
}
