package window

// devRect is a rectangle in screen pixels.
type devRect struct {
	X, Y, W, H float64
}

// devHit is one clickable row of the panel.
type devHit struct {
	rect devRect
	kind int
}

const devHitOps = iota

// devScreenRect converts a rectangle in page space to screen pixels. Not
// stretched: subtract scroll. Stretched: scale by screenW/frameW and
// screenH/frameH, the factors drawReplayScaled uses.
func devScreenRect(r devRect, scrollX, scrollY int, stretched bool, frameW, frameH, screenW, screenH int) devRect {
	if !stretched {
		return devRect{X: r.X - float64(scrollX), Y: r.Y - float64(scrollY), W: r.W, H: r.H}
	}

	sx, sy := 1.0, 1.0

	if frameW > 0 && screenW > 0 {
		sx = float64(screenW) / float64(frameW)
	}

	if frameH > 0 && screenH > 0 {
		sy = float64(screenH) / float64(frameH)
	}

	return devRect{X: r.X * sx, Y: r.Y * sy, W: r.W * sx, H: r.H * sy}
}

// devScreen converts a page-space rectangle for this frame.
func (s *shell) devScreen(r devRect) devRect {
	frameW, frameH := s.frameSize()

	return devScreenRect(r, s.scrollX, s.scrollY, s.stretched(), frameW, frameH, s.screenW, s.screenH)
}

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
