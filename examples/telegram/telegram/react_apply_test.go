package telegram_test

import (
	"context"
	"testing"
)

func TestReactionSetsTheBadge(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "", "open-anna")
	longPress(t, ctx, app, "msg-anna-5")
	click(t, ctx, app, "react-love", "")

	if app.View().ReactID != "" {
		t.Fatalf("react id = %q", app.View().ReactID)
	}

	got := ""
	for _, m := range app.View().Thread {
		if m.ID == "anna-5" {
			got = m.Reaction
		}
	}

	if got != "❤️" {
		t.Fatalf("reaction = %q", got)
	}

	if hasBox(app, "react-like") {
		t.Fatal("reaction row still drawn")
	}
}

func TestLongPressOutsideAMessageDoesNothing(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "", "open-anna")
	longPress(t, ctx, app, "attach")

	if app.View().ReactID != "" {
		t.Fatalf("react id = %q", app.View().ReactID)
	}

	if hasBox(app, "react-like") {
		t.Fatal("reaction row drawn")
	}
}
