package workbook

import (
	"errors"

	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/formula"
)

// Apply runs the command, recalculates the formulas it affects, and pushes
// the change onto the undo history. It returns the recalculation summary.
func (w *Workbook) Apply(cmd Command) (RecalcResult, error) {
	if cmd.Empty() {
		return RecalcResult{}, nil
	}

	s := w.Sheet(cmd.Sheet)
	if s == nil {
		return RecalcResult{}, ErrNoSheet
	}

	forward := Command{Sheet: cmd.Sheet}
	inverse := Command{Sheet: cmd.Sheet}
	seen := map[Pos]bool{}

	for _, e := range cmd.Edits {
		if !e.Pos.Valid() {
			return RecalcResult{}, ErrBadRef
		}

		if seen[e.Pos] {
			continue
		}

		seen[e.Pos] = true

		prev, ok := s.cells[e.Pos]
		if ok {
			inverse.Edits = append(inverse.Edits, CellEdit{Pos: e.Pos, Cell: prev})
		} else {
			inverse.Edits = append(inverse.Edits, CellEdit{Pos: e.Pos, Cell: Cell{}})
		}

		forward.Edits = append(forward.Edits, CellEdit{Pos: e.Pos, Cell: e.Cell})
		w.putCell(s, e.Pos, e.Cell)
	}

	w.push(histEntry{inverse: inverse, forward: forward})

	return w.bump(s, forward.Edits), nil
}

// putCell stores one cell and refreshes its parse, or clears the position.
func (w *Workbook) putCell(s *Sheet, p Pos, c Cell) {
	if c.Kind == Blank {
		delete(s.cells, p)
		delete(s.parsed, p)

		return
	}

	s.cells[p] = c

	if c.Kind != Formula {
		delete(s.parsed, p)

		return
	}

	pf := parsedFormula{}

	if e, err := formula.Parse(c.Source, w.lim); err != nil {
		pe := &formula.ParseError{Code: formula.ErrSyntax, Detail: err.Error()}
		errors.As(err, &pe)

		pf.err = pe
	} else {
		pf.expr = e
		pf.refs = formula.Refs(e)
	}

	s.parsed[p] = pf
}

// bump advances the revision and recalculates the affected formulas.
func (w *Workbook) bump(s *Sheet, edits []CellEdit) RecalcResult {
	w.rev++

	return w.recalc(s, changedPositions(edits))
}
