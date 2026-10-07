package collector

import (
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// TestDummyChurn checks that processes exit and that a reused PID comes back
// with a different start value.
func TestDummyChurn(t *testing.T) {
	d := NewDummy(DummyOptions{Seed: 3, Stamp: fixedStamps()})

	first := make(map[domain.ProcessIdentity]bool)
	rows, err := d.Processes(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	for _, row := range rows {
		first[row.ID] = true
	}

	seen := make(map[int32]map[uint64]bool, len(rows))

	for range 40 {
		if _, err := d.Sample(t.Context()); err != nil {
			t.Fatal(err)
		}
		if _, err := d.Processes(t.Context()); err != nil {
			t.Fatal(err)
		}

		rows, err := d.Processes(t.Context())
		if err != nil {
			t.Fatal(err)
		}
		for _, row := range rows {
			if seen[row.ID.PID] == nil {
				seen[row.ID.PID] = make(map[uint64]bool)
			}
			seen[row.ID.PID][row.ID.Start] = true
		}
	}

	rows, err = d.Processes(t.Context())
	if err != nil {
		t.Fatal(err)
	}

	gone, reused := 0, 0
	for _, row := range rows {
		if !first[row.ID] {
			gone++
		}
	}
	for _, starts := range seen {
		if len(starts) > 1 {
			reused++
		}
	}

	if gone == 0 {
		t.Fatal("no process exited over 40 ticks")
	}
	if reused == 0 {
		t.Fatal("no PID was reused with a new start value")
	}
}
