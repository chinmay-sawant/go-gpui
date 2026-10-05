package emoji

import "unicode/utf8"

// match returns the longest supported sequence at byte i and its file.
// Sequences are three, four, or six bytes: a three-byte symbol with or
// without its three-byte variation selector, or one four-byte face.
func match(html string, i int) (string, string) {
	for _, n := range []int{6, 4, 3} {
		if i+n <= len(html) {
			if file, ok := bySeq[html[i:i+n]]; ok {
				return html[i : i+n], file
			}
		}
	}

	return "", ""
}

// decodeRune returns the rune and its byte size at i, or a one-byte step
// on invalid input.
func decodeRune(html string, i int) (rune, int) {
	r, size := utf8.DecodeRuneInString(html[i:])
	if r == utf8.RuneError && size <= 1 {
		return r, 1
	}

	return r, size
}
