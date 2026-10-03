package flappy

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

	a.paintScene(d)
	a.paintBird(d)
	a.paintPipes(d)
	a.paintText(d)
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
	// nothing and moves off the canvas: a zero-size rounded rect still
	// paints a corner sparkle at its place.
	op.X, op.Y = -1000, -1000
	op.W, op.H = 0, 0
	op.Alpha = 0
}
