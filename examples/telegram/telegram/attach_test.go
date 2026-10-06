package telegram_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/go-gpui/examples/telegram/telegram"
)

func hasBox(app *telegram.App, id string) bool {
	for _, b := range app.Boxes() {
		if b.ID == id {
			return true
		}
	}

	return false
}

func TestAttachDrawerOpensAndCloses(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "", "open-anna")
	click(t, ctx, app, "attach", "")

	if !app.View().AttachOpen {
		t.Fatal("drawer did not open")
	}

	boxByID(t, app, "attach-camera")
	boxByID(t, app, "attach-gallery")

	click(t, ctx, app, "attach", "")

	if app.View().AttachOpen {
		t.Fatal("drawer did not close")
	}

	if hasBox(app, "attach-camera") {
		t.Fatal("drawer still drawn")
	}
}

func TestAttachRowsAskForTheRightKind(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "", "open-anna")
	click(t, ctx, app, "attach", "")
	click(t, ctx, app, "attach-camera", "")

	if kind := app.TakeAttach(); kind != "camera" {
		t.Fatalf("kind = %q", kind)
	}

	click(t, ctx, app, "attach", "")
	click(t, ctx, app, "attach-gallery", "")

	if kind := app.TakeAttach(); kind != "gallery" {
		t.Fatalf("kind = %q", kind)
	}
}

func TestAttachDrawerInEveryChat(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "", "open-max")
	click(t, ctx, app, "attach", "")

	if !app.View().AttachOpen {
		t.Fatal("drawer missing in Max's chat")
	}

	boxByID(t, app, "attach-camera")
}
