package page

import "strings"

func openTag(tag, raw string, extra []string, input bool) string {
	var b strings.Builder
	b.WriteByte('<')
	b.WriteString(tag)
	for _, attr := range splitAttrs(attrRegion(raw)) {
		if skipAttr(attrKey(attr), input) {
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

func skipAttr(name string, input bool) bool {
	if name == "data-gpui-focus" {
		return true
	}
	if !input {
		return false
	}

	return name == "type" || name == "checked" || name == "disabled"
}

func attrKey(attr string) string {
	cut := strings.IndexByte(attr, '=')
	if cut >= 0 {
		attr = attr[:cut]
	}

	return strings.ToLower(strings.TrimSpace(attr))
}

func attrRegion(raw string) string {
	end := len(raw)
	quote := byte(0)
	for i := 0; i < len(raw); i++ {
		c := raw[i]
		if quote != 0 {
			if c == quote {
				quote = 0
			}
			continue
		}
		if c == '"' || c == '\'' {
			quote = c
			continue
		}
		if c == '>' {
			end = i
			break
		}
	}

	i := 0
	if raw != "" && raw[0] == '<' {
		i = 1
	}
	for i < end && !isSpace(raw[i]) && raw[i] != '/' {
		i++
	}
	s := strings.TrimSpace(raw[i:end])

	return strings.TrimSpace(strings.TrimSuffix(s, "/"))
}

func splitAttrs(s string) []string {
	out := make([]string, 0)
	i := 0
	for i < len(s) {
		i = skipSpace(s, i)
		if i >= len(s) {
			break
		}
		start := i
		for i < len(s) && !isSpace(s[i]) && s[i] != '=' {
			i++
		}
		j := skipSpace(s, i)
		if j < len(s) && s[j] == '=' {
			i = skipValue(s, skipSpace(s, j+1))
		}
		out = append(out, strings.TrimSpace(s[start:i]))
	}

	return out
}
