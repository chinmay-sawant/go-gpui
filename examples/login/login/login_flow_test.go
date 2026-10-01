package login_test

import (
	"context"
	"testing"
)

func TestEmptyLoginShowsUnknown(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "", "login")

	if got := app.View().Error; got != "Unknown email or password." {
		t.Fatalf("error = %q", got)
	}

	if got := boxText(t, app, "message"); got != "Unknown email or password." {
		t.Fatalf("message = %q", got)
	}
}

func TestSubmitRejectsEmpty(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	if err := app.Submit(ctx); err != nil {
		t.Fatal(err)
	}

	if got := app.View().Error; got != "Unknown email or password." {
		t.Fatalf("error = %q", got)
	}
}
