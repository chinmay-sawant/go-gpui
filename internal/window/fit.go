package window

import "math"

const scrollStep = 48

// framePoint maps a pointer on the current screen back onto a stretched picture.
// The picture fills the screen, so a point halfway across the window is halfway
// across the picture.
func framePoint(x, y, frameW, frameH, screenW, screenH int) (float64, float64) {
	if frameW <= 0 || frameH <= 0 || screenW <= 0 || screenH <= 0 {
		return 0, 0
	}

	return float64(x) * float64(frameW) / float64(screenW),
		float64(y) * float64(frameH) / float64(screenH)
}

// contentPoint maps a cursor into the painted page.
// When the page is taller than the window, scrollX and scrollY are the
// page position at the window's top left. A stretched picture uses framePoint.
func contentPoint(x, y, scrollX, scrollY int, stretched bool, frameW, frameH, screenW, screenH int) (float64, float64) {
	if stretched {
		return framePoint(x, y, frameW, frameH, screenW, screenH)
	}

	return float64(x + scrollX), float64(y + scrollY)
}

// panScroll moves the window through a page that is larger than the view.
// A positive wheel moves toward the start of that axis.
func panScroll(scrollX, scrollY int, wheelX, wheelY float64, contentW, contentH, viewW, viewH int) (int, int) {
	nextX := int(math.Round(float64(scrollX) - wheelX*scrollStep))
	nextY := int(math.Round(float64(scrollY) - wheelY*scrollStep))

	return clampScroll(nextX, nextY, contentW, contentH, viewW, viewH)
}

func clampScroll(scrollX, scrollY, contentW, contentH, viewW, viewH int) (int, int) {
	return clampAxis(scrollX, contentW-viewW), clampAxis(scrollY, contentH-viewH)
}

func clampAxis(value, limit int) int {
	if limit < 0 {
		limit = 0
	}

	if value < 0 {
		return 0
	}

	if value > limit {
		return limit
	}

	return value
}
