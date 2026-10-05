package store

import (
	"reflect"
	"testing"

	"github.com/chinmay-sawant/go-gpui/examples/teams/calendar"
)

// checkRoundTrip compares every menu after the reopen and checks each change
// came back.
func checkRoundTrip(t *testing.T, got, want Data) {
	t.Helper()

	bad := func(name string, ok bool) {
		if !ok {
			t.Error(name)
		}
	}

	same := reflect.DeepEqual

	bad("activity", same(got.Activity, want.Activity))
	bad("chat", same(got.Chat, want.Chat))
	bad("channels", same(got.Channels, want.Channels))
	bad("calendar", sameEvents(got.Calendar.AllEvents(), want.Calendar.AllEvents()))
	bad("calls", got.Calls.Tab == "voicemail" && same(got.Calls, want.Calls))
	bad("files", same(got.Files, want.Files))
	bad("profile", got.Presence == "busy" && !got.Dark)

	unread := false
	for _, it := range got.Activity.Items {
		unread = unread || it.Unread
	}
	bad("activity-mark", !unread)

	sent := false
	for _, m := range got.Chat.Messages {
		sent = sent || m.Text == "Round-trip message"
	}
	bad("chat-send", sent)

	meeting := false
	for _, e := range got.Calendar.AllEvents() {
		meeting = meeting || e.Title == "Round-trip meeting"
	}
	bad("calendar-create", meeting)

	reply := false
	for _, p := range got.Channels.Posts {
		for _, r := range p.Replies {
			reply = reply || r.Text == "Round-trip reply"
		}
	}
	bad("channels-reply", reply)

	star := false
	for _, f := range got.Files.AllFiles() {
		if f.ID == "f2" {
			star = f.Starred
		}
	}
	bad("files-star", star)
}

// sameEvents reports whether two calendar event lists hold the same events.
// Load orders events by week and position, so the lists are compared as a set
// keyed by event id.
func sameEvents(got, want []calendar.Event) bool {
	if len(got) != len(want) {
		return false
	}

	byID := map[string]calendar.Event{}

	for _, e := range got {
		byID[e.ID] = e
	}

	for _, e := range want {
		if byID[e.ID] != e {
			return false
		}
	}

	return true
}
