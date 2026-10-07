package storage

import (
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// TestRetainClockBack checks that a wall clock step backwards pauses age
// deletion: a row stamped in the future survives.
func TestRetainClockBack(t *testing.T) {
	st := openTest(t)
	ctx := t.Context()
	id := newTestSession(t, st)

	now := time.Unix(1_700_000_000, 0)
	if _, err := st.Append(ctx, id, []Row{rec(domain.MetricCPU, "", now.Add(2*time.Hour), 7)}); err != nil {
		t.Fatal(err)
	}

	res, err := st.Retain(ctx, now)
	if err != nil {
		t.Fatal(err)
	}
	if res.Folded != 0 {
		t.Fatalf("folded = %d", res.Folded)
	}

	raw, err := st.History(ctx, Query{SessionID: id})
	if err != nil || len(raw) != 1 {
		t.Fatalf("raw = %+v err = %v", raw, err)
	}
}
