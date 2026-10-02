package bindhooks_test

import (
	"context"
	"strings"
	"testing"

	"github.com/chinmay-sawant/go-gpui/examples/bind-hooks/bindhooks"
)

func TestBindHooksRender(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	if len(app.PNG()) == 0 {
		t.Fatal("PNG is empty")
	}
}

func TestEmailChangeWritesViewAndStatus(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "email")
	if err := app.Type(ctx, ".net"); err != nil {
		t.Fatal(err)
	}

	if got := app.View().Email; got != "you@example.com.net" {
		t.Fatalf("Email = %q, want you@example.com.net", got)
	}

	if got := app.View().Status; !strings.Contains(got, "email") {
		t.Fatalf("Status = %q, want a note about email", got)
	}
}

func newApp(t *testing.T, ctx context.Context) *bindhooks.App {
	t.Helper()

	app, err := bindhooks.New()
	if err != nil {
		t.Fatal(err)
	}

	app.SetSize(bindhooks.DefaultWidth, bindhooks.DefaultHeight)
	if err := app.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	return app
}

// click uses the last matching box so a nested element wins over its parent.
func click(t *testing.T, ctx context.Context, app *bindhooks.App, id string) {
	t.Helper()

	var x, y float64
	found := false

	for _, b := range app.Boxes() {
		if b.ID != id || b.W <= 0 || b.H <= 0 {
			continue
		}

		x = b.X + b.W/2
		y = b.Y + b.H/2
		found = true
	}

	if !found {
		t.Fatalf("no box id=%q", id)
	}

	if err := app.Click(ctx, x, y); err != nil {
		t.Fatal(err)
	}
}
