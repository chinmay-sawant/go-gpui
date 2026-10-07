package workbook

import (
	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/formula"
)

// sheetEnv lets formula evaluation read cells of one sheet.
type sheetEnv struct{ s *Sheet }

// Cell returns the computed value of one cell.
func (e sheetEnv) Cell(row, col int) formula.Value {
	return e.s.Value(Pos{Row: row, Col: col})
}

// setFormulaValue caches a computed value on a formula cell.
func setFormulaValue(s *Sheet, p Pos, v formula.Value) {
	c := s.cells[p]
	c.Value = v
	s.cells[p] = c
}

// errorValue builds an error value, dropping an empty detail.
func errorValue(code, detail string) formula.Value {
	if detail == "" {
		return formula.Err(code)
	}

	return formula.ErrDetail(code, detail)
}
