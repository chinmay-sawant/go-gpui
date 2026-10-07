package storage

import (
	"testing"
	"time"
)

// openTest opens a file-backed store in a temp dir.
func openTest(t *testing.T) *Store {
	t.Helper()

	st, err := Open(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = st.Close() })

	return st
}

// newTestSession inserts a recording session.
func newTestSession(t *testing.T, st *Store) int64 {
	t.Helper()

	id, err := st.StartSession(t.Context(), Session{
		Name: "test", Source: "live", Mode: "live", Sampling: time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}

	return id
}

// bucketBase returns a time aligned to an aggregate bucket, offset back by
// back, so two rows one second apart always share a bucket.
func bucketBase(now time.Time, back time.Duration) time.Time {
	ns := now.Add(-back).UnixNano()
	bucket := AggregateBucket.Nanoseconds()

	return time.Unix(0, (ns/bucket)*bucket)
}
