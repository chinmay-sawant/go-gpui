package window

// devPanelRect places a w x h panel at the bottom left, above the horizontal
// scrollbar strip, without covering the fallback badge in the top right.
func devPanelRect(screenW, screenH int, w, h float64, hScroll bool) devRect {
	strip := 0.0
	if hScroll {
		strip = scrollbarThickness
	}

	bottom := float64(screenH) - devPanelGap - strip
	x, y := float64(devPanelGap), bottom-h

	bx, by, _, bh := badgeRect(float64(screenW), badgeLabel)
	if x+w > bx-4 && y < by+bh+4 {
		y = by + bh + 4
	}

	if y+h > bottom {
		h = bottom - y
	}

	if h < 0 {
		h = 0
	}

	return devRect{X: x, Y: y, W: w, H: h}
}

// devInRect reports whether a point is inside a rectangle.
func devInRect(r devRect, x, y float64) bool {
	return x >= r.X && x <= r.X+r.W && y >= r.Y && y <= r.Y+r.H
}
