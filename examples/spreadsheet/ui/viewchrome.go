package ui

// buildToolbar lays out the frozen toolbar and the formula bar.
func (a *App) buildToolbar() []Box {
	x := a.scrollX + 6
	y := a.scrollY + 5
	theme := "Theme: Light"
	if a.dark {
		theme = "Theme: Dark"
	}

	out := make([]Box, 0, 6)
	for _, b := range []Box{
		{Action: "undo", Text: "Undo", W: 52},
		{Action: "redo", Text: "Redo", W: 52},
		{Action: "import", Text: "Import", W: 64},
		{Action: "export", Text: "Export", W: 64},
		{Action: "theme", Text: theme, W: 104},
	} {
		b.Y, b.H, b.X = y, ToolH-10, x
		b.Class = "btn"
		out = append(out, b)
		x += b.W + 6
	}

	fy := a.scrollY + ToolH + 4
	out = append(out, Box{Action: "formula", Text: "fx", Class: "fx", X: a.scrollX + 6, Y: fy, W: 22, H: FormulaH - 8})
	out = append(out, Box{Action: "formula", Text: a.activeRef(), Class: "reflab", X: a.scrollX + 30, Y: fy, W: 54, H: FormulaH - 8})
	out = append(out, Box{Action: "formula", Text: a.formulaText(), Class: "ftext", X: a.scrollX + 88, Y: fy, W: max(40, a.viewW-88-8), H: FormulaH - 8})

	return out
}
