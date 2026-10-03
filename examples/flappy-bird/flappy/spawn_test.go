package flappy

import "testing"

func TestSpawnGapsVaryInsideField(t *testing.T) {
	app, _ := newTestApp(t)
	seen := map[float64]bool{}

	for range 200 {
		app.game.pipes = nil
		app.game.spawn(app.rng)

		p := app.game.pipes[0]
		if p.gapY-p.gap/2 < margin || p.gapY+p.gap/2 > groundY-margin {
			t.Fatalf("gap at %v leaves the field", p.gapY)
		}

		seen[p.gapY] = true
	}

	if len(seen) < 10 {
		t.Fatalf("gap centres did not vary: %d", len(seen))
	}
}

func TestSpawnGapShrinks(t *testing.T) {
	app, _ := newTestApp(t)

	cases := []struct {
		score int
		want  float64
	}{
		{0, gapStart},
		{10, gapStart - 30},
		{100, gapMin},
		{1000, gapMin},
	}

	for _, c := range cases {
		app.game.score = c.score
		app.game.pipes = nil
		app.game.spawn(app.rng)

		if got := app.game.pipes[0].gap; got != c.want {
			t.Fatalf("score %d: gap = %v, want %v", c.score, got, c.want)
		}
	}
}

func TestSpawnSpacingShrinks(t *testing.T) {
	app, _ := newTestApp(t)

	app.game.distance = 0
	start := app.game.spacing(app.rng)

	app.game.distance = 10000
	late := app.game.spacing(app.rng)

	if start < spacingStart || start > spacingStart+40 {
		t.Fatalf("early spacing = %v", start)
	}

	if late < spacingMin || late >= start || late > spacingMin+40 {
		t.Fatalf("late spacing = %v, want under %v", late, start)
	}
}
