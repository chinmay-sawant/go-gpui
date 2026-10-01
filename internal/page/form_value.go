package page

// FormValue returns the current value of id, or "" when id is absent.
func (p *Page) FormValue(id string) string {
	c, _ := p.control(id)

	return c.Value
}

// SetFormValue stores value for id and does not draw.
// It does nothing when that control is not on the page yet.
// A select marks each option whose value matches.
// The focused control loses its selection.
func (p *Page) SetFormValue(id, value string) {
	c, ok := p.control(id)
	if !ok {
		return
	}

	c.Value = value
	if c.Tag == "select" {
		for i := range c.Options {
			c.Options[i].Selected = c.Options[i].Value == value
		}
	}

	p.form.byID[id] = c
	if p.form.focusID == id {
		p.form.selected = false
	}
}

// FormChecked reports whether id is checked.
// A missing control returns false.
func (p *Page) FormChecked(id string) bool {
	c, _ := p.control(id)

	return c.Checked
}

// SetFormChecked sets a checkbox or radio and does not draw.
// It does nothing when id is missing.
// A radio set true clears other radios with the same non-empty name.
// A radio set false clears only that radio.
func (p *Page) SetFormChecked(id string, checked bool) {
	c, ok := p.control(id)
	if !ok {
		return
	}

	switch c.Type {
	case "checkbox":
		c.Checked = checked
		p.form.byID[id] = c
	case "radio":
		c.Checked = checked
		p.form.byID[id] = c
		if checked {
			p.uncheckRadios(c.Name, id)
		}
	}
}

// FocusedField returns the focused control id, or "" when none is focused.
func (p *Page) FocusedField() string {
	if p.form == nil {
		return ""
	}

	return p.form.focusID
}

func (p *Page) control(id string) (Control, bool) {
	if p.form == nil {
		return Control{}, false
	}

	c, ok := p.form.byID[id]

	return c, ok
}
