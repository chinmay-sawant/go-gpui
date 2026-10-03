package dino

import "testing"

func TestBirdHeights(t *testing.T) {
	app, _ := newTestApp(t)
	app.game.phase = running
	app.game.score = 400

	heights := map[float64]bool{}

	for range 200 {
		app.game.obstacles = nil
		app.game.spawn(app.rng)

		if app.game.obstacles[0].kind == bird {
			heights[app.game.obstacles[0].bottom] = true
		}
	}

	if !heights[6] || !heights[36] {
		t.Fatalf("bird heights = %v, want both low and high", heights)
	}
}

func TestSpawnOnlyUsesKnownKinds(t *testing.T) {
	app, _ := newTestApp(t)
	app.game.phase = running

	for range 300 {
		app.game.obstacles = nil
		app.game.spawn(app.rng)

		switch app.game.obstacles[0].kind {
		case cactusSmall, cactusBig, bird:
		default:
			t.Fatalf("kind = %d", app.game.obstacles[0].kind)
		}
	}
}
