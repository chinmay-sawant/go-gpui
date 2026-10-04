package page

import "context"

// valueEdit computes one edit to a field: the new value, the caret and the
// anchor to store, and whether anything changed.
type valueEdit func(c Control, caret, anchor int) (Control, int, int, bool)

func (p *Page) editField(ctx context.Context, fn func() error, edit valueEdit) error {
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
		if err := p.change(ctx, p.form.focusID); err != nil {
			return err
		}
	}

	return p.Redraw(ctx)
}

func (p *Page) applyEdit(ctx context.Context, edit valueEdit) (bool, error) {
	c, ok := p.typingTarget()
	if !ok {
		return false, nil
	}

	next, caret, anchor, changed := edit(c, p.form.caret, p.form.anchor)
	if !changed {
		return false, nil
	}

	if err := p.beforeEdit(ctx, p.form.focusID); err != nil {
		return false, err
	}

	p.form.byID[p.form.focusID] = next
	p.form.caret, p.form.anchor = caret, anchor
	p.form.all = false
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
