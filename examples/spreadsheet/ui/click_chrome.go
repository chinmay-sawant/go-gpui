package ui

import (
	"context"
	"time"
)

// doubleClickGap is the window that separates a click from a double click.
const doubleClickGap = 400 * time.Millisecond

// clickCell selects a cell, extends with Shift, or opens the editor on a
// second click inside the double-click gap.
func (a *App) clickCell(ctx context.Context, action string) error {
	r, c, ok := parseCellAction(action)
	if !ok {
		return nil
	}

	now := time.Now()
	double := now.Sub(a.clickAt) < doubleClickGap && a.clickedR == r && a.clickedC == c
	a.clickAt, a.clickedR, a.clickedC = now, r, c

	if double {
		cell, _ := a.cellAt(r, c)
		a.startEdit(r, c, cell.Raw)

		return nil
	}

	if a.edit != nil {
		a.commitEdit(ctx)
	}

	if a.shift {
		s := a.selection()
		s.ActiveR, s.ActiveC = r, c
		a.setSelection(s)

		return nil
	}

	a.setSelection(newSelection(r, c))

	return nil
}

// selectRect replaces the selection with a rectangle.
func (a *App) selectRect(r0, c0, r1, c1 int) {
	a.setSelection(Selection{AnchorR: r0, AnchorC: c0, ActiveR: r1, ActiveC: c1})
}

// switchSheet activates one sheet and resets the scroll to its top left.
func (a *App) switchSheet(id string) {
	if _, ok := a.info[id]; !ok || id == a.active {
		return
	}

	a.edit = nil
	a.active = id
	a.win = Window{}
	a.fetchWin = Window{}
	a.page.ScrollTo(0, 0)
	a.fetchGen++
	a.status = a.sheet().Name
}

// history queues an undo or redo step. An open editor cancels first, the
// way Escape does.
func (a *App) history(undo bool) {
	if a.edit != nil {
		a.cancelEdit()

		return
	}

	kind := jobRedo
	if undo {
		kind = jobUndo
	}

	a.work.post(job{kind: kind})
}
