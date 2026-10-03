package dino

// dinoBox is the collision box for the current pose. It is a little
// smaller than the sprite, so near misses count as misses.
func (g *game) dinoBox() (x, y, w, h float64) {
	top := groundY - g.feet

	if g.onGround && g.ducking {
		return dinoX + 4, top - 26, 46, 26
	}

	return dinoX + 4, top - 48, 38, 48
}

// hits reports whether the dinosaur overlaps an obstacle.
func (g *game) hits() bool {
	dx, dy, dw, dh := g.dinoBox()

	for _, ob := range g.obstacles {
		ox := ob.x + 3
		oy := groundY - ob.bottom - ob.h + 3
		ow := ob.w - 6
		oh := ob.h - 6

		if dx < ox+ow && ox < dx+dw && dy < oy+oh && oy < dy+dh {
			return true
		}
	}

	return false
}
