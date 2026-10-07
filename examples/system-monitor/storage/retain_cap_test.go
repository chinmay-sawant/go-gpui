package storage

import (
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// TestRetainTrimsRawCap checks that the raw table cap removes the oldest rows
// even when they are inside the retention window.
func TestRetainTrimsRawCap(t *testing.T) {
	oldRaw, oldAgg := MaxRawRows, MaxAggregateRows
	MaxRawRows, MaxAggregateRows = 10, 5

	t.Cleanup(func() { MaxRawRows, MaxAggregateRows = oldRaw, oldAgg })

	st := openTest(t)
	ctx := t.Context()
	id := newTestSession(t, st)

	now := time.Now()
	rows := make([]Row, 0, 25)

	for i := range 25 {
		rows = append(rows, rec(domain.MetricCPU, "", now.Add(-time.Minute).Add(time.Duration(i)*time.Second), float64(i)))
	}
	if _, err := st.Append(ctx, id, rows); err != nil {
		t.Fatal(err)
	}

	res, err := st.Retain(ctx, now)
	if err != nil {
		t.Fatal(err)
	}
	if res.Trimmed != 15 {
		t.Fatalf("trimmed = %d", res.Trimmed)
	}

	left, err := st.History(ctx, Query{SessionID: id, Limit: 100})
	if err != nil || len(left) != 10 {
		t.Fatalf("left = %d err = %v", len(left), err)
	}
	if left[0].Value != 15 {
		t.Fatalf("oldest kept = %v", left[0].Value)
	}
}

// TestRetainTrimsAggregateCap is in retain_aggcap_test.go.
