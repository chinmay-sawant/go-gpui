package corebackend

import (
	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/formula"
	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/ui"
	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/workbook"
)

// maxRangeCells bounds one fetch, so a caller cannot ask the adapter to
// materialize a whole sheet.
const maxRangeCells = 50000

// Range returns the cells of a bounded rectangle in row-major order.
func (b *Backend) Range(id string, a ui.Area) ([]ui.Cell, error) {
	sh, ok := b.sheet(id)
	if !ok {
		return nil, ErrNoSheet
	}

	if a.Empty() {
		return nil, nil
	}

	if a.Count() > maxRangeCells {
		return nil, ErrRangeTooLarge
	}

	out := make([]ui.Cell, 0, a.Count())
	for r := a.R0; r <= a.R1; r++ {
		for c := a.C0; c <= a.C1; c++ {
			p := workbook.Pos{Row: r, Col: c}
			cell, _ := sh.Cell(p)
			out = append(out, toUICell(cell, sh.Display(p)))
		}
	}

	return out, nil
}

// Used returns the sheet's used rectangle.
func (b *Backend) Used(id string) (ui.Area, bool) {
	sh, ok := b.sheet(id)
	if !ok {
		return ui.Area{}, false
	}

	r, ok := sh.UsedRange()
	if !ok {
		return ui.Area{}, false
	}

	return ui.Area{R0: r.Min.Row, C0: r.Min.Col, R1: r.Max.Row, C1: r.Max.Col}, true
}

// toUICell maps one core cell and its display value to the UI shape.
func toUICell(c workbook.Cell, display string) ui.Cell {
	out := ui.Cell{Raw: c.Raw(), Display: display}

	switch c.Kind {
	case workbook.Number:
		out.Num = true
	case workbook.Formula:
		out.Num = c.Value.Kind == formula.KindNumber
		if c.Value.IsError() {
			out.Err = display
		}
	}

	return out
}
