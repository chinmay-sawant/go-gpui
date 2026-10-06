package scheduler

import (
	"context"
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
)

// TestProgressCoalescesAndStatesSurvive keeps one progress event per job
// and every state event in order.
func TestProgressCoalescesAndStatesSurvive(t *testing.T) {
	eng, _ := newTestEngine(t, &burstTransport{steps: 500}, Options{Workers: 1})
	eng.Start(context.Background())

	job := add(t, eng, t.TempDir(), "https://example.invalid/burst.bin")
	waitLive(t, eng, job.ID, domain.StateCompleted)

	var events []Event

	deadline := time.Now().Add(2 * time.Second)

	for time.Now().Before(deadline) {
		events = append(events, eng.Drain(1024)...)
		if hasState(events, job.ID, domain.StateCompleted) {
			break
		}

		time.Sleep(time.Millisecond)
	}

	progress, states := 0, []domain.State{}

	for _, ev := range events {
		if ev.Job.ID != job.ID {
			continue
		}

		switch ev.Kind {
		case EventProgress:
			progress++
		case EventState:
			states = append(states, ev.Job.State)
		}
	}

	if progress != 1 {
		t.Errorf("%d progress events survived, want 1", progress)
	}

	want := []domain.State{domain.StateQueued, domain.StateRunning, domain.StateCompleted}
	if len(states) != len(want) {
		t.Fatalf("state events %v, want %v", states, want)
	}

	for i, state := range want {
		if states[i] != state {
			t.Fatalf("state events %v, want %v", states, want)
		}
	}
}

// hasState reports whether the events hold a state event for id.
func hasState(events []Event, id string, state domain.State) bool {
	for _, ev := range events {
		if ev.Kind == EventState && ev.Job.ID == id && ev.Job.State == state {
			return true
		}
	}

	return false
}
