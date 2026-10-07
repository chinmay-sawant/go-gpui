package ui

// dropCovered removes overlapping tiles so a newer range wins lookups.
func (t *tileCache) dropCovered(sheet string, a Area) {
	kept := t.tiles[:0]
	for _, tl := range t.tiles {
		if tl.sheet == sheet && !disjoint(tl.area, a) {
			t.cells -= len(tl.cells)
			continue
		}

		kept = append(kept, tl)
	}

	t.tiles = kept
}

// setCell replaces one cached cell, so an optimistic edit paints at once.
func (t *tileCache) setCell(sheet string, r, c int, cell Cell) {
	for i := range t.tiles {
		tl := &t.tiles[i]
		if tl.sheet != sheet || c < tl.area.C0 || c > tl.area.C1 || r < tl.area.R0 || r > tl.area.R1 {
			continue
		}

		idx := (r-tl.area.R0)*tl.area.W() + (c - tl.area.C0)
		if idx >= 0 && idx < len(tl.cells) {
			tl.cells[idx] = cell
		}

		return
	}
}

func (t *tileCache) evictOldest() {
	if len(t.tiles) == 0 {
		return
	}

	old := 0
	for i := range t.tiles {
		if t.tiles[i].used < t.tiles[old].used {
			old = i
		}
	}

	t.cells -= len(t.tiles[old].cells)
	t.tiles = append(t.tiles[:old], t.tiles[old+1:]...)
}

func disjoint(a, b Area) bool {
	return a.R1 < b.R0 || b.R1 < a.R0 || a.C1 < b.C0 || b.C1 < a.C0
}

// overlapsWindow reports whether a fetched area touches the rendered
// window, so storing it changes a painted cell.
func overlapsWindow(a Area, w Window) bool {
	return !w.Empty() && !disjoint(a, w.Area())
}
