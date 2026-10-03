package flappy

import "github.com/chinmay-sawant/go-gpui"

// paintScene drifts the clouds and the ground marks. The tick changes only
// the place along the sky or the ground; the layout holds the rest.
func (a *App) paintScene(d *gpui.Display) {
	for i := range cloudMax {
		setX(a.parts.clouds[i], a.clouds[i]*d.PixelPerPoint)
	}

	for i := range stripeMax {
		setX(a.parts.stripes[i], a.stripes[i]*d.PixelPerPoint)
	}
}

// setX moves one operation along the ground.
func setX(op *gpui.DisplayOp, x float64) {
	if op != nil {
		op.X = x
	}
}
