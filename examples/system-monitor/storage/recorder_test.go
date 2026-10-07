package storage

import (
	"context"
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// TestRecorderDropsWhenFull checks the bounded queue: a blocked writer makes
// the queue drop and count instead of blocking the caller.
func TestRecorderDropsWhenFull(t *testing.T) {
	st := openTest(t)
	ctx := t.Context()
	id := newTestSession(t, st)

	// Hold the only connection so the recorder's write waits.
	tx, err := st.db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = tx.Rollback() }()

	writer := NewRecorder(st, id, 1)

	for i := range 50 {
		writer.Record([]Row{rec(domain.MetricCPU, "", time.Now().Add(time.Duration(i)), 1)})
	}

	if writer.Stats().Dropped == 0 {
		t.Fatal("nothing was dropped")
	}

	_ = tx.Rollback()
}

// TestRecorderReportsFailure checks that a write against a closed store is
// counted and returned by Close, never hidden.
func TestRecorderReportsFailure(t *testing.T) {
	st := openTest(t)
	id := newTestSession(t, st)

	rec := NewRecorder(st, id, 8)
	if err := st.Close(); err != nil {
		t.Fatal(err)
	}

	if !rec.RecordSample(fixtureSamples()[0]) {
		t.Fatal("sample refused")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := rec.Close(ctx); err == nil {
		t.Fatal("Close hid the failed write")
	}
	if stats := rec.Stats(); stats.Failed == 0 || stats.LastError == "" {
		t.Fatalf("stats = %+v", stats)
	}
}

// TestRecorderCloseDeadline is in recorder_deadline_test.go.
