package page

import "unicode"

// isWordRune reports whether r is part of a word for a double-click.
func isWordRune(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_'
}

// wordBounds returns the word around pos as [start, end).
// A point past the last rune measures the word before it. A non-word rune
// returns an empty range at pos.
func wordBounds(s string, pos int) (int, int) {
	runes := []rune(s)
	n := len(runes)
	if n == 0 {
		return 0, 0
	}

	if pos >= n {
		pos = n - 1
	}

	if pos < 0 {
		pos = 0
	}

	if !isWordRune(runes[pos]) {
		return pos, pos
	}

	start := pos
	for start > 0 && isWordRune(runes[start-1]) {
		start--
	}

	end := pos + 1
	for end < n && isWordRune(runes[end]) {
		end++
	}

	return start, end
}

// lineBounds returns the line around pos as [start, end).
func lineBounds(s string, pos int) (int, int) {
	runes := []rune(s)
	n := len(runes)
	pos = clampPos(pos, n)

	start := 0
	for i := pos; i > 0; i-- {
		if runes[i-1] == '\n' {
			start = i

			break
		}
	}

	end := n
	for i := pos; i < n; i++ {
		if runes[i] == '\n' {
			end = i

			break
		}
	}

	return start, end
}

// prevWordStart returns the start of the word before pos.
func prevWordStart(s string, pos int) int {
	runes := []rune(s)
	i := clampPos(pos, len(runes))

	for i > 0 && unicode.IsSpace(runes[i-1]) {
		i--
	}

	for i > 0 && !unicode.IsSpace(runes[i-1]) {
		i--
	}

	return i
}

// nextWordStart returns the start of the word after pos.
func nextWordStart(s string, pos int) int {
	runes := []rune(s)
	i := clampPos(pos, len(runes))

	for i < len(runes) && unicode.IsSpace(runes[i]) {
		i++
	}

	for i < len(runes) && !unicode.IsSpace(runes[i]) {
		i++
	}

	for i < len(runes) && unicode.IsSpace(runes[i]) {
		i++
	}

	return i
}
