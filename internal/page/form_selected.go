package page

// FormSelected reports whether id is the focused control and its whole
// value is selected. It is false when no control is focused, when the
// page has no form yet, or when another control is focused.
func (p *Page) FormSelected(id string) bool {
	if p.form == nil {
		return false
	}

	return p.form.selected && p.form.focusID == id
}
