// Package formula implements the spreadsheet example's formula language:
// arithmetic, cell references, rectangular ranges, SUM, and AVERAGE.
// Anything else evaluates to a visible error. Parsing and evaluation are
// bounded by Limits; this package stores no state.
package formula

import (
	"math"
	"strconv"
)

// Kind is the type of a computed value.
type Kind uint8

// The value kinds. A blank cell is distinct from the number zero.
const (
	KindBlank Kind = iota
	KindNumber
	KindText
	KindError
)

// Error codes the evaluator can produce.
const (
	ErrSyntax = "#ERROR!"
	ErrRef    = "#REF!"
	ErrDiv    = "#DIV/0!"
	ErrName   = "#NAME?"
	ErrValue  = "#VALUE!"
	ErrNum    = "#NUM!"
	ErrCycle  = "#CYCLE!"
	ErrLimit  = "#LIMIT!"
)

// FormatNumber renders a number without an exponent for everyday sizes.
func FormatNumber(n float64) string {
	if math.IsNaN(n) || math.IsInf(n, 0) {
		return ErrNum
	}

	if n == math.Trunc(n) && math.Abs(n) < 1e15 {
		return strconv.FormatFloat(n, 'f', -1, 64)
	}

	return strconv.FormatFloat(n, 'g', -1, 64)
}
