package wire

import (
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
	"github.com/chinmay-sawant/ownframe/examples/download-manager/scheduler"
)

func TestTrackSpeed(t *testing.T) {
	b := &Backend{speeds: map[string]speed{}}
	at := time.Now()

	// The first progress event is only a baseline.
	b.track(scheduler.Event{Kind: scheduler.EventProgress, Job: domain.Job{ID: "j"}, At: at})

	if b.speeds["j"].bytes != 0 {
		t.Fatal("the first sample set a rate")
	}

	b.track(scheduler.Event{
		Kind: scheduler.EventProgress,
		Job:  domain.Job{ID: "j", Done: 100},
		At:   at.Add(time.Second),
	})

	if got := b.speeds["j"].bytes; got != 100 {
		t.Fatalf("rate = %v, want 100 B/s", got)
	}

	// A gap longer than MaxSampleGap must not update the rate.
	b.track(scheduler.Event{
		Kind: scheduler.EventProgress,
		Job:  domain.Job{ID: "j", Done: 200},
		At:   at.Add(time.Second + domain.MaxSampleGap + time.Second),
	})

	if got := b.speeds["j"].bytes; got != 100 {
		t.Fatalf("rate after the gap = %v, want 100 B/s", got)
	}

	// A state event resets the baseline.
	b.track(scheduler.Event{
		Kind: scheduler.EventState,
		Job:  domain.Job{ID: "j", Done: 200, State: domain.StatePaused},
		At:   at.Add(time.Minute),
	})

	if got := b.speeds["j"].bytes; got != 0 {
		t.Fatalf("rate after a state event = %v, want 0", got)
	}
}
