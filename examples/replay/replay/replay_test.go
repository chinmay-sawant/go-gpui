package replay_test

import (
	"context"
	"strings"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/replay/replay"
)

func TestReplayAndFallbackRoutes(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	if !app.Replayable() {
		t.Fatal("replay page did not keep a display list")
	}

	if data := app.PNG(); len(data) == 0 {
		t.Fatal("no PNG")
	}

	click(t, ctx, app, "action-fallback")

	if app.Replayable() {
		t.Fatal("fallback page kept a display list")
	}

	if got := app.View().Status; !strings.Contains(got, "bitmap") {
		t.Fatalf("status = %q, want the fallback mode", got)
	}

	click(t, ctx, app, "action-replay")

	if !app.Replayable() {
		t.Fatal("replay route did not return to the display list")
	}
}

func newApp(t *testing.T, ctx context.Context) *replay.App {
	t.Helper()

	app, err := replay.New()
	if err != nil {
		t.Fatal(err)
	}

	app.SetSize(replay.DefaultWidth, replay.DefaultHeight)

	if err := app.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	return app
}

func click(t *testing.T, ctx context.Context, app *replay.App, id string) {
	t.Helper()

	var x, y float64
	found := false

	for _, b := range app.Boxes() {
		if b.ID != id || b.W <= 0 || b.H <= 0 {
			continue
		}

		x = b.X + b.W/2
		y = b.Y + b.H/2
		found = true
	}

	if !found {
		t.Fatalf("no box id=%q", id)
	}

	if err := app.Click(ctx, x, y); err != nil {
		t.Fatal(err)
	}
}
