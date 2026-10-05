package page

import "context"

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
