package ui

// buildEditor draws the active editor inside its cell, with the caret at
// the stored position. A cell outside the rendered window is skipped; the
// buffer stays in the app and appears again when the cell returns.
func (a *App) buildEditor() *Editor {
	e := a.edit
	if e == nil || !a.winCovers(e.Row, e.Col) {
		return nil
	}

	runes := []rune(e.Text)
	at := clampInt(e.Caret, 0, len(runes))

	return &Editor{
		X:      cellX(e.Col),
		Y:      cellY(e.Row),
		W:      ColW,
		H:      RowH,
		Before: string(runes[:at]),
		After:  string(runes[at:]),
	}
}

// winCovers reports whether the rendered window holds one cell.
func (a *App) winCovers(r, c int) bool {
	w := a.win

	return !w.Empty() && r >= w.R0 && r <= w.R1 && c >= w.C0 && c <= w.C1
}

// cellAt returns the cached cell at a coordinate.
func (a *App) cellAt(r, c int) (Cell, bool) {
	return a.tiles.get(a.active, r, c)
}
