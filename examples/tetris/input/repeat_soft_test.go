package input

import (
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
)

func TestHeldRotationRepeats(t *testing.T) {
	tr := NewTracker(DefaultKeymap())
	tr.KeyDown("x")

	if got := tr.Step(game.FixedStep); !has(got, game.ActionRotateCW) {
		t.Fatalf("the press did not rotate: %v", got)
	}

	if got := feed(tr, 200*time.Millisecond); count(got, game.ActionRotateCW) != 1 {
		t.Fatalf("held rotation emitted %v", got)
	}
}

func TestHeldSoftDropRepeatsFast(t *testing.T) {
	tr := NewTracker(DefaultKeymap())
	tr.KeyDown("s")

	if got := tr.Step(game.FixedStep); !has(got, game.ActionSoftDrop) {
		t.Fatalf("the press did not drop: %v", got)
	}

	if got := feed(tr, 20*time.Millisecond); len(got) != 0 {
		t.Fatalf("repeated before the soft drop delay: %v", got)
	}

	if got := feed(tr, 20*time.Millisecond); count(got, game.ActionSoftDrop) != 1 {
		t.Fatalf("soft drop repeat emitted %v", got)
	}
}

func TestOneShotActionsDoNotRepeat(t *testing.T) {
	tr := NewTracker(DefaultKeymap())
	tr.KeyDown("space")

	if got := feed(tr, 500*time.Millisecond); count(got, game.ActionHardDrop) != 1 {
		t.Fatalf("hard drop emitted %v", got)
	}
}
