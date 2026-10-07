package store

import (
	"strings"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

func TestPageFilters(t *testing.T) {
	st := newStore(t)

	setup, err := st.EnsureDummy(bg(), DummyOptions{Count: 600})
	if err != nil {
		t.Fatal(err)
	}

	sid := setup.Session.ID
	min := entry.Error

	p, err := st.Page(bg(), Query{Session: &sid, MinSeverity: &min}, PageOptions{Limit: 500})
	if err != nil {
		t.Fatal(err)
	}

	if len(p.Entries) == 0 {
		t.Fatal("severity filter returned nothing")
	}

	for _, e := range p.Entries {
		if e.Severity < entry.Error {
			t.Fatalf("severity %v below minimum", e.Severity)
		}
	}

	src := setup.Sources[1].ID

	p2, err := st.Page(bg(), Query{Session: &sid, Source: &src}, PageOptions{Limit: 500})
	if err != nil {
		t.Fatal(err)
	}

	if len(p2.Entries) == 0 {
		t.Fatal("source filter returned nothing")
	}

	for _, e := range p2.Entries {
		if e.Source != src {
			t.Fatalf("entry from source %d", e.Source)
		}
	}

	p3, err := st.Page(bg(), Query{Session: &sid, Text: "connection refused"}, PageOptions{Limit: 500})
	if err != nil {
		t.Fatal(err)
	}

	if len(p3.Entries) == 0 {
		t.Fatal("text filter returned nothing")
	}

	for _, e := range p3.Entries {
		if !strings.Contains(e.Message, "connection refused") {
			t.Fatalf("message %q does not match", e.Message)
		}
	}

	last := p2.Entries[len(p2.Entries)-1].ID

	n, err := st.Count(bg(), Query{Session: &sid, Source: &src, Cursor: last})
	if err != nil {
		t.Fatal(err)
	}

	if n != int64(len(p2.Entries)-1) {
		t.Fatalf("unread count = %d, want %d", n, len(p2.Entries)-1)
	}
}
