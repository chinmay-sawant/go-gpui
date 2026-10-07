package storage

import (
	"context"
	"testing"
	"time"
)

// TestRecorderWrites checks the happy path: samples are written and counted.
func TestRecorderWrites(t *testing.T) {
	st := openTest(t)
	id := newTestSession(t, st)

	rec := NewRecorder(st, id, 8)
	if !rec.RecordSample(fixtureSamples()[0]) {
		t.Fatal("sample refused")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := rec.Close(ctx); err != nil {
		t.Fatal(err)
	}
	if stats := rec.Stats(); stats.Written == 0 || stats.Failed != 0 {
		t.Fatalf("stats = %+v", stats)
	}
	if n, err := st.CountRows(t.Context(), id); err != nil || n == 0 {
		t.Fatalf("rows = %d err = %v", n, err)
	}
}
