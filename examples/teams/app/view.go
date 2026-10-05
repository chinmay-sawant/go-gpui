package app

import (
	"github.com/chinmay-sawant/go-gpui/examples/teams/activity"
	"github.com/chinmay-sawant/go-gpui/examples/teams/calendar"
	"github.com/chinmay-sawant/go-gpui/examples/teams/calls"
	"github.com/chinmay-sawant/go-gpui/examples/teams/channels"
	"github.com/chinmay-sawant/go-gpui/examples/teams/chat"
	"github.com/chinmay-sawant/go-gpui/examples/teams/files"
)

// View is the data the app template prints.
type View struct {
	Section  string
	Note     string
	Dark     bool
	Presence string
	Flyout   string
	Rail     []RailItem
	Apps     []AppTile

	Activity activity.Data
	Chat     chat.Data
	Channels channels.Data
	Calendar calendar.Data
	Calls    calls.Data
	Files    files.Data
}

// DefaultView returns the sample state every menu starts from.
func DefaultView() View {
	return View{
		Section:  "chat",
		Dark:     true,
		Presence: "online",
		Rail:     railItems(),
		Apps:     appTiles(),
		Activity: activity.Default(),
		Chat:     chat.Default(),
		Channels: channels.Default(),
		Calendar: calendar.Default(),
		Calls:    calls.Default(),
		Files:    files.Default(),
	}
}
