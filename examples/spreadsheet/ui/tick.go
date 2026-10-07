package ui

import (
	"context"
)

// drainBudget bounds the worker results one tick applies, so a burst of
// range loads and command answers cannot stall a frame.
const drainBudget = 64

// tick applies worker results on the UI loop with a bounded drain, then
// redraws once when anything painted differently.
func (a *App) tick(ctx context.Context) error {
	redraw := false

	for range drainBudget {
		select {
		case r := <-a.work.out:
			if a.applyResult(r) {
				redraw = true
			}
		default:
			if redraw {
				return a.Redraw(ctx)
			}

			return nil
		}
	}

	if redraw {
		return a.Redraw(ctx)
	}

	return nil
}

// applyResult folds one worker answer into app state and reports whether
// the frame changed.
func (a *App) applyResult(r result) bool {
	switch r.kind {
	case jobFetch:
		return a.applyFetch(r)
	case jobApply:
		return a.applyEdit(r)
	case jobUndo, jobRedo:
		return a.applyHistory(r)
	case jobUsed:
		return a.applyUsed(r)
	case jobReadFile, jobPreview, jobCommitCSV, jobExport:
		return a.applyCSV(r)
	case jobPref:
		if r.err != nil {
			a.status = "preference save failed: " + r.err.Error()

			return true
		}
	}

	return false
}

// applyUsed stores a used range and moves the selection to its last cell.
func (a *App) applyUsed(r result) bool {
	if r.err != nil {
		a.status = "used range failed: " + r.err.Error()

		return true
	}

	if r.sheet != a.active {
		return false
	}

	a.used[r.sheet] = r.area
	if r.ok {
		a.setSelection(newSelection(r.area.R1, r.area.C1))
		a.ensureVisible()
	} else {
		a.setSelection(newSelection(0, 0))
	}

	return true
}
