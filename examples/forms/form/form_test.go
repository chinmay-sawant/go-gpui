package form_test

import (
	"context"
	"testing"
)

func TestSendReadsControls(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "email")
	if err := app.Type(ctx, "ada@ex.com"); err != nil {
		t.Fatal(err)
	}

	click(t, ctx, app, "remember")
	click(t, ctx, app, "pro")
	click(t, ctx, app, "color")

	click(t, ctx, app, "file")
	if err := app.Type(ctx, "notes.txt"); err != nil {
		t.Fatal(err)
	}

	click(t, ctx, app, "send")

	const want = "email=ada@ex.com remember=true plan=pro color=Blue file=notes.txt"
	if got := boxText(t, app, "status"); got != want {
		t.Fatalf("status = %q", got)
	}
}
