package page

import "context"

// Copy returns the text the page wants on the clipboard.
// A focused text field returns its value and does not call the handler.
// ok is false when there is nothing to copy. Copy does not draw.
func (p *Page) Copy(ctx context.Context) (string, bool, error) {
	if err := useContext(ctx); err != nil {
		return "", false, err
	}

	if c, ok := p.focusedEditable(); ok {
		return c.Value, true, nil
	}

	if p.handlers.Copy == nil {
		return "", false, nil
	}

	return p.handlers.Copy(ctx)
}

// Cut returns the text the page wants on the clipboard and draws again.
// A focused text field returns its value, clears it, writes the bound field,
// and fires Change before the redraw. It does not call the handler.
// ok is false when there is nothing to cut.
func (p *Page) Cut(ctx context.Context) (string, bool, error) {
	if err := useContext(ctx); err != nil {
		return "", false, err
	}

	if c, ok := p.focusedEditable(); ok {
		id := p.form.focusID
		if err := p.beforeEdit(ctx, id); err != nil {
			return "", false, err
		}

		text := c.Value
		c.Value = ""
		p.form.byID[id] = c
		p.form.selected = false
		bindWrite(p, c)
		if err := p.change(ctx, id); err != nil {
			return text, true, err
		}
		err := p.Redraw(ctx)

		return text, true, err
	}

	if p.handlers.Cut == nil {
		return "", false, nil
	}

	text, ok, err := p.handlers.Cut(ctx)
	if err != nil || !ok {
		return text, ok, err
	}

	if err := p.Redraw(ctx); err != nil {
		return text, ok, err
	}

	return text, ok, nil
}

func useContext(ctx context.Context) error {
	if ctx == nil {
		return errNilContext
	}

	return ctx.Err()
}
