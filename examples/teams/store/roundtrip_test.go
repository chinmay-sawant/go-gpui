package store

import (
	"testing"

	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/examples/teams/activity"
	"github.com/chinmay-sawant/go-gpui/examples/teams/calendar"
	"github.com/chinmay-sawant/go-gpui/examples/teams/calls"
	"github.com/chinmay-sawant/go-gpui/examples/teams/channels"
	"github.com/chinmay-sawant/go-gpui/examples/teams/chat"
	"github.com/chinmay-sawant/go-gpui/examples/teams/files"
)

// TestRoundTrip seeds from the SQL files under seed/, changes one thing per
// menu through the real handlers, and checks every change survives Save and
// Load.
func TestRoundTrip(t *testing.T) {
	must := func(err error) {
		if err != nil {
			t.Fatal(err)
		}
	}

	st, err := Open(Memory)
	must(err)

	defer st.Close()

	if err := st.Seed(); err != nil {
		t.Fatal(err)
	}

	d, ok, err := st.Load()
	if err != nil || !ok {
		t.Fatalf("Load: ok=%v err=%v", ok, err)
	}

	page, err := gpui.New(gpui.Config{Width: 400, Height: 300, HTML: `<input id="chat-compose"><input id="cal-title"><input id="channels-reply-av1">`})
	must(err)

	ctx := t.Context()
	must(page.Redraw(ctx))

	page.SetFormValue("chat-compose", "Round-trip message")
	chat.Handle(ctx, page, &d.Chat, "chat-send")
	page.SetFormValue("cal-title", "Round-trip meeting")
	calendar.Handle(ctx, page, &d.Calendar, "calendar-create")
	page.SetFormValue("channels-reply-av1", "Round-trip reply")
	channels.Handle(ctx, page, &d.Channels, "channels-reply-send-av1")
	activity.Handle(ctx, nil, &d.Activity, "activity-mark")
	files.Handle(ctx, nil, &d.Files, "files-star-f2")
	calls.Handle(ctx, nil, &d.Calls, "calls-tab-voicemail")
	d.Presence, d.Dark = "busy", false

	must(st.Save(d))

	got, ok, err := st.Load()
	if err != nil || !ok {
		t.Fatalf("Load: ok=%v err=%v", ok, err)
	}

	checkRoundTrip(t, got, d)
}
