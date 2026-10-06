package ui

// applyHistory accepts the page when its generation is current, walks back
// when a later page shrank away, and relayouts on any accepted page.
func (a *App) applyHistory(page *PageResponse) {
	if page == nil {
		return
	}

	ok, retry := a.pager.Accept(*page)
	if retry {
		a.askPage()

		return
	}

	if ok {
		a.syncPage()
		a.geom = true
	}
}

// syncPage copies the pager state into the printable view.
func (a *App) syncPage() {
	a.view.Filter = a.pager.Filter()
	a.view.History = a.pager.Rows()
	a.view.Page = a.pager.PageNumber()
	a.view.Pages = a.pager.Pages()
	a.view.Total = a.pager.Total()
	a.view.HasPrev = a.pager.HasPrev()
	a.view.HasNext = a.pager.HasNext()
	a.view.Loading = a.pager.Loading()
	a.view.Detail = a.findSelected()
}

// findSelected returns the selected job as a Row copy, or nil.
func (a *App) findSelected() *Row {
	id := a.pager.Selected()
	if id == "" {
		return nil
	}

	for _, row := range a.view.Active {
		if row.ID == id {
			selected := row

			return &selected
		}
	}

	for _, row := range a.pager.Rows() {
		if row.ID == id {
			selected := row

			return &selected
		}
	}

	return nil
}
