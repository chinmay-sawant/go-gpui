package page

import "context"

// Swipe reports a one-finger swipe after the finger lifts. A page without
// a Swipe handler ignores the gesture. It does not draw: change state and
// let a SetTick callback paint, or call Redraw.
func (p *Page) Swipe(ctx context.Context, dx, dy float64) error {
	if err := useContext(ctx); err != nil {
		return err
	}

	if p.handlers.Swipe == nil {
		return nil
	}

	return p.handlers.Swipe(ctx, dx, dy)
}
