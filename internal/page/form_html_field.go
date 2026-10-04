package page

import (
	"html"
	"strings"
)

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
