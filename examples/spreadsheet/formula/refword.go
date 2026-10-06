package formula

import (
	"strconv"
	"strings"
)

// ParseRefWord parses a cell word such as A1 or $B$2 into one-based row and
// column numbers.
func ParseRefWord(w string) (row, col int, ok bool) {
	s := strings.TrimPrefix(w, "$")

	i := 0
	for i < len(s) && isASCIILetter(s[i]) {
		i++
	}

	if i == 0 || i > 3 {
		return 0, 0, false
	}

	letters := i

	if i < len(s) && s[i] == '$' {
		i++
	}

	j := i
	for j < len(s) && s[j] >= '0' && s[j] <= '9' {
		j++
	}

	if j == i || j != len(s) {
		return 0, 0, false
	}

	col = 0
	for k := 0; k < letters; k++ {
		col = col*26 + int(upper(s[k])-'A') + 1
	}

	n, err := strconv.Atoi(s[i:j])
	if err != nil {
		return 0, 0, false
	}

	return n, col, true
}

func isASCIILetter(c byte) bool {
	return c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z'
}

func upper(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return c - 'a' + 'A'
	}

	return c
}
