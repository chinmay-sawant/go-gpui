package flappy

import (
	"testing"
	"time"
)

func TestRunMovesPipesAndScores(t *testing.T) {
	app, clock := newTestApp(t)
	press(t, app, "space")

	app.game.pipes = []pipe{
		{x: birdX - pipeW - 1, gapY: sceneH / 2, gap: gapStart},
		{x: 300, gapY: sceneH / 2, gap: gapStart},
	}

	tick(t, app, clock, time.Second/60)
	tick(t, app, clock, time.Second/60)

	if app.game.pipes[1].x >= 300 {
		t.Fatalf("the pipe did not move left: %v", app.game.pipes[1].x)
	}

	if app.game.score != 1 {
		t.Fatalf("score = %d, want 1", app.game.score)
	}
}

func TestSpeedGrowsWithDistance(t *testing.T) {
	app, _ := newTestApp(t)
	app.game.phase = running

	cases := []struct {
		distance, want float64
	}{
		{0, speedStart},
		{2500, 200},
		{10000, speedMax},
	}

	for _, c := range cases {
		app.game.distance = c.distance
		app.game.step(0, app.rng)

		if app.game.speed != c.want {
			t.Fatalf("distance %v: speed = %v, want %v", c.distance, app.game.speed, c.want)
		}
	}
}
