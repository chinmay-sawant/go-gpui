package storage

import (
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// TestRetainFoldsOldRows checks that old raw rows move into aggregates and
// fresh rows stay.
func TestRetainFoldsOldRows(t *testing.T) {
	st := openTest(t)
	ctx := t.Context()
	id := newTestSession(t, st)

	now := time.Unix(1_700_000_000, 0)
	old := bucketBase(now, 25*time.Hour)

	rows := []Row{
		rec(domain.MetricCPU, "", old, 10),
		rec(domain.MetricCPU, "", old.Add(time.Second), 20),
		rec(domain.MetricCPU, "", now.Add(-time.Minute), 30),
	}
	if _, err := st.Append(ctx, id, rows); err != nil {
		t.Fatal(err)
	}

	res, err := st.Retain(ctx, now)
	if err != nil {
		t.Fatal(err)
	}
	if res.Folded != 2 {
		t.Fatalf("folded = %d", res.Folded)
	}

	raw, err := st.History(ctx, Query{SessionID: id, Metric: domain.MetricCPU})
	if err != nil || len(raw) != 1 || raw[0].Value != 30 {
		t.Fatalf("raw = %+v err = %v", raw, err)
	}

	agg, err := st.Aggregates(ctx, Query{SessionID: id, Metric: domain.MetricCPU})
	if err != nil || len(agg) != 1 {
		t.Fatalf("aggregates = %+v err = %v", agg, err)
	}
	if agg[0].Count != 2 || agg[0].Min != 10 || agg[0].Max != 20 || agg[0].Avg != 15 {
		t.Fatalf("aggregate = %+v", agg[0])
	}
}

// TestRetainDropsOldAggregates checks the aggregate retention window.
func TestRetainDropsOldAggregates(t *testing.T) {
	st := openTest(t)
	ctx := t.Context()
	id := newTestSession(t, st)

	now := time.Unix(1_700_000_000, 0)
	ancient := bucketBase(now, 31*24*time.Hour)

	if _, err := st.Append(ctx, id, []Row{rec(domain.MetricCPU, "", ancient, 5)}); err != nil {
		t.Fatal(err)
	}

	res, err := st.Retain(ctx, now)
	if err != nil {
		t.Fatal(err)
	}
	if res.Folded != 1 || res.Aggregates != 1 {
		t.Fatalf("result = %+v", res)
	}

	agg, err := st.Aggregates(ctx, Query{SessionID: id})
	if err != nil || len(agg) != 0 {
		t.Fatalf("aggregates = %+v err = %v", agg, err)
	}
}

// TestRetainClockBack is in retain_clock_test.go.
