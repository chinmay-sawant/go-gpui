package input

import (
	"sort"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
)

// Step advances repeat timers and returns this step's actions.
func (t *Tracker) Step(step time.Duration) []game.Action {
	if step <= 0 {
		step = game.FixedStep
	}

	for _, a := range []game.Action{
		game.ActionLeft, game.ActionRight, game.ActionSoftDrop,
		game.ActionRotateCW, game.ActionRotateCCW,
	} {
		h := t.held[a]
		if h == nil || a != t.dir && (a == game.ActionLeft || a == game.ActionRight) {
			continue
		}

		h.timer += step
		for h.timer >= h.next {
			h.timer -= h.next
			h.next = ARR
			t.queue = append(t.queue, a)
		}
	}

	out := t.queue
	t.queue = nil

	return out
}

// Held lists the currently held key names, sorted.
func (t *Tracker) Held() []string {
	out := make([]string, 0, len(t.keys))
	for k := range t.keys {
		out = append(out, k)
	}

	sort.Strings(out)

	return out
}
