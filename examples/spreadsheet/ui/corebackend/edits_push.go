package corebackend

import "github.com/chinmay-sawant/ownframe/examples/spreadsheet/workbook"

// push records a saved step and drops the redo stack.
func (b *Backend) push(s step) {
	b.undo = append(b.undo, s)
	if len(b.undo) > undoLimit {
		b.undo = b.undo[len(b.undo)-undoLimit:]
	}

	b.redo = nil
}

// inverse captures the cells a command replaces, so a save failure or an
// undo can put them back.
func (b *Backend) inverse(sh *workbook.Sheet, edits []workbook.CellEdit) workbook.Command {
	inv := workbook.Command{Sheet: sh.ID()}

	for _, e := range edits {
		old, _ := sh.Cell(e.Pos)
		inv.Edits = append(inv.Edits, workbook.CellEdit{Pos: e.Pos, Cell: old})
	}

	return inv
}
