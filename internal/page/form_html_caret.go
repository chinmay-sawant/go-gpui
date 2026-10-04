package page

import (
	"html"
	"strings"
)

// fieldSpanState rewrites one text-like control into a span. It paints the
// caret at st.caret, or the range as a selection span, and the whole-value
// select as the field background. A negative caret means the end of the
// value.
func fieldSpanState(raw string, ctrl Control, st caretState, focused bool) string {
	kind := strings.ToLower(strings.TrimSpace(ctrl.Type))
	tag := strings.ToLower(strings.TrimSpace(ctrl.Tag))
	text := shownText(kind, ctrl.Value)
	ph := ""
	if text == "" && kind != "file" {
		ph = attrValue(raw, "placeholder")
	}

	n := runeLen(ctrl.Value)
	caret := st.caret
	if caret < 0 {
		caret = n
	}

	caret = clampPos(caret, n)
	start, end := rangeOf(st.start, st.end, n)
	whole := focused && (st.all || (start != end && start == 0 && end == n))

	extras := []string{`data-gpui-field="` + tag + `"`}
	if focused {
		extras = append(extras, focusAttr)
	}
	if whole {
		extras = append(extras, selectedAttr)
	}

	inner := ""
	switch {
	case ph != "":
		inner = html.EscapeString(ph)
		if focused && !whole {
			inner = caretSpan() + inner
		}
	case whole:
		inner = html.EscapeString(text)
	default:
		inner = markedText(text, focused, start, end, caret)
	}

	return openTagDrop("span", raw, extras, spanDrop) + inner + "</span>"
}

func caretSpan() string {
	return `<span data-gpui-caret="1"></span>`
}

// markedText splits the shown text at the caret, or wraps the selected
// range in a selection span.
func markedText(text string, focused bool, start, end, caret int) string {
	if !focused {
		return html.EscapeString(text)
	}

	runes := []rune(text)
	if end > start {
		return html.EscapeString(string(runes[:start])) +
			`<span data-gpui-selection="1">` + html.EscapeString(string(runes[start:end])) + `</span>` +
			html.EscapeString(string(runes[end:]))
	}

	return html.EscapeString(string(runes[:caret])) + caretSpan() + html.EscapeString(string(runes[caret:]))
}
