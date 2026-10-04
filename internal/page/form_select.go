package page

import "context"

// SelectAt focuses the control under the point and puts the caret at the
// measured rune offset. A point outside every control blurs the form. A
// disabled control is left alone. A control that takes no text just takes
// focus.
func (p *Page) SelectAt(ctx context.Context, x, y float64) error {
	if err := useContext(ctx); err != nil {
		return err
	}

	id := p.boxIDAt(x, y)
	c, ok := p.control(id)
	if !ok {
		if p.form == nil || p.form.focusID == "" {
			return nil
		}

		p.blurForm()

		return p.Redraw(ctx)
	}

	if c.Disabled {
		return nil
	}

	if !canEdit(c) {
		return p.Focus(ctx, id)
	}

	return p.focusCaret(ctx, c, p.offsetAt(p.boxByID(id), c, x, y))
}

// Drag extends the focused field's range to the point. It does nothing when
// no text field is focused.
func (p *Page) Drag(ctx context.Context, x, y float64) error {
	if err := useContext(ctx); err != nil {
		return err
	}

	c, ok := p.typingTarget()
	if !ok {
		return nil
	}

	off := p.offsetAt(p.boxByID(c.ID), c, x, y)
	if off == p.form.caret {
		return nil
	}

	p.form.caret = off
	p.form.all = false

	return p.Redraw(ctx)
}

// SelectWordAt selects the word under the point and draws.
func (p *Page) SelectWordAt(ctx context.Context, x, y float64) error {
	return p.selectSpan(ctx, x, y, wordBounds)
}

// SelectLineAt selects the line under the point and draws.
func (p *Page) SelectLineAt(ctx context.Context, x, y float64) error {
	return p.selectSpan(ctx, x, y, lineBounds)
}

func (p *Page) selectSpan(ctx context.Context, x, y float64, bounds func(string, int) (int, int)) error {
	if err := useContext(ctx); err != nil {
		return err
	}

	id := p.boxIDAt(x, y)
	c, ok := p.control(id)
	if !ok || c.Disabled || !canEdit(c) {
		return nil
	}

	off := p.offsetAt(p.boxByID(id), c, x, y)
	start, end := bounds(c.Value, off)
	p.form.focusID = c.ID
	p.form.anchor, p.form.caret = start, end
	p.form.all = false

	return p.Redraw(ctx)
}
