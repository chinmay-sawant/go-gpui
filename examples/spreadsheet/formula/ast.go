package formula

// Expr is a parsed formula. A parsed expression is immutable and safe to
// share between goroutines.
type Expr interface{ expr() }

type numExpr struct{ v float64 }

type strExpr struct{ v string }

// refExpr covers a single cell and a rectangular range alike.
type refExpr struct{ rect Rect }

type unaryExpr struct {
	op byte
	x  Expr
}

type binaryExpr struct {
	op   byte
	l, r Expr
}

type callExpr struct {
	name string
	args []Expr
}

func (*numExpr) expr()    {}
func (*strExpr) expr()    {}
func (*refExpr) expr()    {}
func (*unaryExpr) expr()  {}
func (*binaryExpr) expr() {}
func (*callExpr) expr()   {}

// Refs returns every cell and range the expression reads, in walk order.
func Refs(e Expr) []Rect {
	var out []Rect
	walk(e, func(r Rect) { out = append(out, r) })

	return out
}

func walk(e Expr, fn func(Rect)) {
	switch t := e.(type) {
	case *refExpr:
		fn(t.rect)
	case *unaryExpr:
		walk(t.x, fn)
	case *binaryExpr:
		walk(t.l, fn)
		walk(t.r, fn)
	case *callExpr:
		for _, a := range t.args {
			walk(a, fn)
		}
	}
}

// ParseError describes why a formula did not parse. Code is an error code
// such as ErrSyntax or ErrRef, suitable for a cell value.
type ParseError struct {
	Code   string
	Detail string
	At     int
}

func (e *ParseError) Error() string {
	if e.Detail == "" {
		return e.Code
	}

	return e.Code + ": " + e.Detail
}
