package telegram_test

import (
	"context"
	"testing"
)

func TestSearchFiltersTheChatList(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "search", "")

	if err := app.Type(ctx, "dev"); err != nil {
		t.Fatal(err)
	}

	view := app.View()
	if got := len(view.Chats); got != 1 {
		t.Fatalf("chats = %d", got)
	}

	if got := view.Chats[0].ID; got != "devs" {
		t.Fatalf("chat = %q", got)
	}
}

func TestSearchCanMatchNothing(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "search", "")

	if err := app.Type(ctx, "zz"); err != nil {
		t.Fatal(err)
	}

	if got := len(app.View().Chats); got != 0 {
		t.Fatalf("chats = %d", got)
	}

	if got := app.View().Unread; got != 20 {
		t.Fatalf("unread = %d", got)
	}
}
