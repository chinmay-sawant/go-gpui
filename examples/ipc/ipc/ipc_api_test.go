package ipc_test

import (
	"context"
	"errors"
	"testing"

	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/examples/ipc/ipc"
)

// TestIPCExampleAPI calls Send and Request directly, then cancels.
func TestIPCExampleAPI(t *testing.T) {
	ctx := context.Background()
	app, err := ipc.New()
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(app.Cancel)

	gpui.Send("demo.log", "one")
	if got := app.View().Received; got != "one" {
		t.Fatalf("Received = %q, want one", got)
	}

	reply, err := gpui.Request(ctx, "demo.double", "21")
	if err != nil {
		t.Fatal(err)
	}

	if reply != "42" {
		t.Fatalf("reply = %q, want 42", reply)
	}

	if _, err := gpui.Request(ctx, "demo.none", "x"); !errors.Is(err, gpui.ErrNoHandler) {
		t.Fatalf("demo.none err = %v, want ErrNoHandler", err)
	}

	app.Cancel()

	gpui.Send("demo.log", "late")
	if got := app.View().Received; got != "one" {
		t.Fatalf("Received = %q after cancel, want one", got)
	}

	if _, err := gpui.Request(ctx, "demo.double", "21"); !errors.Is(err, gpui.ErrNoHandler) {
		t.Fatalf("demo.double err = %v after cancel, want ErrNoHandler", err)
	}
}
