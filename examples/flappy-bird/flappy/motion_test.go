package flappy

import (
	"math"
	"testing"
	"time"
)

func TestFlapRisesThenFalls(t *testing.T) {
	app, clock := newTestApp(t)
	press(t, app, "space")

	mid := app.game.birdY

	for range 10 {
		tick(t, app, clock, time.Second/60)
	}

	if app.game.birdY >= mid {
		t.Fatalf("the flap did not rise: %v", app.game.birdY)
	}

	for range 120 {
		tick(t, app, clock, time.Second/60)
	}

	if app.game.birdY <= mid {
		t.Fatalf("gravity did not bring the bird down: %v", app.game.birdY)
	}
}

func TestReadyBob(t *testing.T) {
	app, clock := newTestApp(t)

	for range 120 {
		tick(t, app, clock, time.Second/60)

		if d := math.Abs(app.game.birdY - sceneH/2); d > 8.001 {
			t.Fatalf("bob = %v px from centre", d)
		}
	}

	if app.game.phase != ready {
		t.Fatalf("the idle screen left ready: %v", app.game.phase)
	}
}

func TestReadySlowDTCapped(t *testing.T) {
	app, clock := newTestApp(t)
	tick(t, app, clock, time.Second)
	tick(t, app, clock, 2*time.Second)

	if math.Abs(app.game.run-0.1) > 1e-9 {
		t.Fatalf("run = %v after a slow frame, want 0.1", app.game.run)
	}
}
