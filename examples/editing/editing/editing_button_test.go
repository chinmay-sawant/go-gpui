package editing_test

import (
	"context"
	"testing"
)

// The select-all button must select the note even though the button click
// blurs the form first.
func TestSelectAllButtonSelectsNote(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "note")
	if err := app.Type(ctx, "abc"); err != nil {
		t.Fatal(err)
	}

	click(t, ctx, app, "selectall")
	if !app.FormSelected("note") {
		t.Fatal("selectall did not select the note")
	}

	if err := app.Type(ctx, "z"); err != nil {
		t.Fatal(err)
	}

	if got := app.FormValue("note"); got != "z" {
		t.Fatalf("note = %q, want z", got)
	}
}
