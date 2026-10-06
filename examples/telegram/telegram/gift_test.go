package telegram_test

import (
	"context"
	"testing"
)

func TestGiftButtonOnlyInOneChat(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "", "open-anna")

	if !app.View().CanGift {
		t.Fatal("gift missing in Anna's chat")
	}

	boxByID(t, app, "gift")

	click(t, ctx, app, "chat-back", "")
	click(t, ctx, app, "", "open-max")

	if app.View().CanGift {
		t.Fatal("gift allowed in Max's chat")
	}
}

func TestGiftSendsABubble(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "", "open-anna")
	click(t, ctx, app, "gift", "")

	thread := app.View().Thread
	last := thread[len(thread)-1]

	if !last.Gift || !last.Own {
		t.Fatalf("last = %+v", last)
	}

	for _, c := range app.View().Chats {
		if c.ID == "anna" && c.Preview != "You: a gift" {
			t.Fatalf("preview = %q", c.Preview)
		}
	}
}
