package ui

// Selection is a rectangular cell range. Anchor is the corner that stays
// while Shift extends the range; Active is the moving cell and names the
// formula bar. Both are zero-based and clamped into the sheet.
type Selection struct {
	AnchorR int
	AnchorC int
	ActiveR int
	ActiveC int
}

// newSelection returns a one-cell selection.
func newSelection(r, c int) Selection {
	return Selection{AnchorR: r, AnchorC: c, ActiveR: r, ActiveC: c}
}

// setSelection stores the active sheet's selection and clamps it.
func (a *App) setSelection(s Selection) {
	sh := a.sheet()
	a.sel[a.active] = s.Clamp(max(sh.Rows, 1), max(sh.Cols, 1))
}
