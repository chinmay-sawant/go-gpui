package ui

import (
	"strings"
	"testing"
)

func TestStalePageDiscarded(t *testing.T) {
	a, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}

	defer func() { _ = a.Close() }()

	a.pager.Load(PageResult{Entries: seedEntries(3), Total: 3})
	oldGen := a.pageGen.Load()
	a.pageGen.Add(1)

	if a.applyPage(out{gen: oldGen, page: PageResult{Entries: seedEntries(9)}}) {
		t.Fatal("stale page applied")
	}

	if a.pager.Len() != 3 {
		t.Fatalf("page replaced by stale result: %d", a.pager.Len())
	}

	if !a.applyPage(out{gen: a.pageGen.Load(), page: PageResult{Entries: seedEntries(9)}}) {
		t.Fatal("current page discarded")
	}
}

func TestStaleDetailDiscarded(t *testing.T) {
	a, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}

	defer func() { _ = a.Close() }()

	a.detailID = 7
	oldGen := a.detailGen.Load()
	a.detailGen.Add(1)

	if a.applyDetail(out{gen: oldGen, id: 7, detail: Detail{Entry: Entry{ID: 7}}}) {
		t.Fatal("stale detail applied")
	}

	if a.detailOpen {
		t.Fatal("stale detail opened the pane")
	}

	a.detailID = 9
	if !a.applyDetail(out{gen: a.detailGen.Load(), id: 9, detail: Detail{Entry: Entry{ID: 9}}}) {
		t.Fatal("current detail discarded")
	}
}

func TestDeletedSourceFallsBack(t *testing.T) {
	ff := &fakeFeed{}
	ff.seed(5)
	a := newApp(t, ff, Options{})

	pumpUntil(t, a, func() bool { return len(a.sources) == 1 }, "sources")

	a.activeSource = "gone"
	a.refreshSources()
	pumpUntil(t, a, func() bool { return a.activeSource == "" }, "fallback")

	if !strings.Contains(a.note, "gone") {
		t.Fatalf("note = %q", a.note)
	}
}
