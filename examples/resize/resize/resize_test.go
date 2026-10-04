package resize_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/examples/resize/resize"
)

func boxByID(boxes []gpui.Box, id string) (gpui.Box, bool) {
	for _, b := range boxes {
		if b.ID == id {
			return b, true
		}
	}

	return gpui.Box{}, false
}

func TestColumnsBarAndParagraphFollowTheSize(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app, err := resize.New()
	if err != nil {
		t.Fatal(err)
	}

	if err := app.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	wideLeft := boxOf(t, app, "left")
	wideRight := boxOf(t, app, "right")

	if wideLeft.Y != wideRight.Y || wideLeft.X >= wideRight.X {
		t.Fatalf("wide columns = left %.0f,%.0f right %.0f,%.0f",
			wideLeft.X, wideLeft.Y, wideRight.X, wideRight.Y)
	}

	if got := boxOf(t, app, "bar").W; got != resize.DefaultWidth {
		t.Fatalf("wide bar = %.0f, want %d", got, resize.DefaultWidth)
	}

	app.SetSize(480, resize.DefaultHeight)
	if err := app.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	if got := boxOf(t, app, "right").Y; got <= wideRight.Y {
		t.Fatalf("narrow right column y = %.0f, want below the left column", got)
	}

	if got := boxOf(t, app, "bar").W; got != 480 {
		t.Fatalf("narrow bar = %.0f, want 480", got)
	}

	// Both these sizes stack the columns, so the paragraph gets narrower
	// and taller as the window shrinks.
	prose480 := boxOf(t, app, "prose").H

	app.SetSize(360, resize.DefaultHeight)
	if err := app.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	if got := boxOf(t, app, "prose").H; got <= prose480 {
		t.Fatalf("paragraph height at 360 = %.0f, want taller than %.0f at 480", got, prose480)
	}
}

func boxOf(t *testing.T, app *resize.App, id string) gpui.Box {
	t.Helper()

	b, ok := boxByID(app.Boxes(), id)
	if !ok {
		t.Fatalf("no box id=%q", id)
	}

	return b
}
