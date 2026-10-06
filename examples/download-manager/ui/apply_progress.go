package ui

// applyProgress replaces one active row. A state change relayouts; a byte
// count or speed change only repaints.
func (a *App) applyProgress(r *Row) {
	if r == nil {
		return
	}

	for i := range a.view.Active {
		if a.view.Active[i].ID != r.ID {
			continue
		}

		changed := a.view.Active[i].State != r.State
		a.view.Active[i] = *r
		if changed {
			a.geom = true
			a.historyDirty = true
		}

		if !r.State.Active() {
			// A terminal job leaves the active set the next time the
			// backend sends one.
			a.needActive = true
		}

		if a.view.Detail != nil && a.view.Detail.ID == r.ID {
			detail := *r
			a.view.Detail = &detail
		}

		return
	}

	// A job the active set does not know about yet: ask for a fresh set.
	a.needActive = true
}
