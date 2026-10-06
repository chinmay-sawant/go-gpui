package page

import "github.com/chinmay-sawant/ownframe/internal/emoji"

// emojiCSS sizes the replacement images to the surrounding text. The
// negative alignment sinks the picture onto the baseline like a glyph.
const emojiCSS = `<style>img[data-ownframe-emoji]{width:1em;height:1em;vertical-align:-0.125em}</style>`

// emojiPass swaps supported emoji text for images and adds the sizing rule
// when anything changed. It runs after the control rewrite, so field
// values and painted text share one pass and caret spans are untouched.
func emojiPass(source string) string {
	out, changed := emoji.Replace(source)
	if !changed {
		return out
	}

	at := headOpen(out)
	if at < 0 {
		at = findHead(out)
	}

	if at < 0 {
		return emojiCSS + out
	}

	return out[:at] + emojiCSS + out[at:]
}
