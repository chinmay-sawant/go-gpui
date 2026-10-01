package page

import (
	"unicode"
	"unicode/utf8"
)

func dropLastRune(s string) string {
	if s == "" {
		return ""
	}

	_, size := utf8.DecodeLastRuneInString(s)

	return s[:len(s)-size]
}

func dropLastWord(s string) string {
	runes := []rune(s)
	i := len(runes)
	for i > 0 && unicode.IsSpace(runes[i-1]) {
		i--
	}

	for i > 0 && !unicode.IsSpace(runes[i-1]) {
		i--
	}

	return string(runes[:i])
}
