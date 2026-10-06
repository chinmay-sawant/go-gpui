package window

import "github.com/hajimehoshi/ebiten/v2"

// swipeMin is how far a one-finger drag must travel to count as a swipe.
const swipeMin = 24

// swipeDelta is a one-finger swipe that lifted past the tap slop.
type swipeDelta struct {
	dx, dy int
}

// liftSwipe returns the swipe a lifting finger made, or nil. A finger that
// never moved is a tap, an eaten finger belongs to the screen, and a
// gesture with a second finger never swipes.
func liftSwipe(f touchFinger, multi bool) *swipeDelta {
	if !f.moved || f.eaten || multi {
		return nil
	}

	dx := f.x - f.startX
	dy := f.y - f.startY
	if absInt(dx) < swipeMin && absInt(dy) < swipeMin {
		return nil
	}

	return &swipeDelta{dx: dx, dy: dy}
}

// touchPos is one finger position this frame.
type touchPos struct {
	id   ebiten.TouchID
	x, y int
}

// touchFinger is one finger in flight.
type touchFinger struct {
	id     ebiten.TouchID
	x, y   int
	startX int
	startY int
	moved  bool
	eaten  bool
}
