package input

import (
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
)

func TestDefaultKeymapMapsAliases(t *testing.T) {
	km := DefaultKeymap()

	cases := map[string]game.Action{
		"arrowleft":  game.ActionLeft,
		"a":          game.ActionLeft,
		"arrowright": game.ActionRight,
		"d":          game.ActionRight,
		"arrowdown":  game.ActionSoftDrop,
		"s":          game.ActionSoftDrop,
		"space":      game.ActionHardDrop,
		"arrowup":    game.ActionRotateCW,
		"x":          game.ActionRotateCW,
		"z":          game.ActionRotateCCW,
		"p":          game.ActionPause,
		"r":          game.ActionRestart,
		"enter":      game.ActionStart,
	}

	for key, want := range cases {
		if got := km.actionFor(key); got != want {
			t.Fatalf("key %q maps to %v, want %v", key, got, want)
		}
	}

	if got := km.actionFor("f9"); got != game.ActionNone {
		t.Fatalf("f9 maps to %v", got)
	}
}

func TestFirstActionNamesThePrimaryKey(t *testing.T) {
	km := DefaultKeymap()

	if got := km.FirstAction(game.ActionLeft); got != "arrowleft" {
		t.Fatalf("left hint %q", got)
	}

	if got := km.FirstAction(game.ActionHardDrop); got != "space" {
		t.Fatalf("hard drop hint %q", got)
	}
}

func TestEmptyKeymapIgnoresEverything(t *testing.T) {
	tr := NewTracker(Keymap{})
	tr.KeyDown("arrowleft")

	if got := tr.Step(game.FixedStep); len(got) != 0 {
		t.Fatalf("an empty keymap emitted %v", got)
	}
}
