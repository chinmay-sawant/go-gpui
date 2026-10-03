package login_test

import (
	"context"
	"testing"
)

func TestEmptyLoginDoesNothing(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "", "login")

	if got := app.View().Error; got != "" {
		t.Fatalf("error = %q", got)
	}

	if got := app.View().Status; got != "" {
		t.Fatalf("status = %q", got)
	}

	if got := boxText(t, app, "message"); got != "" {
		t.Fatalf("message = %q", got)
	}
}

func TestSubmitIgnoresEmpty(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	if err := app.Submit(ctx); err != nil {
		t.Fatal(err)
	}

	if got := app.View().Error; got != "" {
		t.Fatalf("error = %q", got)
	}

	if got := app.View().Status; got != "" {
		t.Fatalf("status = %q", got)
	}
}
