package controls_test

import (
	"context"
	"strings"
	"testing"
)

func TestToggleRadioAndSelect(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "agree")
	if !app.FormChecked("agree") {
		t.Fatal("agree = false, want true")
	}

	click(t, ctx, app, "pro")
	if !app.FormChecked("pro") {
		t.Fatal("pro = false, want true")
	}

	if app.FormChecked("free") {
		t.Fatal("free = true, want false after checking pro")
	}

	if got := app.FocusedField(); got != "pro" {
		t.Fatalf("FocusedField = %q, want pro", got)
	}

	click(t, ctx, app, "color")
	if got := app.FormValue("color"); got != "green" {
		t.Fatalf("color = %q, want green", got)
	}

	click(t, ctx, app, "color")
	click(t, ctx, app, "color")
	if got := app.FormValue("color"); got != "red" {
		t.Fatalf("color = %q, want red after the third click", got)
	}

	status := boxText(t, app, "status")
	for _, want := range []string{"plan=pro", "color=red"} {
		if !strings.Contains(status, want) {
			t.Fatalf("status = %q, want %q", status, want)
		}
	}
}

func TestTextareaAndFileTakeTyping(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "bio")
	if err := app.Type(ctx, "hello"); err != nil {
		t.Fatal(err)
	}

	if got := app.FormValue("bio"); got != "hello" {
		t.Fatalf("bio = %q, want hello", got)
	}

	click(t, ctx, app, "doc")
	if err := app.Type(ctx, "notes.txt"); err != nil {
		t.Fatal(err)
	}

	if got := app.FormValue("doc"); got != "notes.txt" {
		t.Fatalf("doc = %q, want notes.txt", got)
	}
}
