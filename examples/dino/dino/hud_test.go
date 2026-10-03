package dino

import (
	"strings"
	"testing"
	"time"
)

func TestScoreTextUpdates(t *testing.T) {
	app, clock := newTestApp(t)
	press(t, app, "space")
	app.game.distance = 123 * 18
	tick(t, app, clock, time.Second/60)

	if got := app.parts.score.Text; !strings.Contains(got, "00123") {
		t.Fatalf("score text = %q", got)
	}
}

func TestFPSStartsAfterTheFirstWindow(t *testing.T) {
	app, clock := newTestApp(t)
	tick(t, app, clock, time.Second/60)

	if got := app.parts.fps.Text; got != "060 FPS" {
		t.Fatalf("fps text = %q before a window", got)
	}

	for range 30 {
		tick(t, app, clock, time.Second/30)
	}

	if got := app.parts.fps.Text; got != "030 FPS" {
		t.Fatalf("fps text = %q after a window", got)
	}
}

func TestOverlaysFollowThePhase(t *testing.T) {
	app, clock := newTestApp(t)
	tick(t, app, clock, time.Second/60)

	if app.parts.start.Text != startText || app.parts.over.Text != "" {
		t.Fatal("the ready overlay is wrong")
	}

	press(t, app, "space")
	tick(t, app, clock, time.Second/60)

	if app.parts.start.Text != "" || app.parts.keys.Text != "" {
		t.Fatal("the ready overlay did not hide")
	}

	app.game.phase = over
	tick(t, app, clock, time.Second/60)

	if app.parts.over.Text != overText || app.parts.again.Text != againText {
		t.Fatal("the game-over overlay is missing")
	}

	if app.parts.start.Text != "" {
		t.Fatal("the ready overlay came back")
	}
}
