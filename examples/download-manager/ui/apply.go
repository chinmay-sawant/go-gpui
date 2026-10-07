package ui

// drain applies at most budget worker results without waiting.
func (a *App) drain(budget int) {
	for _, u := range a.backend.Poll(budget) {
		a.apply(u)
	}
}

// apply folds one worker result into the view. A progress update of a row
// that stays in place is paint-only; anything structural sets geom, so the
// next tick relayouts before it repaints.
func (a *App) apply(u Update) {
	a.view.Stats.Applied++

	switch u.Kind {
	case UpdateProgress:
		a.applyProgress(u.Row)
	case UpdateActive:
		a.mergeActive(u.Active)
		a.refreshDetail()
		a.historyDirty = true
		a.geom = true
	case UpdateHistory:
		a.applyHistory(u.Page)
	case UpdateSummary:
		if u.Gen == a.summaryGen && u.Summary != nil {
			a.view.Summary = *u.Summary
		}
	case UpdateNotice:
		a.view.Notice = u.Notice
		a.geom = true
	}
}
