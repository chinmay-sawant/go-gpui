package ui

import (
	"strings"
	"testing"
	"time"
)

func TestCloseStopsWorker(t *testing.T) {
	ff := &fakeFeed{}
	ff.seed(100)
	ff.delay = 5 * time.Second

	a := newApp(t, ff, Options{Poll: time.Hour})
	a.loadPage(intentNewest)
	time.Sleep(20 * time.Millisecond)

	start := time.Now()
	if err := a.Close(); err != nil {
		t.Fatalf("close: %v", err)
	}

	if took := time.Since(start); took > time.Second {
		t.Fatalf("close took %v", took)
	}
}

func TestSettingsApply(t *testing.T) {
	ff := &fakeFeed{}
	ff.seed(20)
	ff.settings = Settings{Dark: true, Follow: false, Severity: "error", Source: "s1"}
	a := newApp(t, ff, Options{})

	pumpUntil(t, a, func() bool { return a.settingsLoaded }, "settings")

	if !a.dark || a.follow.Follow || a.activeSource != "s1" || a.minSev != "error" {
		t.Fatalf("settings not applied: %+v", a.currentSettings())
	}
}

func TestRetentionNote(t *testing.T) {
	a, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}

	defer func() { _ = a.Close() }()

	a.pageGen.Add(1)
	a.applyPage(out{
		gen:  a.pageGen.Load(),
		page: PageResult{Entries: seedEntries(2), Total: 2, Expired: true},
	})

	if !strings.Contains(a.note, "retained history") {
		t.Fatalf("note = %q", a.note)
	}
}

func TestAppendTailTrimsPage(t *testing.T) {
	a, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}

	defer func() { _ = a.Close() }()

	a.pager.Load(PageResult{Entries: seedEntries(PageLimit), Total: PageLimit})

	added := a.appendTail([]Entry{{ID: 201}, {ID: 202}, {ID: 203}})
	if added != 3 {
		t.Fatalf("added = %d", added)
	}

	if a.pager.Len() != PageLimit {
		t.Fatalf("len = %d", a.pager.Len())
	}

	if a.pager.Entries[0].ID != 4 || a.pager.Entries[PageLimit-1].ID != 203 {
		t.Fatalf("window = %d..%d", a.pager.Entries[0].ID, a.pager.Entries[PageLimit-1].ID)
	}

	if !a.pager.HasOlder || a.pager.HasNewer {
		t.Fatal("page flags after trim")
	}
}
