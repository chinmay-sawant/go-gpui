package window

import (
	"math"
	"testing"
	"time"
)

func TestTouchScrollMotionFollowsFingerAndFlingDecays(t *testing.T) {
	start := time.Unix(0, 0)
	m := newTouchScrollMotion(1, 5)
	m.begin(start)
	dx, dy := m.move(0, -20, start.Add(20*time.Millisecond))
	if dx != 0 || dy != 20 {
		t.Fatalf("drag delta = %d, %d, want 0, 20", dx, dy)
	}
	m.release()
	_, first := m.step(start.Add(70 * time.Millisecond))
	_, second := m.step(start.Add(120 * time.Millisecond))
	if first <= 0 || second <= 0 || second >= first {
		t.Fatalf("fling deltas = %d, %d, want positive deceleration", first, second)
	}
}

func TestTouchScrollFlingIsUpdateRateIndependent(t *testing.T) {
	if a, b := flingDistance(60), flingDistance(90); math.Abs(a-b) > 2 {
		t.Fatalf("60 Hz distance %.2f, 90 Hz distance %.2f", a, b)
	}
}

func TestTouchScrollSensitivityScalesFingerMovement(t *testing.T) {
	start := time.Unix(0, 0)
	m := newTouchScrollMotion(.5, 5)
	m.begin(start)
	_, dy := m.move(0, -10, start.Add(20*time.Millisecond))
	if dy != 5 {
		t.Fatalf("scroll delta = %d, want 5", dy)
	}
}

func flingDistance(rate int) float64 {
	start := time.Unix(0, 0)
	m := newTouchScrollMotion(1, 5)
	m.begin(start)
	m.move(0, -200, start.Add(50*time.Millisecond))
	m.release()
	var distance float64
	for i := 1; i <= rate/2; i++ {
		_, dy := m.step(start.Add(50*time.Millisecond + time.Duration(i)*time.Second/time.Duration(rate)))
		distance += float64(dy)
	}
	return distance
}
