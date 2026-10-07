package formula

import (
	"strconv"
)

func (p *parser) unary() (Expr, error) {
	t := p.peek()
	if t.kind == tokOp && (t.text == "+" || t.text == "-") {
		p.take()

		x, err := p.unary()
		if err != nil {
			return nil, err
		}

		if t.text == "+" {
			return x, nil
		}

		return &unaryExpr{op: '-', x: x}, nil
	}

	return p.power()
}

func (p *parser) power() (Expr, error) {
	base, err := p.atom()
	if err != nil {
		return nil, err
	}

	t := p.peek()
	if t.kind == tokOp && t.text == "^" {
		p.take()

		// Right associative, and 2^-1 is allowed.
		exp, err := p.unary()
		if err != nil {
			return nil, err
		}

		return &binaryExpr{op: '^', l: base, r: exp}, nil
	}

	return base, nil
}

func (p *parser) number(text string) (Expr, error) {
	n, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return nil, &ParseError{Code: ErrSyntax, Detail: "bad number " + text}
	}

	return &numExpr{v: n}, nil
}
