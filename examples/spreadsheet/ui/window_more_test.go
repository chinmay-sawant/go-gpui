package ui

import (
	"context"
	"testing"
)

func TestCtrlArrowDownWithTrackedModifier(t *testing.T) {
	app := newTestApp(t, newFake())
	settle(t, app)

	ctx := context.Background()
	if err := app.onKeyDown(ctx, "control"); err != nil {
		t.Fatal(err)
	}

	if err := app.onKeyDown(ctx, "arrowdown"); err != nil {
		t.Fatal(err)
	}

	if s := app.selection(); s.ActiveR != 199 {
		t.Fatalf("active row = %d", s.ActiveR)
	}
}

func TestHorizontalWindow(t *testing.T) {
	app := newTestApp(t, newFake())
	settle(t, app)

	app.page.SetScrollOffset(1000, 0)
	app.page.StepScrollWindow()
	flush(t, app)
	if err := app.Redraw(context.Background()); err != nil {
		t.Fatal(err)
	}

	if app.win.C0 == 0 {
		t.Fatalf("window did not move horizontally: %+v", app.win)
	}

	if b, ok := cellBox(app.View(), 0, app.win.C0); !ok || b.X != cellX(app.win.C0) {
		t.Fatalf("stable column coordinates: %+v ok=%v", b, ok)
	}
}
