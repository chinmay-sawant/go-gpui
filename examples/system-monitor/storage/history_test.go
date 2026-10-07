package storage

import (
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// TestHistoryFilters checks filters, limits, and keyset paging.
func TestHistoryFilters(t *testing.T) {
	st := openTest(t)
	ctx := t.Context()
	id := newTestSession(t, st)

	base := time.Unix(1_700_000_000, 0)
	rows := []Row{
		rec(domain.MetricCPU, "", base, 10),
		rec(domain.MetricCPU, "", base.Add(time.Second), 20),
		rec(domain.MetricCPU, "", base.Add(2*time.Second), 30),
		rec(domain.MetricMemUsed, "", base.Add(time.Second), 40),
		rec(domain.MetricNetRX, "eth0", base.Add(time.Second), 50),
		rec(domain.MetricNetRX, "wlan0", base.Add(time.Second), 60),
	}
	if kept, err := st.Append(ctx, id, rows); err != nil || kept != 6 {
		t.Fatalf("kept = %d err = %v", kept, err)
	}

	cpu, err := st.History(ctx, Query{SessionID: id, Metric: domain.MetricCPU})
	if err != nil || len(cpu) != 3 {
		t.Fatalf("cpu = %+v err = %v", cpu, err)
	}
	if cpu[0].Value != 10 || cpu[2].Value != 30 || cpu[0].At.UnixNano() != base.UnixNano() {
		t.Fatalf("cpu rows = %+v", cpu)
	}

	eth, err := st.History(ctx, Query{SessionID: id, Metric: domain.MetricNetRX, Device: "eth0"})
	if err != nil || len(eth) != 1 || eth[0].Value != 50 {
		t.Fatalf("eth0 = %+v err = %v", eth, err)
	}

	windowed, err := st.History(ctx, Query{
		SessionID: id,
		Since:     base.Add(time.Second),
		Until:     base.Add(time.Second),
	})
	if err != nil || len(windowed) != 4 {
		t.Fatalf("window = %d err = %v", len(windowed), err)
	}

	page, err := st.History(ctx, Query{SessionID: id, Limit: 2})
	if err != nil || len(page) != 2 {
		t.Fatalf("page = %d err = %v", len(page), err)
	}

	next, err := st.History(ctx, Query{SessionID: id, AfterID: page[1].ID, Limit: 2})
	if err != nil || len(next) != 2 || next[0].ID <= page[1].ID {
		t.Fatalf("next = %+v err = %v", next, err)
	}

	if _, err := st.History(ctx, Query{}); err == nil {
		t.Fatal("query without a session accepted")
	}
}
