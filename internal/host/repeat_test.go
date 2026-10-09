package host

import (
	"testing"
	"time"
)

func TestRepeatSchedulerDelayRatesAndStop(t *testing.T) {
	start := time.Unix(1, 0)
	policy := RepeatPolicy{Delay: 350 * time.Millisecond, Rates: []RepeatRate{
		{After: 0, Interval: 120 * time.Millisecond},
		{After: 2 * time.Second, Interval: 80 * time.Millisecond},
		{After: 4 * time.Second, Interval: 50 * time.Millisecond},
	}}
	var s RepeatScheduler
	if !s.Start(policy, start) {
		t.Fatal("valid policy rejected")
	}
	checks := []struct {
		at  time.Duration
		due bool
	}{{349 * time.Millisecond, false}, {350 * time.Millisecond, true},
		{469 * time.Millisecond, false}, {470 * time.Millisecond, true},
		{2 * time.Second, true}, {2*time.Second + 79*time.Millisecond, false},
		{2*time.Second + 80*time.Millisecond, true}, {4 * time.Second, true},
		{4*time.Second + 49*time.Millisecond, false}, {4*time.Second + 50*time.Millisecond, true}}
	for _, check := range checks {
		if got := s.Tick(start.Add(check.at)); got != check.due {
			t.Fatalf("Tick(%v) = %v, want %v", check.at, got, check.due)
		}
	}
	s.Stop()
	if s.Tick(start.Add(5 * time.Second)) {
		t.Fatal("stopped scheduler repeated")
	}
}

func TestRepeatSchedulerRejectsInvalidRates(t *testing.T) {
	var s RepeatScheduler
	if s.Start(RepeatPolicy{Rates: []RepeatRate{{Interval: time.Second}, {After: 0, Interval: time.Second}}}, time.Now()) {
		t.Fatal("unordered policy accepted")
	}
}
