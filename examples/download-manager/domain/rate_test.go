package domain

import (
	"testing"
	"time"
)

// TestRate covers normal, first, stalled, and suspended samples.
func TestRate(t *testing.T) {
	base := time.Now()

	got, ok := Rate(0, base, 100, base.Add(time.Second), MaxSampleGap)
	if !ok || got != 100 {
		t.Errorf("normal sample = %v ok=%v", got, ok)
	}

	if _, ok := Rate(0, time.Time{}, 100, base, MaxSampleGap); ok {
		t.Error("first sample accepted")
	}

	if _, ok := Rate(100, base, 100, base, MaxSampleGap); ok {
		t.Error("non-advancing clock accepted")
	}

	if _, ok := Rate(100, base, 50, base.Add(time.Second), MaxSampleGap); ok {
		t.Error("negative progress accepted")
	}

	if _, ok := Rate(0, base, 100, base.Add(time.Hour), MaxSampleGap); ok {
		t.Error("suspended gap accepted")
	}

	if _, ok := Rate(0, base, 100, base.Add(time.Hour), 0); !ok {
		t.Error("unlimited gap refused")
	}
}
