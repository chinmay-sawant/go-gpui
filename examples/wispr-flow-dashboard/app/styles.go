package app

import (
	"strings"

	"github.com/chinmay-sawant/go-gpui/examples/wispr-flow-dashboard/dictation"
	"github.com/chinmay-sawant/go-gpui/examples/wispr-flow-dashboard/dictionary"
	"github.com/chinmay-sawant/go-gpui/examples/wispr-flow-dashboard/freemonth"
	"github.com/chinmay-sawant/go-gpui/examples/wispr-flow-dashboard/help"
	"github.com/chinmay-sawant/go-gpui/examples/wispr-flow-dashboard/insights"
	"github.com/chinmay-sawant/go-gpui/examples/wispr-flow-dashboard/invite"
	"github.com/chinmay-sawant/go-gpui/examples/wispr-flow-dashboard/notetaker"
	"github.com/chinmay-sawant/go-gpui/examples/wispr-flow-dashboard/scratchpad"
	"github.com/chinmay-sawant/go-gpui/examples/wispr-flow-dashboard/settings"
	"github.com/chinmay-sawant/go-gpui/examples/wispr-flow-dashboard/snippets"
	"github.com/chinmay-sawant/go-gpui/examples/wispr-flow-dashboard/style"
	"github.com/chinmay-sawant/go-gpui/examples/wispr-flow-dashboard/transforms"
	"github.com/chinmay-sawant/go-gpui/examples/wispr-flow-dashboard/voice"
)

// appStyles are the shared stylesheets, in order.
var appStyles = []string{"base.css", "ui.css", "shell.css", "sidebar.css", "getapp.css"}

// styleSheets concatenates the shared styles and every screen stylesheet.
func styleSheets() string {
	var b strings.Builder

	for _, name := range appStyles {
		b.WriteString(file("components/" + name))
	}

	for _, sheet := range []string{
		insights.CSS(),
		dictation.CSS(),
		notetaker.CSS(),
		dictionary.CSS(),
		snippets.CSS(),
		style.CSS(),
		transforms.CSS(),
		scratchpad.CSS(),
		settings.CSS(),
		invite.CSS(),
		freemonth.CSS(),
		help.CSS(),
		voice.CSS(),
	} {
		b.WriteString(sheet)
	}

	return b.String()
}
