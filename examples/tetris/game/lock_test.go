package game

import (
	"testing"
	"time"
)

func TestLockDelayLocksAfter500ms(t *testing.T) {
	g := running(PieceT, Rot0, 3, 18)

	ev := g.Step(400 * time.Millisecond)
	if haveEvent(ev, EventLock) {
		t.Fatal("locked before the delay")
	}

	if g.Pieces != 0 {
		t.Fatalf("pieces = %d, want 0", g.Pieces)
	}

	ev = g.Step(100 * time.Millisecond)
	if !haveEvent(ev, EventLock) {
		t.Fatal("did not lock at the delay")
	}

	if g.Pieces != 1 {
		t.Fatalf("pieces = %d, want 1", g.Pieces)
	}
}

func TestRotationResetsTheLockDelay(t *testing.T) {
	g := running(PieceT, Rot0, 3, 18)

	g.Step(400 * time.Millisecond)

	if !g.rotate(true) {
		t.Fatal("the floor kick should fit")
	}

	g.Step(400 * time.Millisecond)

	if g.Pieces != 0 {
		t.Fatal("the rotation did not reset the lock timer")
	}

	g.Step(100 * time.Millisecond)

	if g.Pieces != 1 {
		t.Fatal("the piece never locked")
	}
}

func TestLockResetsAreCapped(t *testing.T) {
	g := running(PieceT, Rot0, 3, 18)

	for i := 0; i < maxLockResets; i++ {
		dx := -1
		if i%2 == 1 {
			dx = 1
		}

		if !g.move(dx, 0) {
			t.Fatalf("move %d blocked", i)
		}

		g.Step(400 * time.Millisecond)
	}

	if g.resets != maxLockResets {
		t.Fatalf("resets = %d, want %d", g.resets, maxLockResets)
	}

	if !g.move(-1, 0) {
		t.Fatal("the capped move should still shift")
	}

	g.Step(100 * time.Millisecond)

	if g.Pieces != 1 {
		t.Fatal("lock delay kept resetting past the cap")
	}
}

func TestGroundedPieceLocksWithoutSteps(t *testing.T) {
	g := running(PieceI, RotR, -2, 16)

	if g.canDrop() {
		t.Fatal("the fixture piece should rest on the well floor")
	}

	ev := g.Step(LockDelay)

	if !haveEvent(ev, EventLock) || g.Pieces != 1 {
		t.Fatal("a full lock delay did not lock the piece")
	}
}
