package page

import (
	"html"
	"strings"
	"unicode/utf8"
)

const (
	plainStyle = "display:inline-block;border:1px solid #c8c2b4;" +
		"padding:4px 6px;min-width:10em;min-height:1.6em;white-space:pre"
	focusStyle = "display:inline-block;border:2px solid #1a56db;" +
		"padding:3px 5px;min-width:10em;min-height:1.6em;white-space:pre"
	focusAttr = `data-gpui-focus="1"`
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

func textSpan(id, text string, focus bool) string {
	style := plainStyle
	if focus {
		style = focusStyle
	}

	return `<span id="` + html.EscapeString(id) + `" style="` + style + `">` + html.EscapeString(text) + "</span>"
}
