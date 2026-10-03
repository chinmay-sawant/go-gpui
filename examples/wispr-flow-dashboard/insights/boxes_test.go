package insights

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/go-gpui"
)

func TestBoxesCoverTheDashboard(t *testing.T) {
	app, err := New()
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	if err := app.Redraw(context.Background()); err != nil {
		t.Fatalf("Redraw: %v", err)
	}

	boxes := app.Boxes()

	wpm, ok := findBox(boxes, "card-wpm")
	if !ok {
		t.Fatal("no card-wpm box")
	}

	fixes, ok := findBox(boxes, "card-fixes")
	if !ok {
		t.Fatal("no card-fixes box")
	}

	words, ok := findBox(boxes, "card-words")
	if !ok {
		t.Fatal("no card-words box")
	}

	apps, ok := findBox(boxes, "card-apps")
	if !ok {
		t.Fatal("no card-apps box")
	}

	streak, ok := findBox(boxes, "card-streak")
	if !ok {
		t.Fatal("no card-streak box")
	}

	if wpm.Y != fixes.Y || wpm.Y != words.Y {
		t.Fatalf("top cards y = %v %v %v", wpm.Y, fixes.Y, words.Y)
	}

	if wpm.W > 250 || words.W < 480 {
		t.Fatalf("top card widths = %v %v", wpm.W, words.W)
	}

	if apps.Y != streak.Y || apps.Y <= wpm.Y {
		t.Fatalf("bottom cards y = %v %v", apps.Y, streak.Y)
	}

	if apps.W < 480 || streak.W < 480 {
		t.Fatalf("bottom card widths = %v %v", apps.W, streak.W)
	}
}

func TestClickSwitchesTab(t *testing.T) {
	app, err := New()
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	if err := app.Redraw(context.Background()); err != nil {
		t.Fatalf("Redraw: %v", err)
	}

	tab, ok := findBox(app.Boxes(), "tab-voice")
	if !ok {
		t.Fatal("no tab-voice box")
	}

	if err := app.Click(context.Background(), tab.X+tab.W/2, tab.Y+tab.H/2); err != nil {
		t.Fatalf("Click: %v", err)
	}

	if got := app.View().ActiveTab; got != "voice" {
		t.Fatalf("ActiveTab = %q", got)
	}

	if got := app.View().EmptyTitle; got != "Your voice" {
		t.Fatalf("EmptyTitle = %q", got)
	}
}

func findBox(boxes []gpui.Box, id string) (gpui.Box, bool) {
	for _, box := range boxes {
		if box.ID == id {
			return box, true
		}
	}

	return gpui.Box{}, false
}
