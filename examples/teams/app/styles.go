package app

import (
	"strings"

	"github.com/chinmay-sawant/ownframe/examples/teams/activity"
	"github.com/chinmay-sawant/ownframe/examples/teams/calendar"
	"github.com/chinmay-sawant/ownframe/examples/teams/calls"
	"github.com/chinmay-sawant/ownframe/examples/teams/channels"
	"github.com/chinmay-sawant/ownframe/examples/teams/chat"
	"github.com/chinmay-sawant/ownframe/examples/teams/files"
)

// appStyles are the shared stylesheets, in order.
var appStyles = []string{"base.css", "ui.css", "rail.css", "flyout.css"}

// styleSheets concatenates the shared styles and every menu stylesheet.
func styleSheets() string {
	var b strings.Builder

	for _, name := range appStyles {
		b.WriteString(file("components/" + name))
	}

	for _, sheet := range []string{
		activity.CSS(),
		chat.CSS(),
		channels.CSS(),
		calendar.CSS(),
		calls.CSS(),
		files.CSS(),
	} {
		b.WriteString(sheet)
	}

	return b.String()
}
