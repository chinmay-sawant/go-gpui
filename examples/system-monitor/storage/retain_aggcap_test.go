package storage

import (
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// TestRetainTrimsAggregateCap checks the aggregate cap after a fold.
func TestRetainTrimsAggregateCap(t *testing.T) {
	oldRaw, oldAgg := MaxRawRows, MaxAggregateRows
	MaxRawRows, MaxAggregateRows = 1000, 5

	t.Cleanup(func() { MaxRawRows, MaxAggregateRows = oldRaw, oldAgg })

	st := openTest(t)
	ctx := t.Context()
	id := newTestSession(t, st)

	now := time.Unix(1_700_000_000, 0)
	base := bucketBase(now, 25*time.Hour)

	rows := make([]Row, 0, 12)

	for i := range 12 {
		rows = append(rows, rec(domain.MetricCPU, "", base.Add(time.Duration(i)*AggregateBucket), float64(i)))
	}
	if _, err := st.Append(ctx, id, rows); err != nil {
		t.Fatal(err)
	}

	res, err := st.Retain(ctx, now)
	if err != nil {
		t.Fatal(err)
	}
	if res.Folded != 12 || res.Trimmed != 7 {
		t.Fatalf("result = %+v", res)
	}

	agg, err := st.Aggregates(ctx, Query{SessionID: id, Limit: 100})
	if err != nil || len(agg) != 5 {
		t.Fatalf("aggregates = %d err = %v", len(agg), err)
	}
}
