package ui

// buildView fills a View from the app state, the rendered window, and the
// tile cache.
func (a *App) buildView() View {
	v := View{
		Dark:    a.dark,
		ViewW:   a.viewW,
		ViewH:   a.viewH,
		ScrollX: a.scrollX,
		ScrollY: a.scrollY,
		Ref:     a.activeRef(),
		Formula: a.formulaText(),
		Status:  a.status,
	}
	sh := a.sheet()
	v.TotalW = HeadW + sh.Cols*ColW
	v.TotalH = ChromeH + sh.Rows*RowH
	v.Toolbar = a.buildToolbar()
	v.Grid = a.buildGrid()
	v.SelRect = a.buildSelection()
	if w := a.win; !w.Empty() {
		v.GridX = cellX(w.C0)
		v.GridY = cellY(w.R0)
		v.GridW = (w.C1 - w.C0 + 1) * ColW
		v.GridH = (w.R1 - w.R0 + 1) * RowH
	}

	v.Cols = a.buildCols()
	v.Rows = a.buildRows()
	v.Cells = a.buildCells()
	v.Editor = a.buildEditor()
	v.Tabs = a.buildTabs()
	v.Dialog = a.buildDialog()

	return v
}

// sheet returns the active sheet, or an empty placeholder when the backend
// has no sheets at all.
func (a *App) sheet() Sheet {
	if sh, ok := a.info[a.active]; ok {
		return sh
	}

	return Sheet{ID: a.active, Rows: 0, Cols: 0}
}

// activeRef names the active cell.
func (a *App) activeRef() string {
	s := a.selection()

	return ref(s.ActiveR, s.ActiveC)
}
