package page

import (
	"html"
	"strings"
)

func replaceControl(raw string, ctrl Control, focusID string) string {
	kind := strings.ToLower(strings.TrimSpace(ctrl.Type))
	tag := strings.ToLower(strings.TrimSpace(ctrl.Tag))
	switch {
	case tag == "textarea":
		return boxElement(raw, "textarea", ctrl.Value, ctrl.ID, focusID)
	case tag == "select":
		return boxElement(raw, "select", selectLabel(ctrl), ctrl.ID, focusID)
	case kind == "checkbox" || kind == "radio":
		return openTag("input", raw, inputExtras(ctrl, kind, focusID), true)
	case textLike(kind):
		return textSpan(ctrl.ID, shownText(kind, ctrl.Value), focusID == ctrl.ID && focusID != "")
	default:
		return raw
	}
}

func boxElement(raw, tag, body, id, focusID string) string {
	extra := []string{}
	if focusID != "" && id == focusID {
		extra = append(extra, focusAttr)
	}

	return openTag(tag, raw, extra, false) + html.EscapeString(body) + "</" + tag + ">"
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
