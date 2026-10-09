package page

import "context"

type pageDrag struct {
	active bool
	x, y   float64
}

// BeginDrag offers the box under the point to Handlers.DragStart.
func (p *Page) BeginDrag(ctx context.Context, x, y float64) (bool, error) {
	if err := useContext(ctx); err != nil {
		return false, err
	}
	if p.handlers.DragStart == nil {
		return false, nil
	}
	box := p.boxByID(p.boxIDAt(x, y))
	claimed, err := p.handlers.DragStart(ctx, box)
	if err == nil && claimed {
		p.gesture = pageDrag{active: true, x: x, y: y}
		err = p.Press(ctx, x, y)
	}
	return claimed, err
}

// MoveDrag sends movement while a claimed gesture is active.
func (p *Page) MoveDrag(ctx context.Context, x, y float64) error {
	if err := useContext(ctx); err != nil {
		return err
	}
	if !p.gesture.active {
		return nil
	}
	dx, dy := x-p.gesture.x, y-p.gesture.y
	p.gesture.x, p.gesture.y = x, y
	if p.handlers.DragMove == nil || (dx == 0 && dy == 0) {
		return nil
	}
	return p.handlers.DragMove(ctx, dx, dy)
}

// EndDrag releases the claimed gesture without synthesizing a click.
func (p *Page) EndDrag(ctx context.Context) error {
	if err := useContext(ctx); err != nil {
		return err
	}
	active := p.gesture.active
	p.gesture = pageDrag{}
	if !active {
		return nil
	}
	if p.handlers.DragEnd != nil {
		if err := p.handlers.DragEnd(ctx); err != nil {
			return err
		}
	}
	return p.Release(ctx)
}
