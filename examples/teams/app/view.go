package app

import (
	"github.com/chinmay-sawant/ownframe/examples/teams/activity"
	"github.com/chinmay-sawant/ownframe/examples/teams/calendar"
	"github.com/chinmay-sawant/ownframe/examples/teams/calls"
	"github.com/chinmay-sawant/ownframe/examples/teams/channels"
	"github.com/chinmay-sawant/ownframe/examples/teams/chat"
	"github.com/chinmay-sawant/ownframe/examples/teams/files"
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

// DefaultView returns the view before the database loads: the rail, the
// demo tiles, and empty menu state.
func DefaultView() View {
	return View{
		Section:  "chat",
		Dark:     true,
		Presence: "online",
		Rail:     railItems(),
		Apps:     appTiles(),
	}
}
