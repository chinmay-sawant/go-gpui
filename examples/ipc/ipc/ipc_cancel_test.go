package ipc_test

import (
	"context"
	"strings"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/ipc/ipc"
)

// TestIPCExampleCancel drives Cancel and Register again through the button.
func TestIPCExampleCancel(t *testing.T) {
	ctx := context.Background()
	app, err := ipc.New()
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(app.Cancel)

	if err := app.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	click(t, ctx, app, "send")
	if got := app.View().Sends; got != 1 {
		t.Fatalf("Sends = %d before cancel, want 1", got)
	}

	click(t, ctx, app, "cancel")
	if app.View().Live {
		t.Fatal("Live = true after cancel")
	}

	if got := app.View().Status; !strings.Contains(got, "removed") {
		t.Fatalf("Status = %q, want the removal line", got)
	}

	click(t, ctx, app, "send")
	if got := app.View().Sends; got != 1 {
		t.Fatalf("Sends = %d after cancel, want 1", got)
	}

	if got := app.View().Status; !strings.Contains(got, "no listeners") {
		t.Fatalf("Status = %q, want the dropped message", got)
	}

	click(t, ctx, app, "cancel")
	if !app.View().Live {
		t.Fatal("Live = false after register again")
	}

	click(t, ctx, app, "send")
	if got := app.View().Sends; got != 2 {
		t.Fatalf("Sends = %d after register again, want 2", got)
	}

	if got := app.View().Events; len(got) == 0 || !strings.Contains(got[0], "Send demo.log") {
		t.Fatalf("Events top = %q, want the newest line first", got)
	}
}
