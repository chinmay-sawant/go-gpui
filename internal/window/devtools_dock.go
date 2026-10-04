package window

const (
	// devDockWidth is the default width of the right-hand panel.
	devDockWidth = 340.0

	// devDockMin is the narrowest the panel may become while resizing.
	devDockMin = 220.0

	// devDockGrab is the half-width of the resize edge beside the panel.
	devDockGrab = 4.0
)

// devDockRect places the devtools panel on the right edge, full height. A
// width larger than the window is clamped, so a phone shows the panel alone.
func devDockRect(screenW, screenH int, width float64) devRect {
	if width < 0 {
		width = 0
	}

	if width > float64(screenW) {
		width = float64(screenW)
	}

	return devRect{X: float64(screenW) - width, Y: 0, W: width, H: float64(screenH)}
}
