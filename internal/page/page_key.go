package page

import "context"

// KeyDown calls the key-down handler. It does not draw. A nil handler
// does nothing.
func (p *Page) KeyDown(ctx context.Context, key string) error {
	if err := useContext(ctx); err != nil {
		return err
	}

	if p.handlers.KeyDown == nil {
		return nil
	}

	return p.handlers.KeyDown(ctx, key)
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
