package input

import (
	"time"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
)

// heldKey is one held repeating action.
type heldKey struct {
	timer time.Duration
	next  time.Duration
}

// press starts the action and its repeat timer.
func (t *Tracker) press(a game.Action) {
	switch a {
	case game.ActionLeft, game.ActionRight:
		t.setDir(a)
	case game.ActionSoftDrop:
		t.queue = append(t.queue, a)
		t.held[a] = &heldKey{next: SoftDropDelay}
	case game.ActionRotateCW, game.ActionRotateCCW:
		t.queue = append(t.queue, a)
		t.held[a] = &heldKey{next: DAS}
	default:
		t.queue = append(t.queue, a)
	}
}

// setDir makes the freshly pressed direction active and queues it.
func (t *Tracker) setDir(a game.Action) {
	if a == game.ActionLeft {
		t.left = true
	} else {
		t.right = true
	}

	if other := opposite(a); t.held[other] != nil {
		delete(t.held, other)
	}

	t.dir = a
	t.held[a] = &heldKey{next: DAS}
	t.queue = append(t.queue, a)
}

// releaseDir hands the active direction to the other held key, if any,
// and queues the resumed move at once.
func (t *Tracker) releaseDir(a game.Action) {
	delete(t.held, a)
	if a != t.dir {
		return
	}

	switch {
	case a == game.ActionLeft && t.right:
		t.dir = game.ActionRight
	case a == game.ActionRight && t.left:
		t.dir = game.ActionLeft
	default:
		t.dir = game.ActionNone

		return
	}

	t.held[t.dir] = &heldKey{next: DAS}
	t.queue = append(t.queue, t.dir)
}

// opposite returns the other horizontal direction.
func opposite(a game.Action) game.Action {
	if a == game.ActionLeft {
		return game.ActionRight
	}

	return game.ActionLeft
}
