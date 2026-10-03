package window

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// chordKeys are the keys that start a text chord.
var chordKeys = [...]ebiten.Key{
	ebiten.KeyC,
	ebiten.KeyV,
	ebiten.KeyX,
	ebiten.KeyA,
	ebiten.KeyZ,
	ebiten.KeyY,
	ebiten.KeyInsert,
	ebiten.KeyDelete,
}

// chordRepeatGap is the shortest time a key can be up between two real
// presses. A server that reports auto-repeat as a release and a press
// leaves far less than this between the two.
const chordRepeatGap = 50 * time.Millisecond

// chordWatch fires a chord once per press, and keeps a fired key from
// typing while it stays down. A server can keep repeating the key after
// the modifier is released, and those repeats must not turn into text.
type chordWatch struct {
	released [len(chordKeys)]time.Time
	eaten    [len(chordKeys)]string
}

// newChordWatch returns a watch ready for the first press.
func newChordWatch() chordWatch {
	return chordWatch{}
}

// pressed returns the chord one just-pressed key starts, if any.
func (w *chordWatch) pressed(mods modifiers) chord {
	now := time.Now()

	for i, key := range chordKeys {
		justPressed := inpututil.IsKeyJustPressed(key)
		justReleased := inpututil.IsKeyJustReleased(key)

		if !w.step(i, justPressed, justReleased, now) {
			continue
		}

		if c := shortcutChord(mods, key); c != chordNone {
			w.eaten[i] = ebiten.KeyName(key)

			return c
		}
	}

	return chordNone
}

// step records one key's frame and reports a press that is not a repeat.
func (w *chordWatch) step(i int, justPressed, justReleased bool, now time.Time) bool {
	if justReleased {
		w.released[i] = now
	}

	if !justPressed {
		return false
	}

	return w.released[i].IsZero() || now.Sub(w.released[i]) >= chordRepeatGap
}
