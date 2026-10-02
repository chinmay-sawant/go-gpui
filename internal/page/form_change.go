package page

import "context"

// change calls the Change handler for id.
// A nil handler does nothing. The handler gets the last box in p.boxes with
// that id, or Box{ID: id} when the page has no such box.
func (p *Page) change(ctx context.Context, id string) error {
	if p.handlers.Change == nil {
		return nil
	}

	box := Box{ID: id}
	for _, b := range p.boxes {
		if b.ID == id {
			box = b
		}
	}

	return p.handlers.Change(ctx, box)
}
