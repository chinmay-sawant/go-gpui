package ui

import "context"

// Edit bounds. A selection larger than this is not copied, cleared, or
// pasted as a whole, so a stray Ctrl+A on a 100k-row sheet stays cheap.
const (
	maxClipCells  = 4096
	maxPasteCells = 4096
)

// onCopy returns the editor text, or the selection as tab-separated rows.
func (a *App) onCopy(_ context.Context) (string, bool, error) {
	if a.edit != nil {
		return a.edit.Text, a.edit.Text != "", nil
	}

	return a.copyRange(false)
}

// onCut returns the same text and clears the source.
func (a *App) onCut(_ context.Context) (string, bool, error) {
	if a.edit != nil {
		e := a.edit
		a.clearArea(Area{e.Row, e.Col, e.Row, e.Col})
		a.edit = nil

		return e.Text, e.Text != "", nil
	}

	return a.copyRange(true)
}

// copyRange encodes the selection. With cut set, the source is cleared.
func (a *App) copyRange(cut bool) (string, bool, error) {
	area := a.selection().Area()
	if area.Count() > maxClipCells {
		a.status = "range too large to copy"

		return "", false, nil
	}

	cells := make([]Cell, 0, area.Count())
	for r := area.R0; r <= area.R1; r++ {
		for c := area.C0; c <= area.C1; c++ {
			cell, _ := a.cellAt(r, c)
			cells = append(cells, cell)
		}
	}

	if cut {
		a.clearArea(area)
	}

	return encodeTSV(cells, area.W(), area.H()), true, nil
}
