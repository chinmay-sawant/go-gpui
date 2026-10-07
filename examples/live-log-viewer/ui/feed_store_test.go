package ui

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/store"
)

func TestStoreFeedDummyPaging(t *testing.T) {
	a, _, newest := newStoreApp(t)

	pumpUntil(t, a, func() bool { return a.pager.Len() == PageLimit }, "newest page")

	if got := a.lastID(); got != newest {
		t.Fatalf("newest = %d want %d", got, newest)
	}

	if !a.pager.HasOlder || a.pager.HasNewer {
		t.Fatalf("flags = %v/%v", a.pager.HasOlder, a.pager.HasNewer)
	}

	a.loadPage(intentOlder)
	wantOlder := newest - 2*PageLimit + 1
	pumpUntil(t, a, func() bool { return a.pager.Entries[0].ID == wantOlder }, "older page")

	if !a.pager.HasNewer {
		t.Fatal("older page lost the newer boundary")
	}

	a.loadPage(intentNewer)
	pumpUntil(t, a, func() bool { return a.pager.Entries[0].ID == newest-PageLimit+1 }, "newer page")
}

func TestStoreFeedSeverityFilter(t *testing.T) {
	a, feed, _ := newStoreApp(t)

	pumpUntil(t, a, func() bool { return a.pager.Len() == PageLimit }, "first page")

	sev := entry.Error
	want, err := feed.st.Count(context.Background(), store.Query{MinSeverity: &sev})
	if err != nil {
		t.Fatal(err)
	}

	a.minSev = "error"
	a.applyFilterChange()
	pumpUntil(t, a, func() bool { return a.pager.Total == int(want) }, "filtered page")

	for _, e := range a.pager.Entries {
		if c := sevClass(e.Severity); c != "error" && c != "fatal" {
			t.Fatalf("severity %q in the error filter", e.Severity)
		}
	}
}
