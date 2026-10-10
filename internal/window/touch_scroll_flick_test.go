package window

import (
	"testing"
	"time"
)

func TestTouchScrollFlickStartsAtReducedVelocity(t *testing.T) {
	start := time.Unix(0, 0)
	m := newTouchScrollMotion(1, 5)
	m.begin(start)
	m.move(0, -200, start.Add(50*time.Millisecond))
	m.release()
	var distance int
	for i := 1; i <= 30; i++ {
		_, dy := m.step(start.Add(50*time.Millisecond + time.Duration(i)*time.Second/60))
		distance += dy
	}
	if distance < 250 || distance > 500 {
		t.Fatalf("quick flick distance = %d, want 250..500px", distance)
	}
}

func TestTouchScrollPauseBeforeReleaseCancelsFling(t *testing.T) {
	start := time.Unix(0, 0)
	m := newTouchScrollMotion(1, 5)
	m.begin(start)
	m.move(0, -40, start.Add(20*time.Millisecond))
	m.idle(start.Add(70 * time.Millisecond))
	m.release()
	_, dy := m.step(start.Add(90 * time.Millisecond))
	if dy != 0 {
		t.Fatalf("fling after stationary pause = %d, want 0", dy)
	}
}
