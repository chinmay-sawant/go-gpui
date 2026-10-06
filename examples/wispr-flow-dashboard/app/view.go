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

// View is the data the app template prints.
type View struct {
	ActivePage       string
	SidebarCollapsed bool
	GetApp           bool
	Nav              []NavItem
	FootNav          []NavItem
	Note             string

	insights.Data

	Dictation  dictation.Data
	Notetaker  notetaker.Data
	Dictionary dictionary.Data
	Snippets   snippets.Data
	Style      style.Data
	Transforms transforms.Data
	Scratchpad scratchpad.Data
	Settings   settings.Data
	Invite     invite.Data
	FreeMonth  freemonth.Data
	Help       help.Data
	Voice      voice.Data
}
