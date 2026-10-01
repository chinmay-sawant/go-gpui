package login_test

import (
	"context"
	"testing"
)

func TestTypeThenLoginClearsError(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "", "login")
	if got := app.View().Error; got != "Unknown email or password." {
		t.Fatalf("error before type = %q", got)
	}

	click(t, ctx, app, "email", "focus")
	if err := app.Type(ctx, "secret"); err != nil {
		t.Fatal(err)
	}

	click(t, ctx, app, "password", "focus")
	if err := app.Type(ctx, "secret"); err != nil {
		t.Fatal(err)
	}

	if got := boxText(t, app, "email"); got != "secret" {
		t.Fatalf("email text = %q", got)
	}

	if got := boxText(t, app, "password"); got != "******" {
		t.Fatalf("password text = %q", got)
	}

	click(t, ctx, app, "", "login")

	if got := app.View().Error; got != "" {
		t.Fatalf("error after login = %q", got)
	}

	if got := app.View().Status; got != "Signed in." {
		t.Fatalf("status = %q", got)
	}

	if got := boxText(t, app, "message"); got != "Signed in." {
		t.Fatalf("message = %q", got)
	}
}
