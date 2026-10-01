package login_test

import (
	"bytes"
	"context"
	"fmt"
	"image/png"
	"strings"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/login"
)

func TestEmptyLoginShowsUnknown(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "", "login")

	if got := app.View().Error; got != "Unknown email or password." {
		t.Fatalf("error = %q", got)
	}
}

func TestTypeThenLoginClearsError(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "", "login")
	if got := app.View().Error; got != "Unknown email or password." {
		t.Fatalf("error before type = %q", got)
	}

	click(t, ctx, app, "email", "focus")
	if err := app.Type(ctx, "ada@example.com"); err != nil {
		t.Fatal(err)
	}

	click(t, ctx, app, "password", "focus")
	if err := app.Type(ctx, "secret"); err != nil {
		t.Fatal(err)
	}

	if got := boxText(t, app, "email"); got != "ada@example.com" {
		t.Fatalf("email text = %q", got)
	}

	if got := boxText(t, app, "password"); got != "******" {
		t.Fatalf("password text = %q", got)
	}

	click(t, ctx, app, "", "login")

	if got := app.View().Error; got != "" {
		t.Fatalf("error after login = %q", got)
	}
}

func TestSetSizeClampsAndRedraws(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)
	before := app.Generation()

	app.SetSize(10, 10)
	width, height := app.Size()
	if width != login.MinWidth || height != login.MinHeight {
		t.Fatalf("clamped size = %d x %d", width, height)
	}

	if err := app.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	cfg, err := png.DecodeConfig(bytes.NewReader(app.PNG()))
	if err != nil {
		t.Fatal(err)
	}

	if cfg.Width != login.MinWidth || cfg.Height != login.MinHeight {
		t.Fatalf("png = %d x %d", cfg.Width, cfg.Height)
	}

	if app.Generation() <= before {
		t.Fatalf("generation = %d, before = %d", app.Generation(), before)
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

func newApp(t *testing.T, ctx context.Context) *login.App {
	t.Helper()

	app, err := login.New()
	if err != nil {
		t.Fatal(err)
	}

	if err := app.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	if len(app.PNG()) == 0 {
		t.Fatal("png is empty")
	}

	return app
}

func click(t *testing.T, ctx context.Context, app *login.App, id, action string) {
	t.Helper()

	x, y := center(t, app, id, action)
	if err := app.Click(ctx, x, y); err != nil {
		t.Fatal(err)
	}
}

// center uses the last matching box so a nested element wins over its parent.
func center(t *testing.T, app *login.App, id, action string) (float64, float64) {
	t.Helper()

	var x, y float64
	found := false

	for _, b := range app.Boxes() {
		if id != "" && b.ID != id {
			continue
		}

		if action != "" && b.Action != action {
			continue
		}

		if b.W <= 0 || b.H <= 0 {
			continue
		}

		x = b.X + b.W/2
		y = b.Y + b.H/2
		found = true
	}

	if !found {
		t.Fatalf("no box id=%q action=%q:%s", id, action, dumpBoxes(app))
	}

	return x, y
}

func boxText(t *testing.T, app *login.App, id string) string {
	t.Helper()

	var text string
	found := false

	for _, b := range app.Boxes() {
		if b.ID != id {
			continue
		}

		text = b.Text
		found = true
	}

	if !found {
		t.Fatalf("no box id=%q:%s", id, dumpBoxes(app))
	}

	return text
}

func dumpBoxes(app *login.App) string {
	var b strings.Builder

	for _, box := range app.Boxes() {
		fmt.Fprintf(
			&b,
			"\n%s id=%q action=%q text=%q xywh=%.1f,%.1f,%.1f,%.1f",
			box.Tag,
			box.ID,
			box.Action,
			box.Text,
			box.X,
			box.Y,
			box.W,
			box.H,
		)
	}

	if b.Len() == 0 {
		return "\n(no boxes)"
	}

	return b.String()
}
