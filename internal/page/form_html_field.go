package page

import (
	"html"
	"strings"
)

func fieldSpan(raw string, ctrl Control, focus, selected bool) string {
	kind := strings.ToLower(strings.TrimSpace(ctrl.Type))
	tag := strings.ToLower(strings.TrimSpace(ctrl.Tag))
	text := shownText(kind, ctrl.Value)
	ph := ""
	if text == "" && kind != "file" {
		ph = attrValue(raw, "placeholder")
	}
	extras := []string{`data-gpui-field="` + tag + `"`}
	if focus {
		extras = append(extras, focusAttr)
	}
	if selected {
		extras = append(extras, selectedAttr)
	}
	if ph != "" {
		extras = append(extras, placeholderAttr)
	}
	inner := html.EscapeString(text)
	if ph != "" {
		inner = html.EscapeString(ph)
	}
	caret := ""
	if focus && !selected {
		caret = `<span data-gpui-caret="1"></span>`
	}
	if ph != "" {
		inner = caret + inner
	} else {
		inner += caret
	}

	return openTagDrop("span", raw, extras, spanDrop) + inner + "</span>"
}

func attrValue(raw, name string) string {
	for _, attr := range splitAttrs(attrRegion(raw)) {
		if attrKey(attr) != name {
			continue
		}
		cut := strings.IndexByte(attr, '=')
		if cut < 0 {
			return ""
		}
		val := strings.TrimSpace(attr[cut+1:])
		if len(val) >= 2 && (val[0] == '"' || val[0] == '\'') && val[len(val)-1] == val[0] {
			val = val[1 : len(val)-1]
		}

		return html.UnescapeString(val)
	}

	return ""
}

func spanDrop(name string) bool {
	return name == "value" || name == "type"
}
