package flappy

import (
	"testing"
	"time"
)

func TestCrashAndRestart(t *testing.T) {
	app, clock := newTestApp(t)
	press(t, app, "space")

	app.game.score = 7
	app.game.pipes = []pipe{{x: 400, gapY: sceneH / 2, gap: gapStart}}
	app.game.birdY = groundY - 5
	tick(t, app, clock, time.Second/60)

	if app.game.phase != over || app.game.best != 7 {
		t.Fatalf("crash: phase = %v, best = %d", app.game.phase, app.game.best)
	}

	press(t, app, "space")

	if app.game.phase != running || app.game.score != 0 || app.game.best != 7 {
		t.Fatalf("again: phase = %v, score = %d, best = %d", app.game.phase, app.game.score, app.game.best)
	}

	if len(app.game.pipes) != 0 || app.game.birdY != sceneH/2 || app.game.run != 0 {
		t.Fatalf("again did not reset: %+v", app.game)
	}
}

func TestRestartKeepsBest(t *testing.T) {
	app, _ := newTestApp(t)
	app.game.best, app.game.score = 3, 9
	app.game.pipes = []pipe{{x: 200, gapY: 300, gap: gapStart}}
	app.game.distance, app.game.birdY = 500, 100

	app.game.restart()

	if app.game.best != 9 || app.game.score != 0 || app.game.pipes != nil {
		t.Fatalf("restart: %+v", app.game)
	}

	if app.game.phase != ready || app.game.birdY != sceneH/2 || app.game.distance != 0 {
		t.Fatalf("restart did not reset: %+v", app.game)
	}
}
