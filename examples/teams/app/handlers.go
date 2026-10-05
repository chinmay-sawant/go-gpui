package app

import (
	"context"
	"strings"

	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/examples/teams/calls"
	"github.com/chinmay-sawant/go-gpui/examples/teams/channels"
	"github.com/chinmay-sawant/go-gpui/examples/teams/chat"
	"github.com/chinmay-sawant/go-gpui/examples/teams/files"
)

// onChange sends a search box keystroke to the menu that owns the list.
func (a *App) onChange(ctx context.Context, box gpui.Box) error {
	switch box.ID {
	case "chat-search":
		chat.Handle(ctx, a.page, &a.view.Chat, "chat-query")
	case "files-search":
		files.Handle(ctx, a.page, &a.view.Files, "files-query")
	case "calls-search":
		calls.Handle(ctx, a.page, &a.view.Calls, "calls-query")
	default:
		return nil
	}

	a.page.SetData(a.view)

	return nil
}

// onSubmit sends the focused composer on Enter or NumpadEnter.
func (a *App) onSubmit(ctx context.Context) error {
	id := a.page.FocusID()

	switch {
	case id == "chat-compose":
		chat.Handle(ctx, a.page, &a.view.Chat, "chat-send")
	case strings.HasPrefix(id, "channels-reply-"):
		channels.Handle(ctx, a.page, &a.view.Channels, "channels-reply-send-"+strings.TrimPrefix(id, "channels-reply-"))
	default:
		return nil
	}

	a.page.SetData(a.view)

	return nil
}
