package flappy

import (
	"testing"
	"time"
)

func TestHudScoreTextUpdates(t *testing.T) {
	app, clock := newTestApp(t)
	app.game.score = 7
	tick(t, app, clock, time.Second/60)

	if got := app.parts.score.Text; got != "7" {
		t.Fatalf("score text = %q, want %q", got, "7")
	}
}

func TestHudReadyTextHidesWhileRunning(t *testing.T) {
	app, clock := newTestApp(t)
	tick(t, app, clock, time.Second/60)

	if app.parts.title.Text != "FLAPPY BIRD" || app.parts.hint.Text != "SPACE OR CLICK TO FLAP" {
		t.Fatalf("ready text = %q and %q", app.parts.title.Text, app.parts.hint.Text)
	}

	press(t, app, "space")
	tick(t, app, clock, time.Second/60)

	if app.parts.title.Text != "" || app.parts.hint.Text != "" {
		t.Fatal("the ready text did not hide")
	}
}

func TestHudGameOverBoard(t *testing.T) {
	app, clock := newTestApp(t)
	app.game.score = 7
	app.game.best = 11
	app.game.phase = over
	tick(t, app, clock, time.Second/60)

	if app.parts.board == nil {
		t.Fatal("the board was not bound")
	}

	if app.parts.board.W <= 0 {
		t.Fatal("the board is hidden while over")
	}

	if app.parts.over.Text != "GAME OVER" || app.parts.again.Text != "PRESS R OR SPACE" {
		t.Fatalf("overlay = %q and %q", app.parts.over.Text, app.parts.again.Text)
	}

	if app.parts.oscore.Text != "SCORE 7" || app.parts.obest.Text != "BEST 11" {
		t.Fatalf("lines = %q and %q", app.parts.oscore.Text, app.parts.obest.Text)
	}

	if app.parts.score.Text != "" {
		t.Fatalf("score text = %q while over", app.parts.score.Text)
	}
}

func TestHudBoardHidesUntilGameOver(t *testing.T) {
	app, clock := newTestApp(t)
	tick(t, app, clock, time.Second/60)

	if app.parts.board == nil || app.parts.board.W != 0 || app.parts.board.H != 0 {
		t.Fatal("the board is not hidden while ready")
	}

	if app.parts.title.Text != "FLAPPY BIRD" {
		t.Fatalf("title text = %q while ready", app.parts.title.Text)
	}
}
