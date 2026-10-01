package page

import "context"

// Paste calls the paste handler and draws the page again.
// It does nothing when no paste handler is registered.
func (p *Page) Paste(ctx context.Context, text string) error {
	return p.after(ctx, p.handlers.Paste == nil, func() error {
		return p.handlers.Paste(ctx, text)
	})
}

// SelectAll calls the select-all handler and draws the page again.
// It does nothing when no select-all handler is registered.
func (p *Page) SelectAll(ctx context.Context) error {
	return p.after(ctx, p.handlers.SelectAll == nil, func() error {
		return p.handlers.SelectAll(ctx)
	})
}

// Undo calls the undo handler and draws the page again.
// It does nothing when no undo handler is registered.
func (p *Page) Undo(ctx context.Context) error {
	return p.after(ctx, p.handlers.Undo == nil, func() error {
		return p.handlers.Undo(ctx)
	})
}

// Redo calls the redo handler and draws the page again.
// It does nothing when no redo handler is registered.
func (p *Page) Redo(ctx context.Context) error {
	return p.after(ctx, p.handlers.Redo == nil, func() error {
		return p.handlers.Redo(ctx)
	})
}
