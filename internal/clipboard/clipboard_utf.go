package clipboard

import "unicode/utf16"

// utf16Clipboard is CF_UNICODETEXT: UTF-16LE and a trailing zero word.
func utf16Clipboard(s string) []byte {
	u := utf16.Encode([]rune(s))
	b := make([]byte, (len(u)+1)*2)

	for i, v := range u {
		b[i*2] = byte(v)
		b[i*2+1] = byte(v >> 8)
	}

	return b
}
