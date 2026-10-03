package bindhooks_test

import (
	"context"
	"strings"
	"testing"
)

func TestBeforeEditVetoesLockedName(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "name")

	if err := app.Type(ctx, "new"); err == nil {
		t.Fatal("Type on the locked name returned nil, want a veto error")
	}

	if got := app.FormValue("name"); got != "locked" {
		t.Fatalf("FormValue = %q, want locked", got)
	}

	if got := app.View().Name; got != "locked" {
		t.Fatalf("Name = %q, want locked", got)
	}

	if got := app.View().Status; got != "" {
		t.Fatalf("Status = %q, want unchanged and empty", got)
	}
}

func TestAgreeAndPlanWriteView(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "pro")
	if got := app.View().Plan; got != "pro" {
		t.Fatalf("Plan = %q, want pro", got)
	}

	if got := app.View().Status; !strings.Contains(got, "pro") {
		t.Fatalf("Status = %q, want a note about pro", got)
	}

	click(t, ctx, app, "free")
	if got := app.View().Plan; got != "free" {
		t.Fatalf("Plan = %q, want free", got)
	}

	click(t, ctx, app, "agree")
	if !app.View().Agree {
		t.Fatal("Agree = false, want true")
	}

	if got := app.View().Status; !strings.Contains(got, "agree") {
		t.Fatalf("Status = %q, want a note about agree", got)
	}
}
