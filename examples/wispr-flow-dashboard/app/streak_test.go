package app

import (
	"context"
	"reflect"
	"testing"

	"github.com/chinmay-sawant/go-gpui"
)

// TestStreakChevronsPageWeeks checks the newest window starts with the right
// chevron disabled, and that paging back and forward moves the heatmap.
func TestStreakChevronsPageWeeks(t *testing.T) {
	app, err := New()
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	openInsights(t, app)

	next, ok := findBox(app.Boxes(), "streak-next")
	if !ok {
		t.Fatal("no streak-next box")
	}

	if next.Action != "" {
		t.Fatal("the right chevron should start disabled")
	}

	prev, ok := findBox(app.Boxes(), "streak-prev")
	if !ok {
		t.Fatal("no streak-prev box")
	}

	if prev.Action != "streak-prev" {
		t.Fatalf("prev action = %q", prev.Action)
	}

	newest := app.View().Streak.Weeks

	clickBox(t, app, prev)

	if !app.View().Streak.Next {
		t.Fatal("right chevron did not enable after paging back")
	}

	if reflect.DeepEqual(app.View().Streak.Weeks, newest) {
		t.Fatal("weeks did not move")
	}

	next, ok = findBox(app.Boxes(), "streak-next")
	if !ok {
		t.Fatal("no streak-next box after paging")
	}

	clickBox(t, app, next)

	if !reflect.DeepEqual(app.View().Streak.Weeks, newest) {
		t.Fatal("paging forward did not restore the newest window")
	}
}

func clickBox(t *testing.T, app *App, box gpui.Box) {
	t.Helper()

	if err := app.Click(context.Background(), box.X+box.W/2, box.Y+box.H/2); err != nil {
		t.Fatal(err)
	}
}
