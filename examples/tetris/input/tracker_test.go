package input

import (
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
)

func TestLastDirectionWinsAndTheOtherResumes(t *testing.T) {
	tr := NewTracker(DefaultKeymap())
	tr.KeyDown("arrowleft")

	if got := tr.Step(game.FixedStep); !has(got, game.ActionLeft) {
		t.Fatalf("left press emitted %v", got)
	}

	tr.KeyDown("arrowright")

	got := tr.Step(game.FixedStep)
	if !has(got, game.ActionRight) || has(got, game.ActionLeft) {
		t.Fatalf("right press emitted %v", got)
	}

	tr.KeyUp("arrowright")

	got = tr.Step(game.FixedStep)
	if !has(got, game.ActionLeft) {
		t.Fatalf("left did not resume: %v", got)
	}

	tr.KeyUp("arrowleft")

	if got := tr.Step(game.FixedStep); len(got) != 0 {
		t.Fatalf("actions after both releases: %v", got)
	}
}

func TestRapidTapsEachMoveOnce(t *testing.T) {
	tr := NewTracker(DefaultKeymap())

	for i := 0; i < 3; i++ {
		tr.KeyDown("a")
		tr.KeyUp("a")
	}

	if got := tr.Step(game.FixedStep); count(got, game.ActionLeft) != 3 {
		t.Fatalf("three taps emitted %v", got)
	}
}

func TestRepeatedRotationPressesEachTurn(t *testing.T) {
	tr := NewTracker(DefaultKeymap())

	for i := 0; i < 4; i++ {
		tr.KeyDown("z")
		tr.KeyUp("z")
	}

	if got := tr.Step(game.FixedStep); count(got, game.ActionRotateCCW) != 4 {
		t.Fatalf("four presses emitted %v", got)
	}
}

func TestDuplicateKeyDownIsIgnored(t *testing.T) {
	tr := NewTracker(DefaultKeymap())
	tr.KeyDown("a")
	tr.KeyDown("a")

	if got := tr.Step(game.FixedStep); count(got, game.ActionLeft) != 1 {
		t.Fatalf("a double press emitted %v", got)
	}
}
