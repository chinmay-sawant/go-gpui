package storage

import (
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// TestHistoryKeepsInvalid checks that an unavailable reading is stored, not
// dropped.
func TestHistoryKeepsInvalid(t *testing.T) {
	st := openTest(t)
	ctx := t.Context()
	id := newTestSession(t, st)

	row := rec(domain.MetricTemp, "gpu", time.Now(), 0)
	row.Valid = false

	if _, err := st.Append(ctx, id, []Row{row}); err != nil {
		t.Fatal(err)
	}

	got, err := st.History(ctx, Query{SessionID: id, Metric: domain.MetricTemp})
	if err != nil || len(got) != 1 || got[0].Valid {
		t.Fatalf("rows = %+v err = %v", got, err)
	}
}
