package login_test

import (
	"context"
	"testing"
)

func TestPasswordCopyAndDeleteWord(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "password", "")
	if err := app.Type(ctx, "secret"); err != nil {
		t.Fatal(err)
	}

	text, ok, err := app.Page().Copy(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if !ok || text != "secret" {
		t.Fatalf("password copy = %q ok=%v", text, ok)
	}

	if err := app.Page().DeleteWord(ctx); err != nil {
		t.Fatal(err)
	}

	if got := app.Page().FormValue("password"); got != "" {
		t.Fatalf("password after delete word = %q", got)
	}
}
