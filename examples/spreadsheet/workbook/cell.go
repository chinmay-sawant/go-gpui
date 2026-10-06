package workbook

import (
	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/formula"
)

// Kind is the storage type of a cell.
type Kind uint8

// The cell kinds. Blank cells are absent; a Number whose value is zero is a
// stored cell and is not the same as a blank.
const (
	Blank Kind = iota
	Number
	Text
	Formula
)

// Cell is one cell. Formula cells keep their Source and a cached Value.
type Cell struct {
	Kind   Kind
	Number float64
	Text   string
	Source string
	Value  formula.Value
}

// ParseInput turns typed text into a cell: empty is blank, a bare number is
// a number, a leading = starts a formula, anything else is text.
func ParseInput(s string) Cell {
	if s == "" {
		return Cell{}
	}

	if s[0] == '=' {
		return Cell{Kind: Formula, Source: s[1:]}
	}

	if n, ok := plainNumber(s); ok {
		return Cell{Kind: Number, Number: n}
	}

	return Cell{Kind: Text, Text: s}
}

// Display renders the value the cell shows.
func (c Cell) Display() string {
	switch c.Kind {
	case Number:
		return formula.FormatNumber(c.Number)
	case Text:
		return c.Text
	case Formula:
		return c.Value.Display()
	}

	return ""
}

// Raw renders the text an editor loads for the cell.
func (c Cell) Raw() string {
	switch c.Kind {
	case Number:
		return formula.FormatNumber(c.Number)
	case Text:
		return c.Text
	case Formula:
		return "=" + c.Source
	}

	return ""
}
