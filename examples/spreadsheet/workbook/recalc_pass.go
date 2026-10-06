package workbook

import (
	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/formula"
)

// recalc evaluates every formula that transitively depends on changed,
// plus every formula among changed, in dependency order.
func (w *Workbook) recalc(s *Sheet, changed []Pos) RecalcResult {
	res := RecalcResult{}

	affected, blocked := w.collect(s, changed)

	list := make([]Pos, 0, len(affected))
	for p := range affected {
		list = append(list, p)
	}

	sortPositions(list)

	order, cyclic := topo(s, list)

	for _, p := range cyclic {
		if blocked[p] {
			continue
		}

		setFormulaValue(s, p, errorValue(formula.ErrCycle, ""))
		res.Cycled++
		res.Evaluated++
	}

	work := formula.NewWork(w.lim.MaxEvals)

	for _, p := range order {
		if !work.Spend(1) {
			setFormulaValue(s, p, errorValue(formula.ErrLimit, "evaluation budget exhausted"))
			res.Truncated = true
			res.Evaluated++

			continue
		}

		pf := s.parsed[p]

		if pf.err != nil {
			setFormulaValue(s, p, errorValue(pf.err.Code, pf.err.Detail))
			res.Evaluated++

			continue
		}

		v := formula.Eval(pf.expr, sheetEnv{s: s}, w.lim, work)
		setFormulaValue(s, p, v)

		if v.Code == formula.ErrLimit {
			res.Truncated = true
		}

		res.Evaluated++
	}

	for _, p := range list {
		if !blocked[p] {
			continue
		}

		setFormulaValue(s, p, errorValue(formula.ErrLimit, "dependency depth limit"))
		res.Truncated = true
		res.Evaluated++
	}

	return res
}
