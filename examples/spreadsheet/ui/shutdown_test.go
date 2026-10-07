package ui

import (
	"testing"
	"time"
)

// blockingBackend holds Range until the gate opens.
type blockingBackend struct {
	*fakeBackend
	gate chan struct{}
}

func (b *blockingBackend) Range(id string, a Area) ([]Cell, error) {
	<-b.gate

	return b.fakeBackend.Range(id, a)
}

func TestCloseBudgetWhileWorkerBlocked(t *testing.T) {
	b := &blockingBackend{fakeBackend: newFake(), gate: make(chan struct{})}
	app, err := New(Options{Backend: b, Width: 800, Height: 600})
	if err != nil {
		t.Fatal(err)
	}

	app.work.post(job{kind: jobFetch, gen: 99, sheet: "s1", area: Area{0, 0, 0, 0}})
	time.Sleep(20 * time.Millisecond)

	start := time.Now()
	if app.work.close(50 * time.Millisecond) {
		t.Fatal("blocked worker reported a clean join")
	}

	if elapsed := time.Since(start); elapsed > time.Second {
		t.Fatalf("close budget ignored: %v", elapsed)
	}

	close(b.gate)
	select {
	case <-app.work.done:
	case <-time.After(2 * time.Second):
		t.Fatal("worker did not exit after the backend returned")
	}

	if err := app.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestCloseDuringPendingWork(t *testing.T) {
	b := newFake()
	app := newTestApp(t, b)
	settle(t, app)

	app.work.post(job{kind: jobApply, sheet: "s1", edits: []Edit{{Row: 0, Col: 0, Raw: "x"}}})
	if err := app.Close(); err != nil {
		t.Fatal(err)
	}

	if !b.closed {
		t.Fatal("backend not closed")
	}
}
