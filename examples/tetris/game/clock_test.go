package game

import (
	"testing"
	"time"
)

func TestClockBoundsCatchUpAndDropsStalls(t *testing.T) {
	c := &Clock{}
	t0 := time.Unix(0, 0)

	if n := c.Advance(t0); n != 0 {
		t.Fatalf("the first advance ran %d steps", n)
	}

	if n := c.Advance(t0.Add(16 * time.Millisecond)); n != 0 {
		t.Fatalf("16ms ran %d steps", n)
	}

	if n := c.Advance(t0.Add(200 * time.Millisecond)); n != MaxCatchUp {
		t.Fatalf("a 184ms gap ran %d steps, want %d", n, MaxCatchUp)
	}

	if n := c.Advance(t0.Add(time.Second)); n != 0 {
		t.Fatalf("an 800ms stall ran %d steps instead of pausing", n)
	}

	c.Reset(t0.Add(2 * time.Second))

	if n := c.Advance(t0.Add(2*time.Second + FixedStep)); n != 1 {
		t.Fatalf("after reset a single step gap ran %d steps", n)
	}
}

func TestClockResetClearsTheAccumulator(t *testing.T) {
	c := &Clock{}
	t0 := time.Unix(0, 0)
	c.Advance(t0)
	c.Advance(t0.Add(16 * time.Millisecond))
	c.Reset(t0.Add(20 * time.Millisecond))

	if n := c.Advance(t0.Add(20 * time.Millisecond)); n != 0 {
		t.Fatalf("reset left %d steps of debt", n)
	}
}
