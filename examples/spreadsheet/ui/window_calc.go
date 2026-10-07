package ui

// visible returns the first and last cells that intersect the viewport.
// gridW and gridH are the cell area in CSS pixels; rows and cols are the
// sheet size. A window edge is clamped into the sheet.
func visible(scrollX, scrollY, gridW, gridH, rows, cols int) (fr, fc, lr, lc int) {
	if rows <= 0 || cols <= 0 {
		return 0, 0, -1, -1
	}

	if gridW < 1 {
		gridW = 1
	}

	if gridH < 1 {
		gridH = 1
	}

	fr = clampInt(scrollY/RowH, 0, rows-1)
	fc = clampInt(scrollX/ColW, 0, cols-1)
	lr = clampInt((scrollY+gridH-1)/RowH, 0, rows-1)
	lc = clampInt((scrollX+gridW-1)/ColW, 0, cols-1)

	return fr, fc, lr, lc
}

// computeWindow returns the rendered window for a viewport: the visible
// cells expanded by ov rows and columns on every side.
func computeWindow(sheet string, scrollX, scrollY, viewW, viewH, rows, cols, ov int) Window {
	gridW := viewW - HeadW
	gridH := viewH - ChromeH
	fr, fc, lr, lc := visible(scrollX, scrollY, gridW, gridH, rows, cols)
	if lr < fr || lc < fc {
		return Window{Sheet: sheet, R0: 0, C0: 0, R1: -1, C1: -1}
	}

	return Window{
		Sheet: sheet,
		R0:    max(0, fr-ov),
		C0:    max(0, fc-ov),
		R1:    min(rows-1, lr+ov),
		C1:    min(cols-1, lc+ov),
	}
}

// cellX is the content x of a cell column.
func cellX(c int) int { return HeadW + c*ColW }

// cellY is the content y of a cell row.
func cellY(r int) int { return ChromeH + r*RowH }

func clampInt(v, lo, hi int) int {
	if v < lo {
		return lo
	}

	if v > hi {
		return hi
	}

	return v
}
