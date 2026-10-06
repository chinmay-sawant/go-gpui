package ui

import (
	"context"
	"testing"
	"time"
)

func TestPulseAnimatesRetainedOp(t *testing.T) {
	a := newHeadless(t)
	a.follow.SetFollow(true)
	ctx := context.Background()

	if err := a.draw(ctx); err != nil {
		t.Fatal(err)
	}

	a.pulse(time.Now())

	if a.pulseDot == nil {
		t.Fatal("live dot operation not found")
	}

	if alpha := a.pulseDot.Alpha; alpha <= 0.3 || alpha > 0.8 {
		t.Fatalf("pulse alpha = %v", alpha)
	}

	gen := a.pulseGen

	if err := a.draw(ctx); err != nil {
		t.Fatal(err)
	}

	a.pulse(time.Now())

	if a.pulseGen == gen {
		t.Fatal("pulse generation did not advance after a redraw")
	}

	if a.pulseDot == nil {
		t.Fatal("live dot operation not rebound after a redraw")
	}
}
