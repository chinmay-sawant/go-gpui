package storage

import (
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// rec builds one history row.
func rec(metric domain.Metric, device string, at time.Time, value float64) Row {
	return Row{Metric: metric, Device: device, At: at, Mono: time.Duration(at.UnixNano()), Value: value, Valid: true}
}

// TestSessionsRoundTrip checks start, end, list, and the row count.
func TestSessionsRoundTrip(t *testing.T) {
	st := openTest(t)
	ctx := t.Context()

	id := newTestSession(t, st)

	got, ok, err := st.Session(ctx, id)
	if err != nil || !ok || got.State != StateRecording || got.Name != "test" {
		t.Fatalf("session = %+v ok=%v err=%v", got, ok, err)
	}

	base := time.Now()
	if _, err := st.Append(ctx, id, []Row{rec(domain.MetricCPU, "", base, 12)}); err != nil {
		t.Fatal(err)
	}
	if err := st.EndSession(ctx, id, StateDone, "stopped"); err != nil {
		t.Fatal(err)
	}

	got, ok, err = st.Session(ctx, id)
	if err != nil || !ok || got.State != StateDone || got.Note != "stopped" || got.Ended.IsZero() {
		t.Fatalf("session = %+v ok=%v err=%v", got, ok, err)
	}
	if got.Rows != 1 {
		t.Fatalf("rows = %d", got.Rows)
	}

	all, err := st.Sessions(ctx, 10)
	if err != nil || len(all) != 1 || all[0].ID != id {
		t.Fatalf("sessions = %+v err=%v", all, err)
	}

	if _, ok, err := st.Session(ctx, id+999); err != nil || ok {
		t.Fatalf("missing session ok=%v err=%v", ok, err)
	}
	if err := st.EndSession(ctx, id+999, StateDone, ""); err == nil {
		t.Fatal("ending a missing session succeeded")
	}
}

// TestAppendIgnoresDuplicates is in append_test.go.
