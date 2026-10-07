package ui

// Grid geometry, in CSS pixels.
const (
	RowH     = 26
	ColW     = 112
	HeadW    = 52
	HeadH    = 24
	ToolH    = 34
	FormulaH = 28
	TabH     = 28
	// FetchOver is the fetched margin around the visible cells.
	FetchOver = 3
)

// ChromeH is the frozen band above the cells: toolbar, formula bar, and
// column headers.
const ChromeH = ToolH + FormulaH + HeadH

// Area is an inclusive rectangle of cell coordinates.
type Area struct {
	R0, C0, R1, C1 int
}

// W is the width in columns.
func (a Area) W() int { return a.C1 - a.C0 + 1 }

// H is the height in rows.
func (a Area) H() int { return a.R1 - a.R0 + 1 }

// Count is the number of cells.
func (a Area) Count() int { return a.W() * a.H() }

// Empty reports an area that covers no cell.
func (a Area) Empty() bool { return a.W() <= 0 || a.H() <= 0 }

// Window is the rendered cell window: the rows and columns currently in the
// document. The empty sentinel has R1 below R0.
type Window struct {
	Sheet          string
	R0, C0, R1, C1 int
}

// Area returns the window as an inclusive area.
func (w Window) Area() Area { return Area{w.R0, w.C0, w.R1, w.C1} }

// Empty reports a window that renders no cell.
func (w Window) Empty() bool { return w.R1 < w.R0 || w.C1 < w.C0 }

// Cover reports whether win covers every cell of area.
func (w Window) Cover(a Area) bool {
	return !w.Empty() && a.R0 >= w.R0 && a.R1 <= w.R1 && a.C0 >= w.C0 && a.C1 <= w.C1
}
