package store

import (
	"testing"
)

func TestPruneRowsKeepsLimit(t *testing.T) {
	st := newStoreOpts(t, Options{Retention: Retention{MaxRows: 100}})

	if _, err := st.EnsureDummy(bg(), DummyOptions{Count: 300}); err != nil {
		t.Fatal(err)
	}

	out, err := st.Prune(bg())
	if err != nil {
		t.Fatal(err)
	}

	if out.Entries == 0 {
		t.Fatal("nothing pruned")
	}

	s, err := st.Stats(bg())
	if err != nil {
		t.Fatal(err)
	}

	if s.Entries != 100 {
		t.Fatalf("entries = %d, want 100", s.Entries)
	}
}

func TestPruneBytes(t *testing.T) {
	st := newStoreOpts(t, Options{Retention: Retention{MaxBytes: 200}})

	if _, err := st.EnsureDummy(bg(), DummyOptions{Count: 300}); err != nil {
		t.Fatal(err)
	}

	out, err := st.Prune(bg())
	if err != nil {
		t.Fatal(err)
	}

	if out.Entries == 0 {
		t.Fatal("nothing pruned")
	}

	s, err := st.Stats(bg())
	if err != nil {
		t.Fatal(err)
	}

	if s.Bytes > 200 {
		t.Fatalf("bytes = %d, want <= 200", s.Bytes)
	}
}
