package page

import "html"

// selectionSpan wraps selected runes in the library selection element, so
// the highlight paints inside the field and author field backgrounds cannot
// cover it.
func selectionSpan(text string) string {
	return `<span data-gpui-selection="1">` + html.EscapeString(text) + `</span>`
}

// markedText splits the shown text at the caret, or wraps the selected
// range in a selection span. A hidden caret paints no caret span.
func markedText(text string, focused bool, start, end, caret int, hidden bool) string {
	if !focused {
		return html.EscapeString(text)
	}

	runes := []rune(text)
	if end > start {
		return html.EscapeString(string(runes[:start])) +
			selectionSpan(string(runes[start:end])) +
			html.EscapeString(string(runes[end:]))
	}

	if hidden {
		return html.EscapeString(text)
	}

	return html.EscapeString(string(runes[:caret])) + caretSpan() + html.EscapeString(string(runes[caret:]))
}
