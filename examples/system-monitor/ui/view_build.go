package ui

// buildView snapshots the UI-loop state into template data.
func (a *App) buildView() View {
	s := a.state
	v := View{
		Dark:     s.dark,
		Live:     s.live,
		Nav:      s.nav,
		Mode:     s.mode,
		ModeLine: modeLine(s),
		At:       clock(s.at),
		TableAt:  clock(s.table.shown.At),
		Notice:   noticeText(s),
		Query:    s.table.query,
		SortKey:  string(s.table.key),
		SortDir:  sortArrow(s.table.desc),
		HasNew:   s.table.hasNew(),
		NewText:  newText(s.table.hasNew()),
	}

	v.Panels = buildPanels(s)

	pv := s.table.pageView()
	v.Rows = buildRows(pv.Rows, s.sel.id)
	v.PageNo, v.Pages = pv.No, pv.Pages
	v.Total, v.Shown = pv.Total, pv.Shown
	v.HasPrev = pv.No > 1
	v.HasNext = pv.No < pv.Pages
	v.Sel = buildSel(s)

	return v
}

// noticeText combines the transient notice and the collector problems.
func noticeText(s *state) string {
	if s.notice != "" {
		return s.notice
	}

	if len(s.problems) > 0 {
		return "collector: " + s.problems[0]
	}

	return ""
}
