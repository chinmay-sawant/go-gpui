package app

import (
	"github.com/chinmay-sawant/ownframe/examples/teams/activity"
	"github.com/chinmay-sawant/ownframe/examples/teams/calendar"
	"github.com/chinmay-sawant/ownframe/examples/teams/calls"
	"github.com/chinmay-sawant/ownframe/examples/teams/channels"
	"github.com/chinmay-sawant/ownframe/examples/teams/chat"
	"github.com/chinmay-sawant/ownframe/examples/teams/files"
)

// pageEntry is one menu branch in the shell.
type pageEntry struct {
	Name string
	HTML func() string
}

// pages are the menu branches in rail order.
var pages = []pageEntry{
	{"activity", activity.HTML},
	{"chat", chat.HTML},
	{"channels", channels.HTML},
	{"calendar", calendar.HTML},
	{"calls", calls.HTML},
	{"files", files.HTML},
}
