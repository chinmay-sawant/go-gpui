package collector

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// TestDummyShape checks the fixture's advertised size, cores, and unavailable
// values.
func TestDummyShape(t *testing.T) {
	d := NewDummy(DummyOptions{Seed: 1, Stamp: fixedStamps()})

	rows, err := d.Processes(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != DefaultDummyProcesses {
		t.Fatalf("processes = %d", len(rows))
	}

	s, err := d.Sample(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	if len(s.CPUCores) != 8 {
		t.Fatalf("cores = %d", len(s.CPUCores))
	}
	if s.Procs.N != DefaultDummyProcesses {
		t.Fatalf("procs value = %v", s.Procs)
	}
	if s.Handles.Valid {
		t.Fatal("handle count should be unavailable")
	}

	valid := 0
	for _, temp := range s.Temps {
		if temp.Celsius.Valid {
			valid++
		}
	}
	if valid != 1 || len(s.Temps) != 3 {
		t.Fatalf("temps valid=%d total=%d", valid, len(s.Temps))
	}
}

// TestDummyContextCancel checks that a canceled context stops a read.
func TestDummyContextCancel(t *testing.T) {
	d := NewDummy(DummyOptions{Seed: 1})

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, err := d.Sample(ctx); err == nil {
		t.Fatal("Sample ignored a canceled context")
	}
	if _, err := d.Processes(ctx); err == nil {
		t.Fatal("Processes ignored a canceled context")
	}
	if _, err := d.Detail(ctx, domain.ProcessIdentity{}); err == nil {
		t.Fatal("Detail ignored a canceled context")
	}
}
