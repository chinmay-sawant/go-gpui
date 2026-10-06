package history_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/chinmay-sawant/ownframe"
)

func TestRoutesAndHistory(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	if len(app.PNG()) == 0 {
		t.Fatal("png is empty after the first redraw")
	}

	findBox(t, app, "back")
	findBox(t, app, "forward")

	clickAction(t, ctx, app, "red")

	if got := app.HTML(); !strings.Contains(got, "RED") {
		t.Fatalf("html after the red route = %q", got)
	}

	// A routed page replaces the DOM but keeps the toolbar.
	findBox(t, app, "back")

	if err := app.Back(ctx); err != nil {
		t.Fatal(err)
	}

	if got := app.HTML(); strings.Contains(got, "RED") {
		t.Fatalf("html after back = %q, want the first page", got)
	}

	if err := app.Back(ctx); !errors.Is(err, ownframe.ErrNoHistory) {
		t.Fatalf("second back = %v, want ErrNoHistory", err)
	}

	if err := app.Forward(ctx); err != nil {
		t.Fatal(err)
	}

	if got := app.HTML(); !strings.Contains(got, "RED") {
		t.Fatalf("html after forward = %q, want the red page", got)
	}

	// The routed pages carry the toolbar, so Back and Forward work as clicks.
	clickBox(t, ctx, app, "back")
	if got := app.View().Status; got != "html has initial" {
		t.Fatalf("status = %q, want the initial page", got)
	}

	clickBox(t, ctx, app, "back")
	if got := app.View().Status; got != "no history" {
		t.Fatalf("status = %q, want no history", got)
	}

	const native = `<html><body><p id="page">BLUE NATIVE</p></body></html>`
	if err := app.Load(ctx, native); err != nil {
		t.Fatal(err)
	}

	if got := app.HTML(); !strings.Contains(got, "BLUE NATIVE") {
		t.Fatalf("html after a native load = %q", got)
	}
}
