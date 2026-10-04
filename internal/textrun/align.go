package textrun

// align maps each line to its rune offset in target. Only whitespace may be
// skipped between one line and the next, so a soft wrap that drops a space
// keeps the offsets right.
func align(lines []line, target string) ([]int, bool) {
	value := []rune(target)
	bases := make([]int, len(lines))
	pos := 0
	for i := range lines {
		text := []rune(lineText(lines[i]))
		if len(text) == 0 {
			bases[i] = pos

			continue
		}

		at := find(value, text, pos)
		if at < 0 {
			return nil, false
		}

		bases[i] = at
		pos = at + len(text)
	}

	return bases, true
}

// find returns the first index at or after from where want appears, with
// only whitespace between from and the match.
func find(value, want []rune, from int) int {
	for at := from; at+len(want) <= len(value); at++ {
		if !allSpace(value[from:at]) {
			return -1
		}

		if equalRunes(value[at:at+len(want)], want) {
			return at
		}
	}

	return -1
}

func allSpace(runes []rune) bool {
	for _, r := range runes {
		if r != ' ' && r != '\t' && r != '\n' && r != '\r' {
			return false
		}
	}

	return true
}

func equalRunes(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
