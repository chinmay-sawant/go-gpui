package page

import "context"

// KeyDown calls the key-down handler. When a text field is focused, Escape
// clears the focus and the caret keys move the caret; both draw. A nil
// handler does nothing else.
func (p *Page) KeyDown(ctx context.Context, key string) error {
	if err := useContext(ctx); err != nil {
		return err
	}

	if p.handlers.KeyDown != nil {
		if err := p.handlers.KeyDown(ctx, key); err != nil {
			return err
		}
	}

	if key == "escape" {
		return p.Focus(ctx, "")
	}

	return p.caretKey(ctx, key)
}

// KeyUp calls the key-up handler. It does not draw. A nil handler does
// nothing.
func (p *Page) KeyUp(ctx context.Context, key string) error {
	if err := useContext(ctx); err != nil {
		return err
	}

	if p.handlers.KeyUp == nil {
		return nil
	}

	return p.handlers.KeyUp(ctx, key)
}
