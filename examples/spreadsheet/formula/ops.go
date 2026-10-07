package formula

import "math"

// unary applies a sign. Blank counts as zero; text is #VALUE!.
func unary(op byte, v Value) Value {
	if v.IsError() {
		return v
	}

	n, ok := toNumber(v)
	if !ok {
		return Err(ErrValue)
	}

	if op == '-' {
		n = -n
	}

	return Number(n)
}

// toNumber coerces a value for arithmetic. Blank counts as zero.
func toNumber(v Value) (float64, bool) {
	switch v.Kind {
	case KindBlank:
		return 0, true
	case KindNumber:
		return v.Num, true
	}

	return 0, false
}

// binary applies one arithmetic operator.
func binary(op byte, l, r Value) Value {
	a, ok := toNumber(l)
	if !ok {
		return Err(ErrValue)
	}

	b, ok := toNumber(r)
	if !ok {
		return Err(ErrValue)
	}

	var n float64

	switch op {
	case '+':
		n = a + b
	case '-':
		n = a - b
	case '*':
		n = a * b
	case '/':
		if b == 0 {
			return Err(ErrDiv)
		}

		n = a / b
	case '^':
		n = math.Pow(a, b)
	}

	if math.IsInf(n, 0) || math.IsNaN(n) {
		return Err(ErrNum)
	}

	return Number(n)
}
