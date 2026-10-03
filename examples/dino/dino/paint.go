package dino

import "github.com/chinmay-sawant/go-gpui"

// paint moves the cached operations to match the game state. It runs every
// tick and never parses the HTML again.
func (a *App) paint() {
	d := a.page.Display()
	if d == nil {
		return
	}

	if a.page.Generation() != a.bound {
		a.bind(d)
	}

	a.paintDino(d)
	a.paintObstacles(d)
	a.paintClouds(d)
	a.paintPebbles(d)
	a.paintText()
}

// paintDino places the eight dinosaur fills for the current pose.
func (a *App) paintDino(d *gpui.Display) {
	pose := a.game.dinoPose()

	for i := range dinoParts {
		setFill(d, a.parts.dino[i], pose[i], true)
	}
}

// setFill places one fill at r and shows or hides it.
func setFill(d *gpui.Display, op *gpui.DisplayOp, r rect, visible bool) {
	if op == nil {
		return
	}

	op.X = r.x * d.PixelPerPoint
	op.Y = r.y * d.PixelPerPoint
	op.W = r.w * d.PixelPerPoint
	op.H = r.h * d.PixelPerPoint

	if visible {
		op.Alpha = 1

		return
	}

	// The replay reads opacity, not Alpha, so a hidden fill collapses to
	// nothing instead.
	op.W, op.H = 0, 0
	op.Alpha = 0
}
