package remote

import (
	"testing"
	"time"
)

func TestRepeatableControlsAndProfile(t *testing.T) {
	for _, id := range []string{"volup", "voldn", "chup", "chdn", "up", "down", "left", "right"} {
		if !repeatable(id) {
			t.Errorf("%q should repeat", id)
		}
	}
	for _, id := range []string{"mute", "power", "ok", "settings"} {
		if repeatable(id) {
			t.Errorf("%q should not repeat", id)
		}
	}
	if remoteRepeatPolicy.Delay != 350*time.Millisecond {
		t.Fatalf("initial delay = %v", remoteRepeatPolicy.Delay)
	}
	if len(remoteRepeatPolicy.Rates) != 3 {
		t.Fatalf("repeat rates = %d, want 3", len(remoteRepeatPolicy.Rates))
	}
	want := []time.Duration{120 * time.Millisecond, 80 * time.Millisecond, 50 * time.Millisecond}
	for i, rate := range remoteRepeatPolicy.Rates {
		if rate.Interval != want[i] {
			t.Errorf("rate %d interval = %v, want %v", i, rate.Interval, want[i])
		}
	}
	if remoteRepeatPolicy.Rates[0].After != 0 || remoteRepeatPolicy.Rates[1].After != 2*time.Second || remoteRepeatPolicy.Rates[2].After != 4*time.Second {
		t.Fatal("repeat acceleration thresholds changed")
	}
}
