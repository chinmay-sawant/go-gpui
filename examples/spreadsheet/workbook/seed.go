package workbook

import (
	"fmt"
)

// Fixture sizes.
const (
	DummyRows  = 200
	DummyCols  = 20
	StressRows = 100000
)

// SeedDummy builds the shipped fixtures: a three-sheet demo workbook of
// DummyRows by DummyCols cells and a sparse StressRows-row stress workbook.
// The content is deterministic, so seeding twice writes the same cells.
func SeedDummy() []*Workbook {
	demo := New(1, "Demo workbook")

	fillNumbers(demo, demo.AddSheet("Numbers"))
	fillText(demo, demo.AddSheet("Text"))
	fillMixed(demo, demo.AddSheet("Mixed"))
	demo.RecalcAll()

	return []*Workbook{demo, SeedStress(StressRows)}
}

// put stores a cell during seeding.
func (w *Workbook) put(s *Sheet, p Pos, c Cell) { w.putCell(s, p, c) }

// formulaCell builds a formula cell from a format string.
func formulaCell(format string, args ...any) Cell {
	return Cell{Kind: Formula, Source: fmt.Sprintf(format, args...)}
}
