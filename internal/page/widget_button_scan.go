package page

import "strings"

// hasButton reports whether source has a button element outside comments,
// script text, and style text.
func hasButton(source string) bool {
	i := 0
	for i < len(source) {
		next := strings.IndexByte(source[i:], '<')
		if next < 0 {
			return false
		}
		i += next
		if strings.HasPrefix(source[i:], "<!--") {
			end := strings.Index(source[i+4:], "-->")
			if end < 0 {
				return false
			}
			i += end + 7
			continue
		}
		name, _, end, ok := readOpen(source, i)
		if !ok {
			i++
			continue
		}
		low := strings.ToLower(name)
		if low == "button" {
			return true
		}
		if low == "script" || low == "style" {
			_, i = closeSpan(source, end, low)
			continue
		}
		i = end
	}

	return false
}
