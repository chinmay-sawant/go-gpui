package input

import "github.com/chinmay-sawant/ownframe/examples/tetris/game"

// Tracker owns held-key state and repeat timers.
type Tracker struct {
	km    Keymap
	held  map[game.Action]*heldKey
	keys  map[string]game.Action
	dir   game.Action
	left  bool
	right bool
	queue []game.Action
}

// NewTracker returns a tracker using the keymap.
func NewTracker(km Keymap) *Tracker {
	return &Tracker{
		km:   km,
		held: map[game.Action]*heldKey{},
		keys: map[string]game.Action{},
	}
}

// KeyDown records a press and queues its first action.
func (t *Tracker) KeyDown(key string) {
	a := t.km.actionFor(key)
	if a == game.ActionNone {
		return
	}

	if _, ok := t.keys[key]; ok {
		return
	}

	t.keys[key] = a
	t.press(a)
}

// KeyUp records a release; the other held direction resumes.
func (t *Tracker) KeyUp(key string) {
	a, ok := t.keys[key]
	if !ok {
		return
	}

	delete(t.keys, key)

	switch a {
	case game.ActionLeft:
		t.left = false
		t.releaseDir(a)
	case game.ActionRight:
		t.right = false
		t.releaseDir(a)
	default:
		delete(t.held, a)
	}
}

// ReleaseAll clears every held key and timer, for focus loss.
func (t *Tracker) ReleaseAll() {
	t.keys = map[string]game.Action{}
	t.held = map[game.Action]*heldKey{}
	t.dir = game.ActionNone
	t.left, t.right = false, false
	t.queue = nil
}
