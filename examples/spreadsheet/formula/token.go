package formula

import (
	"fmt"
	"strings"
	"unicode"
)

type tokKind uint8

const (
	tokEOF tokKind = iota
	tokNumber
	tokString
	tokWord
	tokOp
	tokLParen
	tokRParen
	tokComma
	tokColon
)

type token struct {
	kind tokKind
	text string
	at   int
}

type lexer struct {
	src []rune
	pos int
}

func lex(src string) ([]token, error) {
	l := &lexer{src: []rune(src)}

	var out []token

	for {
		t, err := l.next()
		if err != nil {
			return nil, err
		}

		out = append(out, t)

		if t.kind == tokEOF {
			return out, nil
		}
	}
}

func (l *lexer) next() (token, error) {
	for l.pos < len(l.src) && unicode.IsSpace(l.src[l.pos]) {
		l.pos++
	}

	at := l.pos

	if l.pos == len(l.src) {
		return token{kind: tokEOF, at: at}, nil
	}

	c := l.src[l.pos]

	switch {
	case c == '(':
		l.pos++
		return token{kind: tokLParen, text: "(", at: at}, nil
	case c == ')':
		l.pos++
		return token{kind: tokRParen, text: ")", at: at}, nil
	case c == ',':
		l.pos++
		return token{kind: tokComma, text: ",", at: at}, nil
	case c == ':':
		l.pos++
		return token{kind: tokColon, text: ":", at: at}, nil
	case strings.ContainsRune("+-*/^", c):
		l.pos++
		return token{kind: tokOp, text: string(c), at: at}, nil
	case c == '"':
		return l.stringLit(at)
	case unicode.IsDigit(c) || c == '.':
		return l.numberLit(at)
	case unicode.IsLetter(c) || c == '$' || c == '_':
		return l.wordLit(at)
	}

	return token{}, &ParseError{
		Code:   ErrSyntax,
		Detail: fmt.Sprintf("unexpected %q", c),
		At:     at,
	}
}
