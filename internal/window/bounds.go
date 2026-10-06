package window

import "github.com/chinmay-sawant/ownframe/internal/host"

// maxSizer is a screen that asks for a largest window size.
type maxSizer interface {
	MaxSize() (int, int)
}

// windowBounds returns the largest window the screen asks for, or -1 for no
// bound. A screen without MaxSize is unbounded, as it was before the cap
// became a window bound.
func windowBounds(app host.Screen) (int, int) {
	maxW, maxH := -1, -1

	sizer, ok := app.(maxSizer)
	if !ok {
		return maxW, maxH
	}

	maxW, maxH = sizer.MaxSize()
	if maxW <= 0 {
		maxW = -1
	}

	if maxH <= 0 {
		maxH = -1
	}

	return maxW, maxH
}
