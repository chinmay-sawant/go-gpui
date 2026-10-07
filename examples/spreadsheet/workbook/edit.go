package workbook

// CellEdit replaces the cell at Pos. A zero Cell clears the position.
type CellEdit struct {
	Pos  Pos
	Cell Cell
}

// Command is one atomic set of cell edits in one sheet.
type Command struct {
	Sheet SheetID
	Edits []CellEdit
}

// Empty reports whether the command changes nothing.
func (c Command) Empty() bool { return len(c.Edits) == 0 }

// changedPositions returns the distinct positions the edits touch.
func changedPositions(edits []CellEdit) []Pos {
	seen := map[Pos]bool{}

	out := make([]Pos, 0, len(edits))

	for _, e := range edits {
		if !seen[e.Pos] {
			seen[e.Pos] = true
			out = append(out, e.Pos)
		}
	}

	return out
}
