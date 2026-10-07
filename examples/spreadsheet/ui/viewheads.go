package ui

// buildCols lays out the frozen column headers for the rendered window.
func (a *App) buildCols() []Box {
	w := a.win
	if w.Empty() {
		return nil
	}

	_, c0, _, c1 := a.selection().Rect()
	out := make([]Box, 0, w.C1-w.C0+1)
	for c := w.C0; c <= w.C1; c++ {
		cls := "head ch"
		if c >= c0 && c <= c1 {
			cls += " on"
		}

		out = append(out, Box{
			Action: act("col:" + itoa(c)),
			Text:   colName(c),
			Class:  cls,
			X:      cellX(c),
			Y:      a.scrollY + ToolH + FormulaH,
			W:      ColW,
			H:      HeadH,
		})
	}

	return out
}

// buildRows lays out the frozen row headers for the rendered window.
func (a *App) buildRows() []Box {
	w := a.win
	if w.Empty() {
		return nil
	}

	r0, _, r1, _ := a.selection().Rect()
	out := make([]Box, 0, w.R1-w.R0+1)
	for r := w.R0; r <= w.R1; r++ {
		cls := "head rh"
		if r >= r0 && r <= r1 {
			cls += " on"
		}

		out = append(out, Box{
			Action: act("row:" + itoa(r)),
			Text:   itoa(r + 1),
			Class:  cls,
			X:      a.scrollX,
			Y:      cellY(r),
			W:      HeadW,
			H:      RowH,
		})
	}

	return out
}
