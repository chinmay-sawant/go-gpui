package scheduler

import (
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
)

// waitState drains events until the job reaches want.
func waitState(t *testing.T, eng *Engine, id string, want domain.State) domain.Job {
	t.Helper()

	deadline := time.Now().Add(3 * time.Second)

	for time.Now().Before(deadline) {
		for _, ev := range eng.Drain(256) {
			if ev.Kind == EventState && ev.Job.ID == id && ev.Job.State == want {
				return ev.Job
			}
		}

		time.Sleep(5 * time.Millisecond)
	}

	t.Fatalf("job %s never reached %s", id, want)

	return domain.Job{}
}

// waitLive polls the in-memory job without draining events.
func waitLive(t *testing.T, eng *Engine, id string, want domain.State) domain.Job {
	t.Helper()

	deadline := time.Now().Add(3 * time.Second)

	for time.Now().Before(deadline) {
		if job, ok := eng.Job(id); ok && job.State == want {
			return job
		}

		time.Sleep(2 * time.Millisecond)
	}

	t.Fatalf("job %s never reached %s", id, want)

	return domain.Job{}
}

// waitStarted blocks until the gate transport begins a download.
func waitStarted(t *testing.T, g *gateTransport, id string) {
	t.Helper()

	select {
	case got := <-g.started:
		if got != id {
			t.Fatalf("started %s, want %s", got, id)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("download never started")
	}
}
