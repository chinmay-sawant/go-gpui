package host

import "context"

// Swiper is a screen that handles a one-finger swipe: a touch that moved
// past the tap slop and lifted without a tap or a claimed long press.
// dx and dy are the movement in window pixels; dy is positive downward.
type Swiper interface {
	Swipe(ctx context.Context, dx, dy float64) error
}
