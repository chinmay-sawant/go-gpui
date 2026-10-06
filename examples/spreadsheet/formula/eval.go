package formula

// Env answers cell reads during evaluation. Row and column are zero-based.
type Env interface {
	Cell(row, col int) Value
}

// Work bounds how many cell reads one evaluation pass may spend. A nil Work
// means no budget.
type Work struct{ Left int }

// NewWork starts a budget of n cell reads.
func NewWork(n int) *Work { return &Work{Left: n} }

// Spend charges n reads and reports whether the budget covered them.
func (w *Work) Spend(n int) bool {
	if w == nil {
		return true
	}

	if n > w.Left {
		w.Left = 0

		return false
	}

	w.Left -= n

	return true
}

// Eval evaluates a parsed expression against env. Limits normalize
// themselves; work bounds the total cell reads.
func Eval(e Expr, env Env, lim Limits, work *Work) Value {
	return eval(e, env, lim.Normalized(), work)
}

func eval(e Expr, env Env, lim Limits, work *Work) Value {
	switch t := e.(type) {
	case *numExpr:
		return Number(t.v)
	case *strExpr:
		return Text(t.v)
	case *refExpr:
		return evalRef(t.rect, env, lim, work)
	case *unaryExpr:
		return unary(t.op, eval(t.x, env, lim, work))
	case *binaryExpr:
		l := eval(t.l, env, lim, work)
		if l.IsError() {
			return l
		}

		r := eval(t.r, env, lim, work)
		if r.IsError() {
			return r
		}

		return binary(t.op, l, r)
	case *callExpr:
		return evalCall(t, env, lim, work)
	}

	return Err(ErrSyntax)
}

// evalRef reads one cell. A range in scalar position has no single value.
func evalRef(r Rect, env Env, lim Limits, work *Work) Value {
	if r.MaxRow >= lim.MaxRows || r.MaxCol >= lim.MaxCols {
		return Err(ErrRef)
	}

	if r.Cells() > lim.MaxRangeCells {
		return Err(ErrLimit)
	}

	if r.MinRow != r.MaxRow || r.MinCol != r.MaxCol {
		return Err(ErrValue)
	}

	if !work.Spend(1) {
		return Err(ErrLimit)
	}

	return env.Cell(r.MinRow, r.MinCol)
}
