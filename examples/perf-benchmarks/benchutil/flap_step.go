package benchutil

// Flap tuning is in CSS pixels and seconds, after the flappy-bird example.
const (
	flapGravity = 1700.0
	flapLift    = -520.0
	flapSpeed   = 150.0
	flapGap     = 190.0
	flapPipeW   = 62.0
	flapBirdX   = 130.0
	flapBirdW   = 34.0
	flapBirdH   = 24.0
	flapGroundY = 624.0
	flapSceneH  = 720.0
)

// Slot tuning: travel between spawns, spawn position, and drop edge.
const (
	flapTravel = 280.0
	flapSpawnX = 490.0
	flapDropX  = -40.0
)

// gaps holds the deterministic spawn gap centers.
var gaps = [4]float64{380, 250, 320, 270}

// Step advances the bird and the pipes by dt seconds: it moves, spawns,
// drops, scores, and resets on a crash, so the demo never needs a Redraw.
func (f *Flap) Step(dt float64) {
	f.vy += flapGravity * dt
	f.y += f.vy * dt

	for i := range f.pipes {
		f.pipes[i].x -= flapSpeed * dt
	}

	f.next -= flapSpeed * dt

	if f.next <= 0 && len(f.pipes) < 3 {
		f.spawn()
		f.next = flapTravel
	}

	kept := f.pipes[:0]

	for _, p := range f.pipes {
		if p.x+flapPipeW > flapDropX {
			kept = append(kept, p)
		}
	}

	f.pipes = kept
	f.scorePass()

	if f.crashed() {
		f.reset()
	}
}

// spawn adds one pair at the right edge. Its slot op is already a fill, so
// starting off-screen is fine; only layout-time clipping deactivates.
func (f *Flap) spawn() {
	f.pipes = append(f.pipes, flapPipe{x: flapSpawnX, gapY: gaps[f.spawns%len(gaps)]})
	f.spawns++
}

// scorePass counts a pipe once when its right edge passes the bird.
func (f *Flap) scorePass() {
	for i := range f.pipes {
		p := &f.pipes[i]

		if !p.scored && p.x+flapPipeW < flapBirdX {
			p.scored = true
			f.score++
		}
	}
}

// crashed reports a pipe, ceiling, or ground hit.
func (f *Flap) crashed() bool {
	if f.y <= 0 || f.y+flapBirdH >= flapGroundY {
		return true
	}

	for _, p := range f.pipes {
		if flapBirdX+flapBirdW < p.x || flapBirdX > p.x+flapPipeW {
			continue
		}

		if f.y < p.gapY-flapGap/2 || f.y+flapBirdH > p.gapY+flapGap/2 {
			return true
		}
	}

	return false
}
