package dino

import "github.com/chinmay-sawant/go-gpui"

// parts caches the display operations the paint step moves. A Redraw
// replaces the display list, so bind rebuilds the cache on a new
// generation.
type parts struct {
	dino    [8]*gpui.DisplayOp
	slots   [slotMax][partsPerSlot]*gpui.DisplayOp
	clouds  [6]*gpui.DisplayOp
	pebbles [6]*gpui.DisplayOp

	score *gpui.DisplayOp
	fps   *gpui.DisplayOp
	start *gpui.DisplayOp
	keys  *gpui.DisplayOp
	over  *gpui.DisplayOp
	again *gpui.DisplayOp
}

// bind caches the operations the paint step changes.
func (a *App) bind(d *gpui.Display) {
	boxes := a.page.Boxes()

	for i, id := range dinoParts {
		a.parts.dino[i] = fillAt(d, boxes, id)
	}

	for slot := range slotMax {
		for i, part := range obstacleParts {
			a.parts.slots[slot][i] = fillAt(d, boxes, slotID(slot, part))
		}
	}

	for i, id := range cloudParts {
		a.parts.clouds[i] = fillAt(d, boxes, id)
	}

	for i, id := range pebbleParts {
		a.parts.pebbles[i] = fillAt(d, boxes, id)
	}

	a.parts.score = textAt(d, boxes, "t-score")
	a.parts.fps = textAt(d, boxes, "t-fps")
	a.parts.start = textAt(d, boxes, "t-start")
	a.parts.keys = textAt(d, boxes, "t-skeys")
	a.parts.over = textAt(d, boxes, "t-over")
	a.parts.again = textAt(d, boxes, "t-again")

	a.bound = a.page.Generation()
}
