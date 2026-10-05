package page

import "context"

// Type calls the type handler and edits a focused text field.
// It draws when the handler ran or the field changed.
// A nil handler with nothing to edit does nothing.
// IME: Ebiten v2.10.4 ships experimental exp/textinput, but the window
// does not wire it yet and the page holds no composing run, so no composing
// text reaches this path. The design that would land here is in plans/v0.0.2/input-interaction.md.
func (p *Page) Type(ctx context.Context, text string) error {
	var fn func() error
	if p.handlers.Type != nil {
		fn = func() error {
			return p.handlers.Type(ctx, text)
		}
	}

	return p.editField(ctx, fn, func(c Control, caret, anchor int) (Control, int, int, bool) {
		return p.insertValue(c, caret, anchor, text)
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

	return p.editField(ctx, fn, func(c Control, caret, anchor int) (Control, int, int, bool) {
		return p.backspaceValue(c, caret, anchor)
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

	return p.editField(ctx, fn, func(c Control, caret, anchor int) (Control, int, int, bool) {
		return p.deleteWordValue(c, caret, anchor)
	})
}
