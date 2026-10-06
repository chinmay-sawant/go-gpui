package page

import (
	"strings"
	"unicode/utf8"
)

const (
	focusAttr       = `data-ownframe-focus="1"`
	selectedAttr    = `data-ownframe-selected="1"`
	placeholderAttr = `data-ownframe-placeholder="1"`
)

func shownText(kind, value string) string {
	if kind == "password" {
		return strings.Repeat("•", utf8.RuneCountInString(value))
	}
	if kind == "file" && value == "" {
		return "No file"
	}

	return value
}
