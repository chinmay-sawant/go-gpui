package storage

import (
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/collector"
	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// coreFlow is the headless acceptance path: seed a fresh
// directory, record a dummy manager, run retention, and close. It returns the
// directory and session ID.
func coreFlow(t *testing.T) (string, int64) {
	t.Helper()

	dir := t.TempDir()
	ctx := t.Context()

	st, err := Open(dir)
	if err != nil {
		t.Fatal(err)
	}

	wrote, err := st.Seed(ctx, collector.DummyFixture(1, time.Now(), 120, 15*time.Second))
	if err != nil || !wrote {
		t.Fatalf("seed wrote=%v err=%v", wrote, err)
	}

	id, err := st.StartSession(ctx, Session{
		Name: "core flow", Source: "dummy", Mode: "dummy", Sampling: 20 * time.Millisecond,
	})
	if err != nil {
		t.Fatal(err)
	}

	rec := NewRecorder(st, id, 64)

	m := collector.New(collector.Options{
		Mode:            collector.ModeDummy,
		SummaryInterval: 20 * time.Millisecond,
		ProcessInterval: time.Hour,
		Deadline:        time.Second,
	})
	m.SetSink(rec)

	if err := m.Start(ctx); err != nil {
		t.Fatal(err)
	}

	deadline := time.Now().Add(5 * time.Second)

	var rows []Point

	for time.Now().Before(deadline) {
		rows, err = st.History(ctx, Query{SessionID: id, Metric: domain.MetricCPU, Limit: 100})
		if err != nil {
			t.Fatal(err)
		}
		if len(rows) >= 5 {
			break
		}

		time.Sleep(20 * time.Millisecond)
	}

	m.SetSink(nil)
	if err := m.Close(ctx); err != nil {
		t.Fatal(err)
	}
	if err := rec.Close(ctx); err != nil {
		t.Fatal(err)
	}
	if err := st.EndSession(ctx, id, StateDone, "core flow"); err != nil {
		t.Fatal(err)
	}

	if stats := rec.Stats(); stats.Written < 5 || stats.Failed != 0 {
		t.Fatalf("recorder = %+v", stats)
	}
	if len(rows) < 5 {
		t.Fatalf("cpu rows while recording = %d", len(rows))
	}

	if _, err := st.Retain(ctx, time.Now()); err != nil {
		t.Fatal(err)
	}

	if err := st.Close(); err != nil {
		t.Fatal(err)
	}

	return dir, id
}
