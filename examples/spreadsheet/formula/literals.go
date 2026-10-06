package formula

import (
	"strconv"
	"strings"
	"unicode"
)

func (l *lexer) numberLit(at int) (token, error) {
	start := l.pos

	for l.pos < len(l.src) && (unicode.IsDigit(l.src[l.pos]) || l.src[l.pos] == '.') {
		l.pos++
	}

	if l.pos < len(l.src) && (l.src[l.pos] == 'e' || l.src[l.pos] == 'E') {
		l.pos++

		if l.pos < len(l.src) && (l.src[l.pos] == '+' || l.src[l.pos] == '-') {
			l.pos++
		}

		for l.pos < len(l.src) && unicode.IsDigit(l.src[l.pos]) {
			l.pos++
		}
	}

	text := string(l.src[start:l.pos])
	if _, err := strconv.ParseFloat(text, 64); err != nil {
		return token{}, &ParseError{
			Code:   ErrSyntax,
			Detail: "bad number " + text,
			At:     at,
		}
	}

	return token{kind: tokNumber, text: text, at: at}, nil
}

func (l *lexer) stringLit(at int) (token, error) {
	l.pos++

	var b strings.Builder

	for l.pos < len(l.src) {
		c := l.src[l.pos]

		if c != '"' {
			b.WriteRune(c)
			l.pos++
			continue
		}

		if l.pos+1 < len(l.src) && l.src[l.pos+1] == '"' {
			b.WriteRune('"')
			l.pos += 2
			continue
		}

		l.pos++

		return token{kind: tokString, text: b.String(), at: at}, nil
	}

	return token{}, &ParseError{
		Code:   ErrSyntax,
		Detail: "unterminated string",
		At:     at,
	}
}

func (l *lexer) wordLit(at int) (token, error) {
	start := l.pos

	for l.pos < len(l.src) && isWordRune(l.src[l.pos]) {
		l.pos++
	}

	return token{kind: tokWord, text: string(l.src[start:l.pos]), at: at}, nil
}

func isWordRune(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c) || c == '$' || c == '_'
}
