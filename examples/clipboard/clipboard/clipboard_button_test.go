package clipboard_test

import (
	"context"
	"strings"
	"testing"
)

func TestButtonsWriteStatusAndText(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	click(t, ctx, app, "copy")
	if got := app.View().Status; !strings.Contains(got, "nothing selected") {
		t.Fatalf("empty copy status = %q", got)
	}

	click(t, ctx, app, "left")
	if err := app.Type(ctx, "hello"); err != nil {
		t.Fatal(err)
	}

	click(t, ctx, app, "copy")
	if got := app.View().Status; !strings.Contains(got, "copied: hello") {
		t.Fatalf("copy status = %q", got)
	}

	click(t, ctx, app, "cut")
	if got := app.FormValue("left"); got != "" {
		t.Fatalf("after cut = %q, want empty", got)
	}

	if got := app.View().Status; !strings.Contains(got, "cut: hello") {
		t.Fatalf("cut status = %q", got)
	}

	click(t, ctx, app, "pastego")
	if got := app.FormValue("left"); !strings.Contains(got, "from Go") {
		t.Fatalf("after pastego = %q, want from Go", got)
	}

	click(t, ctx, app, "selectall")
	if !app.Page().FormSelected("left") {
		t.Fatal("selectall did not select the left field")
	}
}
