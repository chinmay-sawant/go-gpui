package page

import "context"

// Type calls the type handler and draws the page again.
// It does nothing when no type handler is registered.
func (p *Page) Type(ctx context.Context, text string) error {
	return p.after(ctx, p.handlers.Type == nil, func() error {
		return p.handlers.Type(ctx, text)
	})
}

// Backspace calls the backspace handler and draws the page again.
// It does nothing when no backspace handler is registered.
func (p *Page) Backspace(ctx context.Context) error {
	return p.after(ctx, p.handlers.Backspace == nil, func() error {
		return p.handlers.Backspace(ctx)
	})
}

// Submit calls the submit handler and draws the page again.
// It does nothing when no submit handler is registered.
func (p *Page) Submit(ctx context.Context) error {
	return p.after(ctx, p.handlers.Submit == nil, func() error {
		return p.handlers.Submit(ctx)
	})
}

// DeleteWord calls the delete-word handler and draws the page again.
// It does nothing when no delete-word handler is registered.
func (p *Page) DeleteWord(ctx context.Context) error {
	return p.after(ctx, p.handlers.DeleteWord == nil, func() error {
		return p.handlers.DeleteWord(ctx)
	})
}

func (p *Page) after(ctx context.Context, skip bool, fn func() error) error {
	if err := useContext(ctx); err != nil {
		return err
	}

	if skip {
		return nil
	}

	if err := fn(); err != nil {
		return err
	}

	return p.Redraw(ctx)
}
