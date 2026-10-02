package page

func (p *Page) syncForm(body string) string {
	if p.form == nil || p.form.doc != p.pastAt {
		p.form = &formState{
			byID: map[string]Control{},
			doc:  p.pastAt,
		}
	}

	body = rewriteButtons(body)
	spans := scanControls(body)
	byID, order := mergeControls(p.form.byID, spans)
	for id, ctrl := range byID {
		if ctrl.Bind == "" {
			continue
		}

		if c, ok := bindRead(p, ctrl); ok {
			byID[id] = c
		}
	}

	p.form.byID = byID
	p.form.order = order
	if _, ok := byID[p.form.focusID]; !ok {
		p.form.focusID = ""
		p.form.selected = false
	}

	return rewriteControls(body, spans, byID, p.form.focusID, p.form.selected)
}
