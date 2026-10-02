package page

func skipSpace(s string, i int) int {
	for i < len(s) && isSpace(s[i]) {
		i++
	}

	return i
}

func skipValue(s string, i int) int {
	if i >= len(s) {
		return i
	}
	if s[i] == '"' || s[i] == '\'' {
		q := s[i]
		i++
		for i < len(s) && s[i] != q {
			i++
		}
		if i < len(s) {
			i++
		}

		return i
	}
	for i < len(s) && !isSpace(s[i]) {
		i++
	}

	return i
}

func isSpace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r' || c == '\f'
}
