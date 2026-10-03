package flappy

// flap starts the run from the ready screen, lifts the bird while running,
// and restarts after a crash.
func (g *game) flap() {
	switch g.phase {
	case ready:
		g.phase = running
		g.vy = flapV
		g.flapAge = 0
	case over:
		g.again()
	default:
		g.vy = flapV
		g.flapAge = 0
	}
}

// again starts a fresh run after a crash: a reset plus the first flap.
func (g *game) again() {
	g.restart()
	g.phase = running
	g.vy = flapV
	g.flapAge = 0
}

// restart resets the run and keeps the best score.
func (g *game) restart() {
	best := max(g.best, g.score)
	*g = newGame()
	g.best = best
}
