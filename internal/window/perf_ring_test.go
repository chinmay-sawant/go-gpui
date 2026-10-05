package window

import (
	"testing"
	"time"
)

func TestPerfRingMath(t *testing.T) {
	t.Parallel()
	var r perfRing
	if r.avg() != 0 || r.p95() != 0 || r.p99() != 0 || r.longFrames(perfLongThreshold) != 0 {
		t.Fatal("empty ring must read zero")
	}
	for i := 0; i < 10; i++ {
		r.record(10 * time.Millisecond)
	}
	if got := r.avg(); got != 10*time.Millisecond {
		t.Fatalf("avg = %v, want 10ms", got)
	}
	if got := r.p95(); got != 10*time.Millisecond {
		t.Fatalf("p95 = %v, want 10ms", got)
	}
	if got := r.p99(); got != 10*time.Millisecond {
		t.Fatalf("p99 = %v, want 10ms", got)
	}
	r.record(100 * time.Millisecond)
	if got := r.longFrames(perfLongThreshold); got != 1 {
		t.Fatalf("long = %d, want 1", got)
	}
	if got := r.p95(); got <= 10*time.Millisecond {
		t.Fatalf("p95 = %v, want above 10ms", got)
	}
	for i := 0; i < 200; i++ {
		r.record(time.Millisecond)
	}
	if got := r.count(); got != perfRingCap {
		t.Fatalf("count = %d, want %d", got, perfRingCap)
	}
}
