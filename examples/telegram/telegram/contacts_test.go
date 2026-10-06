package telegram_test

import (
	"context"
	"testing"
)

func TestContactWithoutAChatGetsOne(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "tab-contacts", "")
	click(t, ctx, app, "", "contact-nina")

	if got := app.View().Active; got != "nina" {
		t.Fatalf("active = %q", got)
	}

	if got := len(app.View().Thread); got != 0 {
		t.Fatalf("thread = %d", got)
	}

	click(t, ctx, app, "chat-back", "")

	if got := app.View().Tab; got != "contacts" {
		t.Fatalf("tab = %q", got)
	}
}

func TestContactChatIsReused(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "tab-contacts", "")
	click(t, ctx, app, "", "contact-nina")
	click(t, ctx, app, "chat-back", "")
	click(t, ctx, app, "", "contact-nina")

	if got := app.View().Active; got != "nina" {
		t.Fatalf("active = %q", got)
	}

	n := 0

	for _, c := range app.View().Chats {
		if c.ID == "nina" {
			n++
		}
	}

	if n != 1 {
		t.Fatalf("nina chats = %d", n)
	}
}
