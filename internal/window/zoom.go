package window

// zoom returns the pinch scale. A shell that never pinched reads 1.
func (s *shell) zoom() float64 {
	return s.fingers.zoomOr1()
}

// contentAt maps a window point into the page for the current scroll and
// pinch zoom.
func (s *shell) contentAt(x, y, frameW, frameH int) (float64, float64) {
	return contentPointZoom(
		x, y, s.scrollX, s.scrollY, s.zoom(), s.stretched(),
		frameW, frameH, s.screenW, s.screenH,
	)
}

// contentPointZoom maps a cursor into the painted page and divides out the
// pinch zoom, so a click stays on the glyph under the finger.
func contentPointZoom(x, y, scrollX, scrollY int, zoom float64, stretched bool, frameW, frameH, screenW, screenH int) (float64, float64) {
	px, py := contentPoint(x, y, scrollX, scrollY, stretched, frameW, frameH, screenW, screenH)
	if zoom <= 0 {
		zoom = 1
	}

	return px / zoom, py / zoom
}
