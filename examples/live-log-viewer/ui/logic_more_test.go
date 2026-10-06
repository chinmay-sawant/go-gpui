package ui

import (
	"testing"
	"time"
)

func TestFilterDebounce(t *testing.T) {
	now := time.Unix(1000, 0)
	f := filterState{}

	f.Edit("abc", now, 250*time.Millisecond)
	if _, _, ok := f.Ready(now.Add(100 * time.Millisecond)); ok {
		t.Fatal("fired before the deadline")
	}

	text, gen, ok := f.Ready(now.Add(300 * time.Millisecond))
	if !ok || text != "abc" || gen != 1 {
		t.Fatalf("due edit = %q %d %v", text, gen, ok)
	}

	if _, _, ok := f.Ready(now.Add(400 * time.Millisecond)); ok {
		t.Fatal("fired twice")
	}
}

func TestFilterGenerations(t *testing.T) {
	now := time.Unix(1000, 0)
	f := filterState{}

	f.Edit("a", now, time.Millisecond)
	_, oldGen, _ := f.Ready(now.Add(time.Second))

	f.Apply("b")
	if !f.Stale(oldGen) {
		t.Fatal("old generation accepted")
	}

	if f.Stale(f.Gen) {
		t.Fatal("current generation rejected")
	}

	f.Edit("b", now, time.Millisecond)
	if _, _, ok := f.Ready(now.Add(time.Second)); ok {
		t.Fatal("unchanged text dispatched")
	}
}

func TestFollowPauseAndUnread(t *testing.T) {
	f := followState{Follow: true}
	f.Pause()

	if f.Live() {
		t.Fatal("paused state is live")
	}

	f.Note([]Entry{{ID: 3}, {ID: 4}}, 2)
	f.Note(nil, 1)

	if f.Unread != 3 {
		t.Fatalf("unread = %d", f.Unread)
	}

	if f.LastSeen != 4 {
		t.Fatalf("last seen = %d", f.LastSeen)
	}

	f.Resume()
	if f.Live() {
		t.Fatal("resume restarted following")
	}

	f.SetFollow(true)
	f.Note([]Entry{{ID: 5}}, 1)

	if f.Unread != 0 {
		t.Fatalf("live unread = %d", f.Unread)
	}
}

func TestFollowSeen(t *testing.T) {
	f := followState{Follow: true}
	f.Seen(10)
	f.Note(nil, 0)

	if f.LastSeen != 10 {
		t.Fatalf("seen = %d", f.LastSeen)
	}

	f.ClearUnread()
}
