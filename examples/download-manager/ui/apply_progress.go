package ui

// applyProgress replaces one active row. A state change relayouts; a byte
// count or speed change only repaints. A row older than the one already
// shown is a stale snapshot and is ignored.
func (a *App) applyProgress(r *Row) {
	if r == nil {
		return
	}

	for i := range a.view.Active {
		if a.view.Active[i].ID != r.ID {
			continue
		}

		if a.view.Active[i].Updated.After(r.Updated) {
			return
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

		a.updateDetail(*r)

		return
	}

	// A job the active set does not know about yet: ask for a fresh set.
	a.needActive = true
}

// mergeActive replaces the active set, keeping a row whose shown state is
// newer than the snapshot's, so a snapshot captured before a transition
// cannot undo the event that announced it.
func (a *App) mergeActive(rows []Row) {
	shown := make(map[string]Row, len(a.view.Active))
	for _, row := range a.view.Active {
		shown[row.ID] = row
	}

	merged := make([]Row, 0, len(rows))

	for _, row := range rows {
		if old, ok := shown[row.ID]; ok && old.Updated.After(row.Updated) {
			row = old
		}

		merged = append(merged, row)
	}

	a.view.Active = merged
}
