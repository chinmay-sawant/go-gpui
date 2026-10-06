package workbook

import (
	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/formula"
)

// parsedFormula is the cached parse of a formula cell.
type parsedFormula struct {
	expr formula.Expr
	refs []formula.Rect
	err  *formula.ParseError
}

// reads reports whether the formula reads the cell at p.
func (pf parsedFormula) reads(p Pos) bool {
	for _, r := range pf.refs {
		if r.Contains(p.Row, p.Col) {
			return true
		}
	}

	return false
}

// Sheet is one sparse grid.
type Sheet struct {
	id     SheetID
	name   string
	cells  map[Pos]Cell
	parsed map[Pos]parsedFormula
}

func newSheet(id SheetID, name string) *Sheet {
	return &Sheet{
		id:     id,
		name:   name,
		cells:  map[Pos]Cell{},
		parsed: map[Pos]parsedFormula{},
	}
}

// ID returns the stable sheet identifier.
func (s *Sheet) ID() SheetID { return s.id }

// SetID gives the sheet its storage identifier.
func (s *Sheet) SetID(id SheetID) { s.id = id }

// Name returns the sheet name.
func (s *Sheet) Name() string { return s.name }

// Cell returns the stored cell at p; the second result is false for a blank.
func (s *Sheet) Cell(p Pos) (Cell, bool) {
	c, ok := s.cells[p]

	return c, ok
}

// Value returns the computed value at p; a blank cell is KindBlank.
func (s *Sheet) Value(p Pos) formula.Value {
	c, ok := s.cells[p]
	if !ok {
		return formula.Value{}
	}

	return valueOf(c)
}

// Display returns the text the cell shows.
func (s *Sheet) Display(p Pos) string { return s.Value(p).Display() }

// Count is the number of non-blank cells.
func (s *Sheet) Count() int { return len(s.cells) }

func valueOf(c Cell) formula.Value {
	switch c.Kind {
	case Number:
		return formula.Number(c.Number)
	case Text:
		return formula.Text(c.Text)
	case Formula:
		return c.Value
	}

	return formula.Value{}
}
