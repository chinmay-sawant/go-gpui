package store

import (
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

func TestCheckpointCadence(t *testing.T) {
	st := newStoreOpts(t, Options{CheckpointEvery: 1})

	setup, err := st.EnsureDummy(bg(), DummyOptions{Count: 3})
	if err != nil {
		t.Fatal(err)
	}

	src := setup.Sources[0]

	for i := 0; i < 5; i++ {
		if _, err := st.Commit(bg(), src.ID, CommitMeta{
			Generation: 1, Position: src.Position + int64(i+1),
			State: entry.StateLive,
		}, []entry.Entry{{
			Session: setup.Session.ID, Source: src.ID,
			Position: src.Position + int64(i), Generation: 1,
			Message: "x", Bytes: 1, Received: time.Now(),
		}}); err != nil {
			t.Fatal(err)
		}
	}

	if err := st.Checkpoint(bg()); err != nil {
		t.Fatal(err)
	}

	got, err := st.Source(bg(), src.ID)
	if err != nil {
		t.Fatal(err)
	}

	if got.Position != src.Position+5 {
		t.Fatalf("checkpoint = %d, want %d", got.Position, src.Position+5)
	}

	stats, err := st.Stats(bg())
	if err != nil {
		t.Fatal(err)
	}

	if stats.Journal != "wal" {
		t.Fatalf("journal = %q", stats.Journal)
	}
}
