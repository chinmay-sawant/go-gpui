// Package activity is the Activity menu of the Teams example: the feed, its
// filters, and the detail pane for one activity.
package activity

import (
	"context"
	"strings"

	"github.com/chinmay-sawant/ownframe"
)

// Data is the activity state the shell prints.
type Data struct {
	Filter string // "all" | "mentions" | "replies"
	Active string
	Items  []Item
}

// Item is one row of the activity feed.
type Item struct {
	ID       string
	Actor    string
	Initials string
	Color    string // avatar class suffix: "violet", "blue", ...
	Kind     string // "mention" | "reply" | "reaction" | "follow"
	Where    string // e.g. "Avengers Tower › Mission Briefings"
	Text     string // e.g. "mentioned you in Avengers Tower › Mission Briefings"
	Preview  string // the message body
	Time     string // e.g. "9:42 AM"
	Unread   bool
	Replies  int
	Replied  bool
}

// ActiveItem returns the item the detail pane shows, or nil.
func (d Data) ActiveItem() *Item {
	for i := range d.Items {
		if d.Items[i].ID == d.Active {
			return &d.Items[i]
		}
	}

	return nil
}

// Matches reports whether the item appears under a feed filter.
func (i Item) Matches(filter string) bool {
	switch filter {
	case "mentions":
		return i.Kind == "mention"
	case "replies":
		return i.Kind == "reply"
	default:
		return true
	}
}

// Handle applies one click action; it reports whether the action was ours.
func Handle(_ context.Context, _ *ownframe.Page, d *Data, action string) bool {
	switch {
	case strings.HasPrefix(action, "activity-filter-"):
		d.Filter = strings.TrimPrefix(action, "activity-filter-")
		d.Active = ""
	case strings.HasPrefix(action, "activity-open-"):
		d.open(strings.TrimPrefix(action, "activity-open-"))
	case action == "activity-reply":
		d.reply()
	case action == "activity-mark":
		d.markAllRead()
	default:
		return false
	}

	return true
}
