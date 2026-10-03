package window

import "github.com/hajimehoshi/ebiten/v2"

// keyCount is one slot per key Ebiten can report.
const keyCount = int(ebiten.KeyMax) + 1

// keyGuardFrames is how many frames a key must be seen up before another
// press counts, and how many frames it must be held before its release
// counts. A server that reports auto-repeat as a release and a press leaves
// only one frame between the two, so the guard drops the pair while a real
// tap still reports both events.
const keyGuardFrames = 2

// keyWatch reports one press and one release per real key event, not per
// auto-repeat pulse.
type keyWatch struct {
	down [keyCount]int
	up   [keyCount]int
}

// newKeyWatch returns a watch whose keys start up, so the first press fires.
func newKeyWatch() keyWatch {
	var w keyWatch

	for i := range w.up {
		w.up[i] = keyGuardFrames
	}

	return w
}

// step advances one key by one frame and reports the press and the release
// a page should see in this frame.
func (w *keyWatch) step(key ebiten.Key, pressed bool) (down, up bool) {
	i := int(key)
	if i < 0 || i >= keyCount {
		return false, false
	}

	if pressed {
		fire := w.down[i] == 0 && w.up[i] >= keyGuardFrames
		w.up[i] = 0
		w.down[i]++

		return fire, false
	}

	fire := w.down[i] >= keyGuardFrames
	w.down[i] = 0

	if w.up[i] < keyGuardFrames {
		w.up[i]++
	}

	return false, fire
}
