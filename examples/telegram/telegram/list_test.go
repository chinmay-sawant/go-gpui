package telegram_test

import (
	"context"
	"testing"
)

func TestChatListRenders(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	if got := len(app.View().Chats); got != 7 {
		t.Fatalf("chats = %d", got)
	}

	if got := app.View().Unread; got != 20 {
		t.Fatalf("unread = %d", got)
	}

	boxByID(t, app, "chat-anna")
	boxByID(t, app, "tab-chats")
}

func TestOpenChatShowsThread(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "", "open-anna")

	view := app.View()
	if view.Active != "anna" || view.Header != "Anna Petrova" {
		t.Fatalf("active = %q header = %q", view.Active, view.Header)
	}

	if got := len(view.Thread); got != 5 {
		t.Fatalf("thread = %d", got)
	}

	if got := view.Unread; got != 18 {
		t.Fatalf("unread after open = %d", got)
	}

	click(t, ctx, app, "chat-back", "")

	if got := app.View().Active; got != "" {
		t.Fatalf("active after back = %q", got)
	}
}

func TestGroupThreadKeepsAuthors(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "", "open-design")

	view := app.View()
	if !view.Group || view.Status != "5 members" {
		t.Fatalf("group = %v status = %q", view.Group, view.Status)
	}

	if got := view.Thread[0].Author; got != "Mira" {
		t.Fatalf("author = %q", got)
	}
}
