package formula

// evalCall evaluates the two supported functions. Any other name stays
// visible as #NAME?.
func evalCall(c *callExpr, env Env, lim Limits, work *Work) Value {
	switch c.name {
	case "SUM":
		return aggregate(c.args, env, lim, work, false)
	case "AVERAGE":
		return aggregate(c.args, env, lim, work, true)
	}

	return ErrDetail(ErrName, "unknown function "+c.name)
}

// aggregate sums or averages the arguments.
func aggregate(args []Expr, env Env, lim Limits, work *Work, average bool) Value {
	var total float64

	count := 0

	for _, a := range args {
		v, n, ok := foldArg(a, env, lim, work, &total)
		if !ok {
			return v
		}

		count += n
	}

	if average {
		if count == 0 {
			return Err(ErrDiv)
		}

		return Number(total / float64(count))
	}

	return Number(total)
}
