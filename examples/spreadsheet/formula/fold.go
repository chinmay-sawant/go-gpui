package formula

// foldArg adds one argument to total. The Value result is an error to
// return; n counts the numbers the argument contributed; ok is false when
// the returned Value is an error.
func foldArg(a Expr, env Env, lim Limits, work *Work, total *float64) (Value, int, bool) {
	if r, isRef := a.(*refExpr); isRef {
		return foldRange(r.rect, env, lim, work, total)
	}

	v := eval(a, env, lim, work)
	if v.IsError() {
		return v, 0, false
	}

	if v.Kind == KindBlank {
		return Value{}, 0, true
	}

	n, ok := toNumber(v)
	if !ok {
		return Err(ErrValue), 0, false
	}

	*total += n

	return Value{}, 1, true
}

// foldRange sums the numbers in a referenced rectangle. Blanks and text are
// skipped; an error anywhere in the range wins.
func foldRange(r Rect, env Env, lim Limits, work *Work, total *float64) (Value, int, bool) {
	if r.MaxRow >= lim.MaxRows || r.MaxCol >= lim.MaxCols {
		return Err(ErrRef), 0, false
	}

	if r.Cells() > lim.MaxRangeCells {
		return Err(ErrLimit), 0, false
	}

	count := 0

	for row := r.MinRow; row <= r.MaxRow; row++ {
		for col := r.MinCol; col <= r.MaxCol; col++ {
			if !work.Spend(1) {
				return Err(ErrLimit), 0, false
			}

			v := env.Cell(row, col)
			if v.IsError() {
				return v, 0, false
			}

			if v.Kind == KindNumber {
				*total += v.Num
				count++
			}
		}
	}

	return Value{}, count, true
}
