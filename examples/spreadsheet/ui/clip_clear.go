package ui

import (
	"context"
	"fmt"
)

// clearArea queues one batch of empty edits for a rectangle.
func (a *App) clearArea(area Area) {
	if area.Count() > maxClipCells {
		a.status = "range too large to clear"

		return
	}

	edits := make([]Edit, 0, area.Count())
	for r := area.R0; r <= area.R1; r++ {
		for c := area.C0; c <= area.C1; c++ {
			edits = append(edits, Edit{Row: r, Col: c, Raw: ""})
		}
	}

	a.applyOptimistic(a.active, edits)
	a.postEdits(a.active, edits)
	a.status = fmt.Sprintf("cleared %d cells", len(edits))
}

// clearSelection is the Delete key.
func (a *App) clearSelection(ctx context.Context) error {
	a.clearArea(a.selection().Area())

	return a.Redraw(ctx)
}
