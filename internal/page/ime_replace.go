package page

import (
	"context"
	"unicode/utf8"
)

// IMEReplace replaces the focused value's bytes [start, end) with insert
// and puts the caret at the byte offset caret within insert. It draws when
// the value changed.
func (p *Page) IMEReplace(ctx context.Context, start, end int, insert string, caret int) error {
	c, ok := p.typingTarget()
	if !ok {
		return nil
	}

	runes := []rune(c.Value)
	startRune := runeAtByte(c.Value, start)
	endRune := runeAtByte(c.Value, end)
	if startRune == endRune && insert == "" {
		return nil
	}

	ins := []rune(insert)
	pos := startRune + clampPos(runeAtByte(insert, caret), len(ins))

	return p.editField(ctx, nil, func(c Control, _, _ int) (Control, int, int, bool) {
		out := make([]rune, 0, len(runes)-(endRune-startRune)+len(ins))
		out = append(out, runes[:startRune]...)
		out = append(out, ins...)
		out = append(out, runes[endRune:]...)

		return setRunes(c, out), pos, pos, true
	})
}

// runeAtByte converts a byte offset into value to a rune offset, clamped to
// the value and snapped to a rune boundary.
func runeAtByte(value string, off int) int {
	if off <= 0 {
		return 0
	}

	if off > len(value) {
		off = len(value)
	}

	for off > 0 && off < len(value) && !utf8.RuneStart(value[off]) {
		off--
	}

	return utf8.RuneCountInString(value[:off])
}
