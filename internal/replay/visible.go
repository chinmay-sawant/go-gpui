package replay

import (
	"image"
	"math"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
	"github.com/hajimehoshi/ebiten/v2"
)

// DrawVisible skips offscreen operations before text shaping or GPU submission.
// Unbounded ink is retained. Translation maps the target back into page space.
func DrawVisible(dst *ebiten.Image, display *layout.Display, dx, dy float64) {
	if dst == nil || display == nil {
		return
	}
	rect := visibleRect(dst.Bounds(), dx, dy)
	visitVisible(display, rect, func(index int) { drawOp(dst, &display.Ops[index], dx, dy) })
}

func visibleRect(bounds image.Rectangle, dx, dy float64) image.Rectangle {
	return image.Rect(
		int(math.Floor(float64(bounds.Min.X)-dx)),
		int(math.Floor(float64(bounds.Min.Y)-dy)),
		int(math.Ceil(float64(bounds.Max.X)-dx)),
		int(math.Ceil(float64(bounds.Max.Y)-dy)),
	)
}

func visitVisible(display *layout.Display, rect image.Rectangle, draw func(int)) {
	for _, index := range display.Order {
		if index < 0 || index >= len(display.Ops) {
			continue
		}
		if opTouches(&display.Ops[index], display.PixelPerPoint, rect) {
			draw(index)
		}
	}
}
