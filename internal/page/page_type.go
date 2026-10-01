package page

import "context"

// Type calls the type handler and edits a focused text field.
// It draws when the handler ran or the field changed.
// A nil handler with nothing to edit does nothing.
func (p *Page) Type(ctx context.Context, text string) error {
	var fn func() error
	if p.handlers.Type != nil {
		fn = func() error {
			return p.handlers.Type(ctx, text)
		}
	}

	return p.editField(ctx, fn, func(c Control) (Control, bool) {
		return p.insertValue(c, text)
	})
}

// Backspace calls the backspace handler and edits a focused text field.
// It draws when the handler ran or the field changed.
// A nil handler with nothing to edit does nothing.
func (p *Page) Backspace(ctx context.Context) error {
	var fn func() error
	if p.handlers.Backspace != nil {
		fn = func() error {
			return p.handlers.Backspace(ctx)
		}
	}

	return p.editField(ctx, fn, func(c Control) (Control, bool) {
		return p.backspaceValue(c)
	})
}

// Submit calls the submit handler and draws the page again.
// It does nothing when no submit handler is registered.
func (p *Page) Submit(ctx context.Context) error {
	return p.after(ctx, p.handlers.Submit == nil, func() error {
		return p.handlers.Submit(ctx)
	})
}

// DeleteWord calls the delete-word handler and edits a focused text field.
// It draws when the handler ran or the field changed.
// A nil handler with nothing to edit does nothing.
func (p *Page) DeleteWord(ctx context.Context) error {
	var fn func() error
	if p.handlers.DeleteWord != nil {
		fn = func() error {
			return p.handlers.DeleteWord(ctx)
		}
	}

	return p.editField(ctx, fn, func(c Control) (Control, bool) {
		return p.deleteWordValue(c)
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
