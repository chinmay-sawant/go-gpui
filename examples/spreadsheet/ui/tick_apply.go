package ui

// applyFetch stores a range load unless it is stale. A load that misses the
// rendered window changes no painted cell, so it is paint-only.
func (a *App) applyFetch(r result) bool {
	if r.err != nil {
		a.status = "load failed: " + r.err.Error()

		return true
	}

	if r.gen != a.fetchGen || r.sheet != a.active {
		return false
	}

	a.tiles.put(r.sheet, r.area, r.cells, a.rev)

	return overlapsWindow(r.area, a.win)
}

// applyEdit confirms a stored edit. A failure restores the editor so the
// typed text is not lost.
func (a *App) applyEdit(r result) bool {
	if r.err != nil {
		a.status = "save failed: " + r.err.Error()
		if len(r.edits) == 1 {
			e := r.edits[0]
			a.startEdit(e.Row, e.Col, e.Raw)
		}

		return true
	}

	a.status = ""
	a.onRev(r.rev)

	return true
}

// onRev folds in a new workbook revision: cached tiles and used ranges are
// invalidated and the fetch window is requested again.
func (a *App) onRev(rev uint64) {
	if rev == a.rev {
		return
	}

	a.rev = rev
	a.tiles.invalidate()
	a.used = map[string]Area{}
	a.requestWindow()
}

// applyHistory confirms an undo or redo step.
func (a *App) applyHistory(r result) bool {
	if r.err != nil {
		a.status = "history failed: " + r.err.Error()

		return true
	}

	if !r.und.OK {
		a.status = "nothing to undo"
		if r.kind == jobRedo {
			a.status = "nothing to redo"
		}

		return true
	}

	a.status = r.und.Label
	a.onRev(r.und.Rev)

	return true
}
