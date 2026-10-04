package page

import "context"

func (p *Page) editField(ctx context.Context, fn func() error, edit func(Control) (Control, bool)) error {
	if err := useContext(ctx); err != nil {
		return err
	}

	ran := false
	if fn != nil {
		if err := fn(); err != nil {
			return err
		}

		ran = true
	}

	changed, err := p.applyEdit(ctx, edit)
	if err != nil {
		return err
	}

	if !changed && !ran {
		return nil
	}

	if changed {
		p.markPending(p.form.focusID)

		if err := p.change(ctx, p.form.focusID); err != nil {
			return err
		}
	}

	return p.Redraw(ctx)
}

func (p *Page) applyEdit(ctx context.Context, edit func(Control) (Control, bool)) (bool, error) {
	c, ok := p.typingTarget()
	if !ok {
		return false, nil
	}

	next, changed := edit(c)
	if !changed {
		return false, nil
	}

	if err := p.beforeEdit(ctx, p.form.focusID); err != nil {
		return false, err
	}

	p.form.byID[p.form.focusID] = next
	p.form.selected = false
	bindWrite(p, next)

	return true, nil
}

func (p *Page) typingTarget() (Control, bool) {
	c, ok := p.focusedEditable()
	if !ok || c.Disabled {
		return Control{}, false
	}

	return c, true
}

func (p *Page) focusedEditable() (Control, bool) {
	if p.form == nil {
		return Control{}, false
	}

	c, ok := p.control(p.form.focusID)
	if !ok || !canEdit(c) {
		return Control{}, false
	}

	return c, true
}

func (p *Page) insertValue(c Control, text string) (Control, bool) {
	if text == "" {
		return c, false
	}

	if p.form.selected {
		c.Value = text

		return c, true
	}

	c.Value += text

	return c, true
}

func (p *Page) backspaceValue(c Control) (Control, bool) {
	if p.form.selected {
		c.Value = ""

		return c, true
	}

	if c.Value == "" {
		return c, false
	}

	c.Value = dropLastRune(c.Value)

	return c, true
}

func (p *Page) deleteWordValue(c Control) (Control, bool) {
	if p.form.selected {
		c.Value = ""

		return c, true
	}

	next := dropLastWord(c.Value)
	if next == c.Value {
		return c, false
	}

	c.Value = next

	return c, true
}
