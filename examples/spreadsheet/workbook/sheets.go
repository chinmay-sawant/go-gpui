package workbook

// SheetByName returns the first sheet with the name, or nil.
func (w *Workbook) SheetByName(name string) *Sheet {
	for _, s := range w.sheets {
		if s.name == name {
			return s
		}
	}

	return nil
}

// AddSheet appends a sheet with a fresh local ID. It returns the existing
// sheet when the name is taken.
func (w *Workbook) AddSheet(name string) *Sheet {
	if s := w.SheetByName(name); s != nil {
		return s
	}

	var next SheetID = 1

	for _, s := range w.sheets {
		if s.id >= next {
			next = s.id + 1
		}
	}

	s := newSheet(next, name)
	w.sheets = append(w.sheets, s)
	w.bySheet[s.id] = s

	return s
}

// LoadSheet adds a sheet with a known storage ID without touching history.
// The storage package rebuilds persisted workbooks with it.
func (w *Workbook) LoadSheet(id SheetID, name string) *Sheet {
	s := newSheet(id, name)
	w.sheets = append(w.sheets, s)
	w.bySheet[id] = s

	return s
}

// RemoveSheet drops a sheet and its history. The last sheet stays.
func (w *Workbook) RemoveSheet(id SheetID) error {
	s := w.bySheet[id]
	if s == nil {
		return ErrNoSheet
	}

	if len(w.sheets) == 1 {
		return ErrNoSheet
	}

	delete(w.bySheet, id)

	for i, other := range w.sheets {
		if other == s {
			w.sheets = append(w.sheets[:i], w.sheets[i+1:]...)

			break
		}
	}

	w.history = nil
	w.future = nil

	return nil
}

// SetCell stores one cell without advancing the revision or the history.
// The storage load path uses it before a single RecalcAll.
func (w *Workbook) SetCell(sheet SheetID, p Pos, c Cell) error {
	s := w.Sheet(sheet)
	if s == nil {
		return ErrNoSheet
	}

	if !p.Valid() {
		return ErrBadRef
	}

	w.putCell(s, p, c)

	return nil
}
