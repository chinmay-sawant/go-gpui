package ui

import (
	"context"
	"testing"
	"time"
)

func TestAppFollowsTail(t *testing.T) {
	ff := &fakeFeed{}
	ff.seed(250)
	a := newApp(t, ff, Options{})

	pumpUntil(t, a, func() bool { return a.pager.Len() == PageLimit }, "initial page")

	if got := a.lastID(); got != 250 {
		t.Fatalf("newest loaded = %d", got)
	}

	if !a.follow.Follow || a.pager.HWM != 0 {
		t.Fatalf("follow = %v hwm = %d", a.follow.Follow, a.pager.HWM)
	}

	ff.add(Entry{ID: 251, Text: "new"}, Entry{ID: 252, Text: "new"})
	pumpUntil(t, a, func() bool { return a.lastID() == 252 }, "tail append")

	if a.follow.Unread != 0 {
		t.Fatalf("unread while live = %d", a.follow.Unread)
	}

	a.togglePause()
	ff.add(Entry{ID: 253, Text: "new"}, Entry{ID: 254, Text: "new"})
	pumpUntil(t, a, func() bool { return a.follow.Unread == 2 }, "paused unread")

	if got := a.lastID(); got != 252 {
		t.Fatalf("paused page moved to %d", got)
	}

	a.togglePause()
	for i := 0; i < 5; i++ {
		_ = a.Tick(context.Background())
		time.Sleep(10 * time.Millisecond)
	}

	if got := a.lastID(); got != 252 {
		t.Fatalf("resume jumped to %d", got)
	}

	a.jumpNewest()
	pumpUntil(t, a, func() bool { return a.lastID() == 254 }, "jump to newest")

	if a.follow.Unread != 0 {
		t.Fatalf("unread after jump = %d", a.follow.Unread)
	}
}
