package dino

import "math/rand/v2"

// gap is the travel before the next obstacle, long enough to clear at the
// current speed.
func (g *game) gap(rng *rand.Rand) float64 {
	return g.speed * (0.62 + rng.Float64()*0.55)
}

// spawn adds one obstacle at the right edge. Birds arrive after the first
// few cacti.
func (g *game) spawn(rng *rand.Rand) {
	kind := cactusSmall
	roll := rng.Float64()

	switch {
	case g.score > 250 && roll < 0.22:
		kind = bird
	case roll < 0.6:
		kind = cactusBig
	}

	ob := obstacle{kind: kind, x: sceneW + 20}

	switch kind {
	case bird:
		ob.w, ob.h = 40, 28
		ob.bottom = 6

		if rng.Float64() < 0.5 {
			ob.bottom = 36
		}
	case cactusBig:
		ob.w, ob.h = 30, 48
	default:
		ob.w, ob.h = 22, 32
	}

	g.obstacles = append(g.obstacles, ob)
}
