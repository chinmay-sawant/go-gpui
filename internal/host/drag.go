package host

import "context"

// PointerDragger can claim a gesture before the window scrolls or clicks.
// Coordinates are CSS pixels. EndDrag also ends a cancelled gesture.
type PointerDragger interface {
	BeginDrag(ctx context.Context, x, y float64) (bool, error)
	MoveDrag(ctx context.Context, x, y float64) error
	EndDrag(ctx context.Context) error
}
