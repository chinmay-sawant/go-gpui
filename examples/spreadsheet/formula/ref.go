package formula

func (p *parser) ref(t token) (Expr, error) {
	row, col, ok := ParseRefWord(t.text)
	if !ok {
		return nil, &ParseError{Code: ErrName, Detail: "unknown name " + t.text, At: t.at}
	}

	if row < 1 || row > p.lim.MaxRows || col < 1 || col > p.lim.MaxCols {
		return nil, &ParseError{Code: ErrRef, Detail: "reference " + t.text + " is out of range", At: t.at}
	}

	r := Rect{MinRow: row - 1, MinCol: col - 1, MaxRow: row - 1, MaxCol: col - 1}

	if p.peek().kind != tokColon {
		return &refExpr{rect: r}, nil
	}

	p.take() // ":"

	end := p.take()
	if end.kind != tokWord {
		return nil, &ParseError{Code: ErrSyntax, Detail: "expected a cell after :", At: end.at}
	}

	row2, col2, ok := ParseRefWord(end.text)
	if !ok {
		return nil, &ParseError{Code: ErrRef, Detail: "bad range end " + end.text, At: end.at}
	}

	if row2 < 1 || row2 > p.lim.MaxRows || col2 < 1 || col2 > p.lim.MaxCols {
		return nil, &ParseError{Code: ErrRef, Detail: "reference " + end.text + " is out of range", At: end.at}
	}

	r2 := Rect{MinRow: row2 - 1, MinCol: col2 - 1, MaxRow: row2 - 1, MaxCol: col2 - 1}

	r = Rect{
		MinRow: min(r.MinRow, r2.MinRow),
		MinCol: min(r.MinCol, r2.MinCol),
		MaxRow: max(r.MaxRow, r2.MaxRow),
		MaxCol: max(r.MaxCol, r2.MaxCol),
	}

	if r.Cells() > p.lim.MaxRangeCells {
		return nil, &ParseError{Code: ErrLimit, Detail: "range is too large", At: t.at}
	}

	return &refExpr{rect: r}, nil
}
