package dino

import "github.com/chinmay-sawant/go-gpui"

// setInk places one visible obstacle fill in the dinosaur grey. The
// elements start in the background colour so the layout looks empty
// before the first tick.
func setInk(d *gpui.Display, op *gpui.DisplayOp, r rect) {
	setFill(d, op, r, true)

	if op != nil {
		op.R, op.G, op.B = ink[0], ink[1], ink[2]
	}
}

// setInkUp places a bird wing, visible or folded.
func setInkUp(d *gpui.Display, op *gpui.DisplayOp, r rect, up bool) {
	if !up {
		setFill(d, op, rect{}, false)

		return
	}

	setInk(d, op, r)
}

// hideParts hides one range of a slot's operations.
func hideParts(d *gpui.Display, p *[partsPerSlot]*gpui.DisplayOp, from, to int) {
	for i := from; i < to; i++ {
		setFill(d, p[i], rect{}, false)
	}
}
