package ui

// buildCells lays out every cell of the rendered window as an inline block
// inside the grid container. The page background and the grid lines paint
// the surface, and one selection rectangle paints the highlight, so a cell
// costs one op unless it holds text. A cell that is being edited is left
// out; buildEditor draws it.
func (a *App) buildCells() []Box {
	w := a.win
	if w.Empty() {
		return nil
	}

	sel := a.selection()
	out := make([]Box, 0, w.Area().Count())
	for r := w.R0; r <= w.R1; r++ {
		for c := w.C0; c <= w.C1; c++ {
			if a.edit != nil && a.edit.Row == r && a.edit.Col == c {
				continue
			}

			cell, ok := a.cellAt(r, c)
			text := cell.Display
			if !ok || (text == "" && !cell.Num) {
				text = cell.Raw
			}

			cls := "cell"
			if cell.Num {
				cls += " num"
			}
			if cell.Err != "" || cell.Raw != "" && cell.Display == "" {
				cls += " err"
			}
			if r == sel.ActiveR && c == sel.ActiveC {
				cls += " act"
			}

			out = append(out, Box{
				Action: act("cell:" + itoa(r) + ":" + itoa(c)),
				Text:   text,
				Class:  cls,
				X:      cellX(c),
				Y:      cellY(r),
				W:      ColW,
				H:      RowH,
			})
		}
	}

	return out
}
