package page

import (
	"html"
	"strings"
)

func replaceControlState(raw string, ctrl Control, st caretState) string {
	focused := st.focus != "" && ctrl.ID == st.focus
	kind := strings.ToLower(strings.TrimSpace(ctrl.Type))
	tag := strings.ToLower(strings.TrimSpace(ctrl.Tag))
	switch {
	case tag == "textarea":
		return fieldSpanState(raw, ctrl, st, focused)
	case tag == "select":
		return boxElement(raw, "select", selectLabel(ctrl), ctrl.ID, st.focus, focused && st.all)
	case kind == "checkbox" || kind == "radio":
		return openTag("input", raw, inputExtras(ctrl, kind, st.focus), true)
	case textLike(kind):
		return fieldSpanState(raw, ctrl, st, focused)
	default:
		return raw
	}
}

func boxElement(raw, tag, body, id, focusID string, selected bool) string {
	extra := []string{`data-gpui-field="` + tag + `"`}
	if focusID != "" && id == focusID {
		extra = append(extra, focusAttr)
	}
	if selected {
		extra = append(extra, selectedAttr)
	}

	return openTagDrop(tag, raw, extra, func(string) bool { return false }) + html.EscapeString(body) + "</" + tag + ">"
}

func inputExtras(ctrl Control, kind, focusID string) []string {
	extra := []string{`type="` + kind + `"`}
	if ctrl.Checked {
		extra = append(extra, "checked")
	}
	if ctrl.Disabled {
		extra = append(extra, "disabled")
	}
	if focusID != "" && ctrl.ID == focusID {
		extra = append(extra, focusAttr)
	}

	return extra
}

func selectLabel(ctrl Control) string {
	for _, opt := range ctrl.Options {
		if opt.Value == ctrl.Value {
			return opt.Label
		}
	}
	if len(ctrl.Options) > 0 {
		return ctrl.Options[0].Label
	}

	return ctrl.Value
}
