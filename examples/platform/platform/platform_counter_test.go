package platform_test

import (
	"context"
	"testing"
)

func TestClickIncrementsTheCounter(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "inc")

	if got := app.View().Count; got != 1 {
		t.Fatalf("Count = %d, want 1", got)
	}

	click(t, ctx, app, "inc")

	if got := app.View().Count; got != 2 {
		t.Fatalf("Count = %d, want 2", got)
	}
}
