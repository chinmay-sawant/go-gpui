package workbook

// RecalcRequest asks for a recalculation pass. Rev is the content revision
// the caller based the request on; when the workbook has moved on, the
// request is stale and nothing runs.
type RecalcRequest struct {
	Rev     int64
	Sheet   SheetID
	Changed []Pos
}

// RecalcResult summarizes one pass.
type RecalcResult struct {
	Stale     bool // the workbook moved on; nothing ran
	Evaluated int  // formulas evaluated or marked with an error
	Truncated bool // a limit stopped the pass
	Cycled    int  // formulas left in a dependency cycle
}

// Recalc runs the request unless its revision is stale.
func (w *Workbook) Recalc(req RecalcRequest) RecalcResult {
	if req.Rev != 0 && req.Rev != w.rev {
		return RecalcResult{Stale: true}
	}

	s := w.Sheet(req.Sheet)
	if s == nil {
		return RecalcResult{}
	}

	return w.recalc(s, req.Changed)
}

// RecalcAll recalculates every formula in every sheet.
func (w *Workbook) RecalcAll() RecalcResult {
	var out RecalcResult

	for _, s := range w.sheets {
		res := w.recalc(s, s.formulaPositions())
		out.Evaluated += res.Evaluated
		out.Truncated = out.Truncated || res.Truncated
		out.Cycled += res.Cycled
	}

	return out
}
