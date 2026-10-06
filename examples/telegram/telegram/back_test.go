package telegram_test

import (
	"context"
	"testing"
)

func TestBackLeavesTheOpenChat(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "", "open-anna")

	if !app.RequestBack() {
		t.Fatal("back refused in a chat")
	}

	if err := app.Tick(ctx); err != nil {
		t.Fatal(err)
	}

	if got := app.View().Active; got != "" {
		t.Fatalf("active = %q", got)
	}
}

func TestRequestBackRefusesOnAList(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	if app.RequestBack() {
		t.Fatal("back accepted while on the list")
	}

	if err := app.Tick(ctx); err != nil {
		t.Fatal(err)
	}

	if got := app.View().Active; got != "" {
		t.Fatalf("active = %q", got)
	}
}
