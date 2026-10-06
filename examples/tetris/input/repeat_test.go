package input

import (
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
)

// feed runs d worth of fixed steps and collects the actions.
func feed(tr *Tracker, d time.Duration) []game.Action {
	var out []game.Action

	for n := int(d / game.FixedStep); n > 0; n-- {
		out = append(out, tr.Step(game.FixedStep)...)
	}

	return out
}

// count returns how often want appears.
func count(got []game.Action, want game.Action) int {
	n := 0

	for _, a := range got {
		if a == want {
			n++
		}
	}

	return n
}

// has reports whether want appears at all.
func has(got []game.Action, want game.Action) bool { return count(got, want) > 0 }

func TestHeldMoveRepeatsAfterDASThenARR(t *testing.T) {
	tr := NewTracker(DefaultKeymap())
	tr.KeyDown("arrowleft")

	if got := tr.Step(game.FixedStep); !has(got, game.ActionLeft) {
		t.Fatalf("the press did not fire: %v", got)
	}

	if got := feed(tr, 100*time.Millisecond); len(got) != 0 {
		t.Fatalf("repeated before DAS: %v", got)
	}

	if got := feed(tr, 100*time.Millisecond); count(got, game.ActionLeft) != 1 {
		t.Fatalf("crossing DAS emitted %v", got)
	}

	if got := feed(tr, ARR); count(got, game.ActionLeft) != 1 {
		t.Fatalf("the ARR repeat emitted %v", got)
	}
}
