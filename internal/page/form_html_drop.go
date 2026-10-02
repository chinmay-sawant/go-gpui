package page

import "strings"

func openTagDrop(tag, raw string, extra []string, drop func(string) bool) string {
	var b strings.Builder
	b.WriteByte('<')
	b.WriteString(tag)
	for _, attr := range splitAttrs(attrRegion(raw)) {
		name := attrKey(attr)
		if strings.HasPrefix(name, "data-gpui-") || drop(name) {
			continue
		}
		b.WriteByte(' ')
		b.WriteString(attr)
	}
	for _, attr := range extra {
		b.WriteByte(' ')
		b.WriteString(attr)
	}
	b.WriteByte('>')

	return b.String()
}
