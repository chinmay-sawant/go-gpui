package ui

// selectRow stores the selected job ID and a snapshot for the details pane.
func (a *App) selectRow(id string) {
	a.pager.Select(id)
	a.detail = a.findRow(id)
	a.view.Detail = a.detail
}

// clearDetail closes the details pane and clears the selection.
func (a *App) clearDetail() {
	a.pager.Select("")
	a.detail = nil
	a.view.Detail = nil
}

// updateDetail refreshes the snapshot when it is for the selected job.
func (a *App) updateDetail(row Row) {
	if a.detail == nil || a.detail.ID != row.ID {
		return
	}

	snapshot := row
	a.detail = &snapshot
	a.view.Detail = a.detail
}

// refreshDetail updates the snapshot from the current rows when the job is
// still visible, and leaves it alone when the job is on another page.
func (a *App) refreshDetail() {
	if a.detail == nil {
		return
	}

	if row := a.findRow(a.detail.ID); row != nil {
		a.detail = row
	}

	a.view.Detail = a.detail
}

// findRow returns a copy of one row from the active set or the current
// history page, or nil.
func (a *App) findRow(id string) *Row {
	if id == "" {
		return nil
	}

	for _, row := range a.view.Active {
		if row.ID == id {
			snapshot := row

			return &snapshot
		}
	}

	for _, row := range a.pager.Rows() {
		if row.ID == id {
			snapshot := row

			return &snapshot
		}
	}

	return nil
}
