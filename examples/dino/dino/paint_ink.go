package dino

import "github.com/chinmay-sawant/go-gpui"

// setInk places one visible obstacle fill in the dinosaur grey. The
// elements start in the background colour so the layout looks empty
// before the first tick.
func (a *App) setInk(d *gpui.Display, op *gpui.DisplayOp, r rect) {
	a.setFill(d, op, r, true)

	if op != nil {
		op.R, op.G, op.B = ink[0], ink[1], ink[2]
	}
}

// setInkUp places a bird wing, visible or folded.
func (a *App) setInkUp(d *gpui.Display, op *gpui.DisplayOp, r rect, up bool) {
	if !up {
		a.setFill(d, op, rect{}, false)

		return
	}

	a.setInk(d, op, r)
}

// hideParts hides one range of a slot's operations.
func (a *App) hideParts(d *gpui.Display, p *[partsPerSlot]*gpui.DisplayOp, from, to int) {
	for i := from; i < to; i++ {
		a.setFill(d, p[i], rect{}, false)
	}
}
