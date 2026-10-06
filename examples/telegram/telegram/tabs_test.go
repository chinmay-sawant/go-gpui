package telegram_test

import (
	"context"
	"testing"
)

func TestTabsSwitchScreens(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "tab-contacts", "")

	if got := app.View().Tab; got != "contacts" {
		t.Fatalf("tab = %q", got)
	}

	boxByID(t, app, "contact-nina")

	click(t, ctx, app, "tab-settings", "")

	if got := app.View().Tab; got != "settings" {
		t.Fatalf("tab = %q", got)
	}

	boxByID(t, app, "dark")

	click(t, ctx, app, "tab-chats", "")

	if got := app.View().Tab; got != "chats" {
		t.Fatalf("tab = %q", got)
	}

	boxByID(t, app, "chat-anna")
}

func TestDarkToggleSwitchesTheTheme(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "tab-settings", "")
	click(t, ctx, app, "dark", "")

	if !app.View().Dark {
		t.Fatal("dark is off after the toggle")
	}

	click(t, ctx, app, "dark", "")

	if app.View().Dark {
		t.Fatal("dark is on after the second toggle")
	}
}
