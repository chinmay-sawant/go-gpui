package game

import "time"

// Step advances the simulation by one fixed step and reports events. A
// step on a ready, paused, or finished game does nothing.
func (g *Game) Step(step time.Duration) []Event {
	if g.Phase != PhaseRunning || !g.Piece.Valid() {
		return nil
	}

	if step <= 0 {
		step = FixedStep
	}

	g.Steps++
	g.Elapsed = min(g.Elapsed+step, MaxDuration)

	if !g.canDrop() {
		g.grounded = true
		g.fall = 0
		g.lock += step

		if g.lock >= LockDelay {
			return g.lockPiece()
		}

		return nil
	}

	g.grounded = false
	g.lock = 0
	g.fall += step

	gravity := g.gravity()
	for g.fall >= gravity && g.canDrop() {
		g.fall -= gravity
		g.Y++
		g.resets = 0
	}

	if !g.canDrop() {
		g.grounded = true
		g.fall = 0
	}

	return nil
}
