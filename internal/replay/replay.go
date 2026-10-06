// Package replay draws a layout display list onto an Ebiten image. The window
// uses it to show the page as vector operations instead of one rasterized
// bitmap. The caller decides replayability first; a page it cannot draw keeps
// the bitmap path.
package replay

import (
	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
	"github.com/hajimehoshi/ebiten/v2"
)

// Draw replays display in paint order. dx and dy move the page's top-left in
// pixels, so the window passes its scroll offset. A viewport-fixed op ignores
// the offset: the engine placed it against the viewport.
func Draw(dst *ebiten.Image, display *layout.Display, dx, dy float64) {
	if display == nil {
		return
	}

	for _, index := range display.Order {
		if index < 0 || index >= len(display.Ops) {
			continue
		}

		drawOp(dst, &display.Ops[index], dx, dy)
	}
}

// DrawUnfixed replays only the operations that scroll with the page. A
// caller that paints a scrolled buffer draws the fixed ops separately, after
// the buffer, so they stay at the viewport.
func DrawUnfixed(dst *ebiten.Image, display *layout.Display, dx, dy float64) {
	if display == nil {
		return
	}

	for _, index := range display.Order {
		if index < 0 || index >= len(display.Ops) {
			continue
		}

		op := &display.Ops[index]
		if !op.Fixed {
			drawOp(dst, op, dx, dy)
		}
	}
}

// DrawFixed replays only the viewport-fixed operations, untranslated.
func DrawFixed(dst *ebiten.Image, display *layout.Display) {
	if display == nil {
		return
	}

	for _, index := range display.Order {
		if index < 0 || index >= len(display.Ops) {
			continue
		}

		op := &display.Ops[index]
		if op.Fixed {
			drawOp(dst, op, 0, 0)
		}
	}
}
