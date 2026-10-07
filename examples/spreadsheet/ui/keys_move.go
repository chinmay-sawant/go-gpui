package ui

import "context"

// onKeyUp tracks the modifier keys the window names.
func (a *App) onKeyUp(_ context.Context, key string) error {
	switch key {
	case "shift", "shiftleft", "shiftright":
		a.shift = false
	case "control", "controlleft", "controlright", "meta", "metaleft", "metaright":
		a.ctrl = false
	case "alt", "altleft", "altright":
		a.alt = false
	}

	return nil
}

// step moves or extends the selection and redraws.
func (a *App) step(ctx context.Context, dr, dc int, extend bool) error {
	s := a.selection()
	sh := a.sheet()
	if extend {
		s = s.Extend(dr, dc, sh.Rows, sh.Cols)
	} else {
		s = s.Move(dr, dc, sh.Rows, sh.Cols)
	}

	a.setSelection(s)
	a.ensureVisible()

	return a.Redraw(ctx)
}

// jump moves the active cell to one coordinate and redraws.
func (a *App) jump(ctx context.Context, c, r int) error {
	s := a.selection()
	s.ActiveR, s.ActiveC = r, c
	if !a.shift {
		s.AnchorR, s.AnchorC = r, c
	}

	a.setSelection(s)
	a.ensureVisible()

	return a.Redraw(ctx)
}

// jumpUsed moves to the last used cell. The used range comes from the
// worker, so an unknown sheet asks for it and moves when the answer lands.
func (a *App) jumpUsed(ctx context.Context) error {
	if area, ok := a.used[a.active]; ok {
		return a.jump(ctx, area.C1, area.R1)
	}

	a.work.post(job{kind: jobUsed, sheet: a.active})
	a.status = "finding used range"

	return a.Redraw(ctx)
}
