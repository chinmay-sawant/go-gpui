package app

import (
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

// pageEntry is one page branch in the shell.
type pageEntry struct {
	Name string
	HTML func() string
}

// pages are the page branches in sidebar order.
var pages = []pageEntry{
	{"dictation", dictation.HTML},
	{"notetaker", notetaker.HTML},
	{"insights", insights.HTML},
	{"dictionary", dictionary.HTML},
	{"snippets", snippets.HTML},
	{"style", style.HTML},
	{"transforms", transforms.HTML},
	{"scratchpad", scratchpad.HTML},
	{"invite", invite.HTML},
	{"free-month", freemonth.HTML},
	{"settings", settings.HTML},
	{"help", help.HTML},
	{"voice", voice.HTML},
}
