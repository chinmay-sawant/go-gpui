package domain

import (
	"testing"
	"time"
)

// TestElapsed checks monotonic distance and its guards.
func TestElapsed(t *testing.T) {
	a := Stamp{Mono: time.Second}
	b := Stamp{Mono: 3 * time.Second}

	if d, ok := Elapsed(a, b); !ok || d != 2*time.Second {
		t.Fatalf("elapsed = %v %v", d, ok)
	}
	if _, ok := Elapsed(b, a); ok {
		t.Fatal("backwards elapsed accepted")
	}
	if _, ok := Elapsed(a, a); ok {
		t.Fatal("same instant accepted")
	}
}

// TestGapSuspend checks that a wall clock that jumps while the monotonic clock
// does not is a gap.
func TestGapSuspend(t *testing.T) {
	prev := Stamp{At: base, Mono: time.Second}
	next := Stamp{At: base.Add(time.Hour), Mono: 2 * time.Second}

	if !Gap(prev, next, 0) {
		t.Fatal("suspend not detected")
	}
}

// TestGapClockStep checks both directions of a wall clock step.
func TestGapClockStep(t *testing.T) {
	prev := Stamp{At: base, Mono: time.Second}

	forward := Stamp{At: base.Add(time.Minute), Mono: 2 * time.Second}
	if !Gap(prev, forward, 0) {
		t.Fatal("forward step not detected")
	}

	back := Stamp{At: base.Add(-time.Minute), Mono: 2 * time.Second}
	if !Gap(prev, back, 0) {
		t.Fatal("backwards step not detected")
	}
}

// TestGapNormal checks that ordinary drift is not a gap.
func TestGapNormal(t *testing.T) {
	prev := Stamp{At: base, Mono: time.Second}
	next := Stamp{At: base.Add(2 * time.Second), Mono: 3 * time.Second}

	if Gap(prev, next, 0) {
		t.Fatal("ordinary interval marked as a gap")
	}
}
