package host

import "context"

// LongPresser is a screen that handles a press held in place, such as a
// long press on a touch screen. The window fires LongPress once after a
// press stays within a small slop for the hold delay. A true answer claims
// the press: the rest of the gesture, including the lift, belongs to the
// screen instead of a page scroll or a tap.
type LongPresser interface {
	LongPress(ctx context.Context, x, y float64) (bool, error)
}
