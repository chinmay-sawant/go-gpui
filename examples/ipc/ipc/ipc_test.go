package ipc_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/examples/ipc/ipc"
)

func TestIPCExample(t *testing.T) {
	ctx := context.Background()
	app, err := ipc.New()
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(app.Cancel)

	if err := app.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	if len(app.Page().PNG()) == 0 {
		t.Fatal("PNG is empty after Redraw")
	}

	gpui.Send("demo.log", "one")
	if got := app.View().Log; !strings.Contains(got, "one") {
		t.Fatalf("Log = %q, want one", got)
	}

	reply, err := gpui.Request(ctx, "demo.double", "21")
	if err != nil {
		t.Fatal(err)
	}

	if reply != "42" {
		t.Fatalf("reply = %q, want 42", reply)
	}

	if _, err := gpui.Request(ctx, "demo.none", "x"); err == nil {
		t.Fatal("Request demo.none err = nil, want ErrNoHandler")
	} else if !errors.Is(err, gpui.ErrNoHandler) {
		t.Fatalf("err = %v, want ErrNoHandler", err)
	}

	app.Page().SetFormValue("num", "21")
	click(t, ctx, app, "request")
	if got := app.View().Status; got != "42" {
		t.Fatalf("Status = %q, want 42", got)
	}

	click(t, ctx, app, "missing")
	if got := app.View().Status; !strings.Contains(got, "no ipc handler") {
		t.Fatalf("Status = %q, want the ErrNoHandler text", got)
	}

	click(t, ctx, app, "send")
	if got := app.View().Log; !strings.Contains(got, "ping") {
		t.Fatalf("Log = %q, want ping", got)
	}

	click(t, ctx, app, "cancel")
	gpui.Send("demo.log", "late")
	if got := app.View().Log; strings.Contains(got, "late") {
		t.Fatalf("Log = %q, want no late after cancel", got)
	}
}

func click(t *testing.T, ctx context.Context, app *ipc.App, id string) {
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
