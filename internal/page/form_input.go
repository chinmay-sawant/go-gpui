package page

import "context"

func (p *Page) clickControl(ctx context.Context, box Box) error {
	c, _ := p.control(box.ID)
	if !c.Disabled {
		p.activate(box.ID)
	}

	if p.handlers.Click != nil {
		if err := p.handlers.Click(ctx, box); err != nil {
			return err
		}
	}

	return p.Redraw(ctx)
}

func (p *Page) activate(id string) {
	c, ok := p.control(id)
	if !ok || c.Disabled {
		return
	}

	switch {
	case c.Type == "checkbox":
		c.Checked = !c.Checked
	case c.Type == "radio":
		c.Checked = true
		p.uncheckRadios(c.Name, id)
	case c.Tag == "select":
		c = nextOption(c)
	}

	p.form.focusID = id
	p.form.selected = false
	p.form.byID[id] = c
}

func (p *Page) uncheckRadios(name, keep string) {
	if name == "" || p.form == nil {
		return
	}

	for id, other := range p.form.byID {
		if id == keep || other.Type != "radio" || other.Name != name {
			continue
		}

		other.Checked = false
		p.form.byID[id] = other
	}
}

func nextOption(c Control) Control {
	n := len(c.Options)
	if n == 0 {
		return c
	}

	next := 0
	for i, opt := range c.Options {
		if opt.Selected {
			next = i + 1
			break
		}
	}

	if next == n {
		next = 0
	}

	for i := range c.Options {
		c.Options[i].Selected = i == next
	}

	c.Value = c.Options[next].Value

	return c
}

func (p *Page) blurForm() {
	if p.form == nil {
		return
	}

	p.form.focusID = ""
	p.form.selected = false
}

func (p *Page) formControl(id string) bool {
	if id == "" {
		return false
	}

	_, ok := p.control(id)

	return ok
}

func canEdit(c Control) bool {
	if c.Tag == "select" || c.Type == "checkbox" || c.Type == "radio" {
		return false
	}

	return c.Tag == "textarea" || textLike(c.Type)
}
