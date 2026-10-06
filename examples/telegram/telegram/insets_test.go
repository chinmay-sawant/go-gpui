package telegram_test

import (
	"context"
	"testing"
)

func TestInsetsPadTheView(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	app.SetInsets(28, 32)

	if err := app.Tick(ctx); err != nil {
		t.Fatal(err)
	}

	view := app.View()
	if view.InsetTop != 28 || view.InsetBottom != 32 {
		t.Fatalf("insets = %d, %d", view.InsetTop, view.InsetBottom)
	}

	// A second tick leaves the values alone.
	if err := app.Tick(ctx); err != nil {
		t.Fatal(err)
	}

	if got := app.View().InsetTop; got != 28 {
		t.Fatalf("inset top = %d", got)
	}
}
