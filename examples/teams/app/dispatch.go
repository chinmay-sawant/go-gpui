package app

import (
	"context"

	"github.com/chinmay-sawant/ownframe/examples/teams/activity"
	"github.com/chinmay-sawant/ownframe/examples/teams/calendar"
	"github.com/chinmay-sawant/ownframe/examples/teams/calls"
	"github.com/chinmay-sawant/ownframe/examples/teams/channels"
	"github.com/chinmay-sawant/ownframe/examples/teams/chat"
	"github.com/chinmay-sawant/ownframe/examples/teams/files"
)

// dispatch offers one click action to every menu. The first menu that
// recognizes the action handles it.
func (a *App) dispatch(ctx context.Context, action string) {
	switch {
	case activity.Handle(ctx, a.page, &a.view.Activity, action):
	case chat.Handle(ctx, a.page, &a.view.Chat, action):
	case channels.Handle(ctx, a.page, &a.view.Channels, action):
	case calendar.Handle(ctx, a.page, &a.view.Calendar, action):
	case calls.Handle(ctx, a.page, &a.view.Calls, action):
	case files.Handle(ctx, a.page, &a.view.Files, action):
	}
}
