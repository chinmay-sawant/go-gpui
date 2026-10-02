package page

import "context"

// boxByID returns the last box in p.boxes with that id, or Box{ID: id}.
func (p *Page) boxByID(id string) Box {
	box := Box{ID: id}
	for _, b := range p.boxes {
		if b.ID == id {
			box = b
		}
	}

	return box
}

// beforeEdit calls the BeforeEdit handler for id.
// A nil handler does nothing.
func (p *Page) beforeEdit(ctx context.Context, id string) error {
	if p.handlers.BeforeEdit == nil {
		return nil
	}

	return p.handlers.BeforeEdit(ctx, p.boxByID(id))
}

// change calls the Change handler for id.
// A nil handler does nothing.
func (p *Page) change(ctx context.Context, id string) error {
	if p.handlers.Change == nil {
		return nil
	}

	return p.handlers.Change(ctx, p.boxByID(id))
}
