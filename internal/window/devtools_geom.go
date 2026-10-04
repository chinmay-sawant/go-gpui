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

// devScreenRect converts a rectangle in page space to screen pixels. The
// page draws through the pinch zoom, so a page point lands at
// point*zoom - scroll. Stretched: scale by screenW/frameW and screenH/frameH
// as well, the factors drawReplayScaled uses.
func devScreenRect(r devRect, scrollX, scrollY int, zoom float64, stretched bool, frameW, frameH, screenW, screenH int) devRect {
	if zoom <= 0 {
		zoom = 1
	}

	if !stretched {
		return devRect{
			X: r.X*zoom - float64(scrollX),
			Y: r.Y*zoom - float64(scrollY),
			W: r.W * zoom,
			H: r.H * zoom,
		}
	}

	sx, sy := 1.0, 1.0

	if frameW > 0 && screenW > 0 {
		sx = float64(screenW) / float64(frameW)
	}

	if frameH > 0 && screenH > 0 {
		sy = float64(screenH) / float64(frameH)
	}

	return devRect{
		X: r.X * sx * zoom,
		Y: r.Y * sy * zoom,
		W: r.W * sx * zoom,
		H: r.H * sy * zoom,
	}
}

// devScreen converts a page-space rectangle for this frame.
func (s *shell) devScreen(r devRect) devRect {
	frameW, frameH := s.frameSize()

	return devScreenRect(r, s.scrollX, s.scrollY, s.zoom(), s.stretched(), frameW, frameH, s.screenW, s.screenH)
}
