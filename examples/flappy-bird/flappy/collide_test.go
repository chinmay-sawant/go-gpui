package flappy

import (
	"testing"
	"time"
)

func TestPipeColumnHits(t *testing.T) {
	app, _ := newTestApp(t)
	app.game.birdY = 300
	app.game.pipes = []pipe{{x: birdX, gapY: 500, gap: gapStart}}

	if !app.game.hits() {
		t.Fatal("a column at the bird's height missed")
	}
}

func TestPipeGapPasses(t *testing.T) {
	app, _ := newTestApp(t)
	app.game.birdY = 300
	app.game.pipes = []pipe{{x: birdX, gapY: 300, gap: gapStart}}

	if app.game.hits() {
		t.Fatal("a gap at the bird's height hit")
	}
}

func TestGroundHits(t *testing.T) {
	app, _ := newTestApp(t)
	app.game.birdY = groundY - 5

	if !app.game.hits() {
		t.Fatal("the ground did not hit")
	}
}

func TestCeilingClampDoesNotHit(t *testing.T) {
	app, clock := newTestApp(t)
	press(t, app, "space")
	app.game.birdY, app.game.vy = 2, -100

	tick(t, app, clock, time.Second/60)

	if app.game.phase != running || app.game.birdY != birdR || app.game.vy != 0 {
		t.Fatalf("ceiling: phase = %v, y = %v, vy = %v", app.game.phase, app.game.birdY, app.game.vy)
	}
}

func TestNearMissMisses(t *testing.T) {
	app, _ := newTestApp(t)
	app.game.birdY = 300
	app.game.pipes = []pipe{{x: birdX, gapY: 300 - 11 + gapStart/2, gap: gapStart}}

	if app.game.hits() {
		t.Fatal("a grazing near miss hit")
	}
}
