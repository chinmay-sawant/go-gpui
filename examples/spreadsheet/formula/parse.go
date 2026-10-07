package formula

// Parse parses a formula source without the leading "=". The returned
// expression is immutable.
func Parse(src string, lim Limits) (Expr, error) {
	lim = lim.Normalized()

	if len([]rune(src)) > lim.MaxLen {
		return nil, &ParseError{Code: ErrLimit, Detail: "formula is too long"}
	}

	ts, err := lex(src)
	if err != nil {
		return nil, err
	}

	p := &parser{ts: ts, lim: lim}

	e, err := p.expr()
	if err != nil {
		return nil, err
	}

	if t := p.peek(); t.kind != tokEOF {
		return nil, &ParseError{Code: ErrSyntax, Detail: "unexpected " + t.text, At: t.at}
	}

	return e, nil
}

type parser struct {
	ts    []token
	pos   int
	lim   Limits
	depth int
}

func (p *parser) peek() token { return p.ts[p.pos] }

func (p *parser) take() token {
	t := p.ts[p.pos]
	p.pos++

	return t
}

func (p *parser) expr() (Expr, error) { return p.additive() }

func (p *parser) additive() (Expr, error) {
	l, err := p.multiplicative()
	if err != nil {
		return nil, err
	}

	for {
		t := p.peek()
		if t.kind != tokOp || (t.text != "+" && t.text != "-") {
			return l, nil
		}

		p.take()

		r, err := p.multiplicative()
		if err != nil {
			return nil, err
		}

		l = &binaryExpr{op: t.text[0], l: l, r: r}
	}
}

func (p *parser) multiplicative() (Expr, error) {
	l, err := p.unary()
	if err != nil {
		return nil, err
	}

	for {
		t := p.peek()
		if t.kind != tokOp || (t.text != "*" && t.text != "/") {
			return l, nil
		}

		p.take()

		r, err := p.unary()
		if err != nil {
			return nil, err
		}

		l = &binaryExpr{op: t.text[0], l: l, r: r}
	}
}
