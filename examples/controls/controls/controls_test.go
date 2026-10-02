package controls_test

import (
	"context"
	"strings"
	"testing"

	"github.com/chinmay-sawant/go-gpui/examples/controls/controls"
)

func TestControlsRender(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	if len(app.PNG()) == 0 {
		t.Fatal("PNG is empty")
	}
}

func TestTypingFocusAndDelete(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "name")
	if got := app.FocusedField(); got != "name" {
		t.Fatalf("FocusedField = %q, want name", got)
	}

	if err := app.Type(ctx, "one two"); err != nil {
		t.Fatal(err)
	}

	if err := app.DeleteWord(ctx); err != nil {
		t.Fatal(err)
	}

	if got := app.FormValue("name"); got != "one " {
		t.Fatalf("after DeleteWord = %q, want %q", got, "one ")
	}

	if err := app.Backspace(ctx); err != nil {
		t.Fatal(err)
	}

	if got := app.FormValue("name"); got != "one" {
		t.Fatalf("after Backspace = %q, want one", got)
	}

	if got := boxText(t, app, "status"); !strings.Contains(got, "name=one") {
		t.Fatalf("status = %q, want name=one", got)
	}
}

func newApp(t *testing.T, ctx context.Context) *controls.App {
	t.Helper()

	app, err := controls.New()
	if err != nil {
		t.Fatal(err)
	}

	app.SetSize(controls.DefaultWidth, controls.DefaultHeight)
	if err := app.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	return app
}
