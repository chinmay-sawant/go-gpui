package window

// framePoint maps a pointer on the current screen back onto the PNG.
// While a resize is still settling, the screen size and the PNG size differ,
// and the picture is stretched to fill the screen.
func framePoint(x, y, frameW, frameH, screenW, screenH int) (float64, float64) {
	if frameW <= 0 || frameH <= 0 || screenW <= 0 || screenH <= 0 {
		return 0, 0
	}

	return float64(x) * float64(frameW) / float64(screenW),
		float64(y) * float64(frameH) / float64(screenH)
}
