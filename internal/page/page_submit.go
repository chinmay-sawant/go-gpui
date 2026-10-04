package page

import "context"

// Submit calls the submit handler and draws the page again.
// It does nothing when no submit handler is registered.
func (p *Page) Submit(ctx context.Context) error {
	return p.after(ctx, p.handlers.Submit == nil, func() error {
		return p.handlers.Submit(ctx)
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
