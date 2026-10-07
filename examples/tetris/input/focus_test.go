package input

import (
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
)

func TestReleaseAllClearsHeldInputs(t *testing.T) {
	tr := NewTracker(DefaultKeymap())
	tr.KeyDown("arrowleft")
	tr.KeyDown("space")
	tr.Step(game.FixedStep)

	tr.ReleaseAll()

	if len(tr.Held()) != 0 {
		t.Fatalf("keys still held: %v", tr.Held())
	}

	if got := tr.Step(game.FixedStep); len(got) != 0 {
		t.Fatalf("stale actions after focus loss: %v", got)
	}

	tr.KeyDown("arrowright")

	if got := tr.Step(game.FixedStep); !has(got, game.ActionRight) {
		t.Fatalf("input stopped working after ReleaseAll: %v", got)
	}
}

func TestUnknownKeysAreIgnored(t *testing.T) {
	tr := NewTracker(DefaultKeymap())
	tr.KeyDown("f9")

	if got := tr.Step(game.FixedStep); len(got) != 0 {
		t.Fatalf("an unmapped key emitted %v", got)
	}
}

func TestReleaseAllDropsQueuedTaps(t *testing.T) {
	tr := NewTracker(DefaultKeymap())
	tr.KeyDown("a")

	tr.ReleaseAll()

	if got := tr.Step(game.FixedStep); len(got) != 0 {
		t.Fatalf("a queued tap survived ReleaseAll: %v", got)
	}
}
