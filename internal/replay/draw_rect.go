package replay

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

// DrawRect replays the operations of display whose painted box meets rect,
// in the same paint order Draw uses. rect is in CSS pixels in the page's own
// space; dx and dy only move the paint, as they do in Draw. The rect is
// clamped to the canvas, so one that reaches outside still paints the part
// that lands inside.
func DrawRect(dst *ebiten.Image, display *layout.Display, rect image.Rectangle, dx, dy float64) {
	if display == nil {
		return
	}

	rect = clipRect(display, rect)
	if rect.Empty() {
		return
	}

	for _, index := range display.Order {
		if index < 0 || index >= len(display.Ops) {
			continue
		}

		op := &display.Ops[index]
		if !opTouches(op, display.PixelPerPoint, rect) {
			continue
		}

		drawOp(dst, op, dx, dy)
	}
}

// canvasRect is the page canvas in CSS pixels.
func canvasRect(display *layout.Display) image.Rectangle {
	if display.Width <= 0 || display.Height <= 0 {
		return image.Rectangle{}
	}

	return image.Rect(0, 0, display.Width, display.Height)
}

// clipRect clamps rect to the page canvas.
func clipRect(display *layout.Display, rect image.Rectangle) image.Rectangle {
	return rect.Intersect(canvasRect(display))
}

// opTouches reports whether an op's painted box meets rect. An op whose ink
// cannot be bounded, such as a rotated run, always touches.
func opTouches(op *layout.DisplayOp, ppt float64, rect image.Rectangle) bool {
	box, ok := opBounds(op, ppt)
	if !ok {
		return true
	}

	return !box.Intersect(rect).Empty()
}
