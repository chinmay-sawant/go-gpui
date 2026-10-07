package collector

import (
	"runtime"
	"testing"
	"time"
)

// TestDummyStress checks the 10,000 process fixture: size, counters, and that
// reading it does not start goroutines per process.
func TestDummyStress(t *testing.T) {
	before := runtime.NumGoroutine()

	d := NewStress(11)
	start := time.Now()

	rows, err := d.Processes(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != StressProcesses {
		t.Fatalf("processes = %d", len(rows))
	}

	s, err := d.Sample(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	if s.Procs.N != StressProcesses {
		t.Fatalf("procs = %v", s.Procs)
	}

	if elapsed := time.Since(start); elapsed > 10*time.Second {
		t.Fatalf("stress read took %v", elapsed)
	}
	if after := runtime.NumGoroutine(); after > before+8 {
		t.Fatalf("goroutines before=%d after=%d", before, after)
	}
}

// TestStressIsLabeled checks that the stress source still reports itself as
// the dummy fixture.
func TestStressIsLabeled(t *testing.T) {
	d := NewStress(1)
	if d.Name() != "dummy" {
		t.Fatalf("name = %q", d.Name())
	}
	if caps := d.Capabilities(); !caps.Processes || caps.Handles {
		t.Fatalf("caps = %+v", caps)
	}
}
