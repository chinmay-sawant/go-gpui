package ui

// Rect returns the normalized inclusive corners.
func (s Selection) Rect() (r0, c0, r1, c1 int) {
	r0, r1 = minMax(s.AnchorR, s.ActiveR)
	c0, c1 = minMax(s.AnchorC, s.ActiveC)

	return r0, c0, r1, c1
}

// Area returns the selection as an area.
func (s Selection) Area() Area {
	r0, c0, r1, c1 := s.Rect()

	return Area{r0, c0, r1, c1}
}

// Contains reports whether a cell is inside the range.
func (s Selection) Contains(r, c int) bool {
	r0, c0, r1, c1 := s.Rect()

	return r >= r0 && r <= r1 && c >= c0 && c <= c1
}

// Move shifts the whole selection and keeps the active cell visible.
func (s Selection) Move(dr, dc, rows, cols int) Selection {
	r := clampInt(s.ActiveR+dr, 0, rows-1)
	c := clampInt(s.ActiveC+dc, 0, cols-1)

	return newSelection(r, c)
}

// Extend moves the active cell and keeps the anchor, growing the range.
func (s Selection) Extend(dr, dc, rows, cols int) Selection {
	s.ActiveR = clampInt(s.ActiveR+dr, 0, rows-1)
	s.ActiveC = clampInt(s.ActiveC+dc, 0, cols-1)

	return s
}

// Clamp pulls both corners into a sheet size, which may have changed after
// a backend reload.
func (s Selection) Clamp(rows, cols int) Selection {
	s.AnchorR = clampInt(s.AnchorR, 0, rows-1)
	s.AnchorC = clampInt(s.AnchorC, 0, cols-1)
	s.ActiveR = clampInt(s.ActiveR, 0, rows-1)
	s.ActiveC = clampInt(s.ActiveC, 0, cols-1)

	return s
}

func minMax(a, b int) (int, int) {
	if a > b {
		return b, a
	}

	return a, b
}
