package resize_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/resize/resize"
)

func TestHoverControlRedraws(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app, err := resize.New()
	if err != nil {
		t.Fatal(err)
	}

	if err := app.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	btn := boxOf(t, app, "hover-btn")
	before := app.Generation()

	if err := app.Hover(ctx, btn.X+btn.W/2, btn.Y+btn.H/2); err != nil {
		t.Fatal(err)
	}

	if app.Generation() <= before {
		t.Fatalf("generation = %d, want past %d", app.Generation(), before)
	}
}
