package workbook

import (
	"sort"
)

// Each visits every stored cell in row-major order. It stops when fn
// returns false.
func (s *Sheet) Each(fn func(Pos, Cell) bool) {
	for _, p := range s.positions() {
		if !fn(p, s.cells[p]) {
			return
		}
	}
}

// UsedRange is the smallest rectangle covering the stored cells.
func (s *Sheet) UsedRange() (Rect, bool) {
	if len(s.cells) == 0 {
		return Rect{}, false
	}

	minR, minC := MaxRows, MaxCols
	maxR, maxC := -1, -1

	for p := range s.cells {
		if p.Row < minR {
			minR = p.Row
		}

		if p.Row > maxR {
			maxR = p.Row
		}

		if p.Col < minC {
			minC = p.Col
		}

		if p.Col > maxC {
			maxC = p.Col
		}
	}

	return Rect{Min: Pos{Row: minR, Col: minC}, Max: Pos{Row: maxR, Col: maxC}}, true
}

// positions returns the stored positions in row-major order.
func (s *Sheet) positions() []Pos {
	out := make([]Pos, 0, len(s.cells))
	for p := range s.cells {
		out = append(out, p)
	}

	sortPositions(out)

	return out
}

// formulaPositions returns the formula cells in row-major order.
func (s *Sheet) formulaPositions() []Pos {
	out := make([]Pos, 0, len(s.parsed))
	for p := range s.parsed {
		out = append(out, p)
	}

	sortPositions(out)

	return out
}

// dependents returns the formulas that read p, in position order.
func (s *Sheet) dependents(p Pos) []Pos {
	var out []Pos

	for f, pf := range s.parsed {
		if pf.reads(p) {
			out = append(out, f)
		}
	}

	sortPositions(out)

	return out
}

func sortPositions(out []Pos) {
	sort.Slice(out, func(i, j int) bool {
		if out[i].Row != out[j].Row {
			return out[i].Row < out[j].Row
		}

		return out[i].Col < out[j].Col
	})
}
