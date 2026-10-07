package storage

import (
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// TestAppendIgnoresDuplicates checks the unique key on session, metric,
// device, and timestamp.
func TestAppendIgnoresDuplicates(t *testing.T) {
	st := openTest(t)
	ctx := t.Context()
	id := newTestSession(t, st)

	base := time.Now()
	row := rec(domain.MetricCPU, "", base, 12)

	kept, err := st.Append(ctx, id, []Row{row})
	if err != nil || kept != 1 {
		t.Fatalf("kept = %d err = %v", kept, err)
	}

	kept, err = st.Append(ctx, id, []Row{row})
	if err != nil || kept != 0 {
		t.Fatalf("duplicate kept = %d err = %v", kept, err)
	}

	if n, err := st.CountRows(ctx, id); err != nil || n != 1 {
		t.Fatalf("count = %d err = %v", n, err)
	}
}
