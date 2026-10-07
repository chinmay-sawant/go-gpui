package formula

import (
	"strings"
)

func (p *parser) atom() (Expr, error) {
	t := p.take()

	switch t.kind {
	case tokNumber:
		return p.number(t.text)
	case tokString:
		return &strExpr{v: t.text}, nil
	case tokLParen:
		p.depth++
		if p.depth > p.lim.MaxDepth {
			return nil, &ParseError{Code: ErrLimit, Detail: "expression nests too deep", At: t.at}
		}

		e, err := p.expr()
		p.depth--

		if err != nil {
			return nil, err
		}

		if close := p.take(); close.kind != tokRParen {
			return nil, &ParseError{Code: ErrSyntax, Detail: "missing )", At: close.at}
		}

		return e, nil
	case tokWord:
		if p.peek().kind == tokLParen {
			return p.call(t)
		}

		return p.ref(t)
	}

	return nil, &ParseError{Code: ErrSyntax, Detail: "unexpected " + t.text, At: t.at}
}

func (p *parser) call(name token) (Expr, error) {
	p.depth++
	if p.depth > p.lim.MaxDepth {
		return nil, &ParseError{Code: ErrLimit, Detail: "expression nests too deep", At: name.at}
	}

	p.take() // "("

	args, err := p.args()
	p.depth--

	if err != nil {
		return nil, err
	}

	return &callExpr{name: strings.ToUpper(name.text), args: args}, nil
}

func (p *parser) args() ([]Expr, error) {
	var args []Expr

	for {
		if p.peek().kind == tokRParen {
			if len(args) == 0 {
				return nil, &ParseError{Code: ErrSyntax, Detail: "no arguments", At: p.peek().at}
			}

			p.take()

			return args, nil
		}

		e, err := p.expr()
		if err != nil {
			return nil, err
		}

		args = append(args, e)

		switch t := p.take(); t.kind {
		case tokComma:
			if p.peek().kind == tokRParen {
				return nil, &ParseError{Code: ErrSyntax, Detail: "expected an argument after ,", At: p.peek().at}
			}
		case tokRParen:
			return args, nil
		default:
			return nil, &ParseError{Code: ErrSyntax, Detail: "expected , or )", At: t.at}
		}
	}
}
