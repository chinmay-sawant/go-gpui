package app

import (
	"github.com/chinmay-sawant/ownframe/examples/wispr-flow-dashboard/dictation"
	"github.com/chinmay-sawant/ownframe/examples/wispr-flow-dashboard/dictionary"
	"github.com/chinmay-sawant/ownframe/examples/wispr-flow-dashboard/freemonth"
	"github.com/chinmay-sawant/ownframe/examples/wispr-flow-dashboard/help"
	"github.com/chinmay-sawant/ownframe/examples/wispr-flow-dashboard/insights"
	"github.com/chinmay-sawant/ownframe/examples/wispr-flow-dashboard/invite"
	"github.com/chinmay-sawant/ownframe/examples/wispr-flow-dashboard/notetaker"
	"github.com/chinmay-sawant/ownframe/examples/wispr-flow-dashboard/scratchpad"
	"github.com/chinmay-sawant/ownframe/examples/wispr-flow-dashboard/settings"
	"github.com/chinmay-sawant/ownframe/examples/wispr-flow-dashboard/snippets"
	"github.com/chinmay-sawant/ownframe/examples/wispr-flow-dashboard/style"
	"github.com/chinmay-sawant/ownframe/examples/wispr-flow-dashboard/transforms"
	"github.com/chinmay-sawant/ownframe/examples/wispr-flow-dashboard/voice"
)

// DefaultView returns the sample data the whole app draws.
func DefaultView() View {
	return View{
		ActivePage: "dictation",
		Nav:        navItems(),
		FootNav:    footNavItems(),
		Data:       insights.Default(),
		Dictation:  dictation.Default(),
		Notetaker:  notetaker.Default(),
		Dictionary: dictionary.Default(),
		Snippets:   snippets.Default(),
		Style:      style.Default(),
		Transforms: transforms.Default(),
		Scratchpad: scratchpad.Default(),
		Settings:   settings.Default(),
		Invite:     invite.Default(),
		FreeMonth:  freemonth.Default(),
		Help:       help.Default(),
		Voice:      voice.Default(),
	}
}
