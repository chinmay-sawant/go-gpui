package window

import (
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// filter drops the runes of a chord key that already fired and is still
// down. The key can keep repeating after the modifier goes up, and those
// repeats must not reach the focused field. A key is forgotten only after
// its release has been quiet for the repeat gap, so a repeat reported as a
// release and a press keeps the key down.
func (w *chordWatch) filter(runes []rune) []rune {
	kept := runes[:0]

	for _, r := range runes {
		if !w.swallows(r) {
			kept = append(kept, r)
		}
	}

	for i, key := range chordKeys {
		if ebiten.IsKeyPressed(key) {
			continue
		}

		if w.released[i].IsZero() || time.Since(w.released[i]) >= chordRepeatGap {
			w.eaten[i] = ""
		}
	}

	return kept
}

// swallows reports whether r belongs to a fired key that is still down.
func (w *chordWatch) swallows(r rune) bool {
	for _, name := range w.eaten {
		if name != "" && strings.EqualFold(string(r), name) {
			return true
		}
	}

	return false
}
