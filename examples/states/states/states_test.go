package states_test

import (
	"bytes"
	"context"
	"testing"
)

func TestStatesAndPointerInput(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)

	plain := app.PNG()
	if len(plain) == 0 {
		t.Fatal("png is empty after the first redraw")
	}

	assertReplayed(t, app)

	// The focus-visible outline is a CSS outline op, once a bitmap-only
	// feature. The page must stay on its display list in every state below,
	// and PNG still changes because it rasterizes with the current state.
	clickBox(t, ctx, app, "focus-input")
	if got := app.FocusedField(); got != "focus-input" {
		t.Fatalf("focused = %q, want focus-input", got)
	}

	focused := app.PNG()
	if bytes.Equal(focused, plain) {
		t.Fatal("focus did not change the png")
	}

	assertReplayed(t, app)

	x, y := boxCenter(t, app, "hover-btn")

	before := app.Generation()
	if err := app.Hover(ctx, x, y); err != nil {
		t.Fatal(err)
	}

	if app.Generation() <= before {
		t.Fatalf("hover generation = %d, want past %d", app.Generation(), before)
	}

	if hovered := app.PNG(); bytes.Equal(hovered, focused) {
		t.Fatal("hover did not change the png")
	}

	assertReplayed(t, app)

	px, py := boxCenter(t, app, "press-btn")
	if err := app.Press(ctx, px, py); err != nil {
		t.Fatal(err)
	}

	assertReplayed(t, app)

	if err := app.Release(ctx); err != nil {
		t.Fatal(err)
	}

	beforeCheck := app.PNG()
	clickBox(t, ctx, app, "agree")

	if !app.FormChecked("agree") {
		t.Fatal("agree is not checked after the click")
	}

	if checked := app.PNG(); bytes.Equal(checked, beforeCheck) {
		t.Fatal("checked did not change the png")
	}

	assertReplayed(t, app)
}
