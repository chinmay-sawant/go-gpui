package domain

import (
	"testing"
	"time"
)

// base is a wall clock anchor so wall and monotonic clocks agree in tests.
var base = time.Unix(1_700_000_000, 0)

// mk builds a sample with monotonic stamp and given counters.
func mk(mono time.Duration, busy, total uint64) Sample {
	return Sample{
		Stamp:    Stamp{At: base.Add(mono), Mono: mono},
		CPUTotal: CPUTimes{Busy: busy, Total: total},
		Mem:      Memory{Total: Bytes(100), Available: Bytes(40)},
		Disks:    []Disk{{Device: "sda", ReadBytes: 100, WriteBytes: 200}},
		Nets:     []Net{{Name: "eth0", RXBytes: 100, TXBytes: 50}},
	}
}

// TestComputeRates checks a normal interval, including disk and net rates.
func TestComputeRates(t *testing.T) {
	prev := mk(time.Second, 100, 1000)
	next := mk(2*time.Second, 200, 1200)
	next.Disks[0].ReadBytes = 1100
	next.Nets[0].RXBytes = 1100

	got := Compute(prev, next)

	if !got.CPUPercent.Valid || got.CPUPercent.N != 50 {
		t.Fatalf("cpu = %+v, want 50", got.CPUPercent)
	}
	if !got.Mem.UsedPercent.Valid || got.Mem.Used.N != 60 {
		t.Fatalf("mem = %+v used %v", got.Mem.UsedPercent, got.Mem.Used)
	}
	if got.Disks[0].ReadRate.N != 1000 {
		t.Fatalf("disk rate = %v, want 1000", got.Disks[0].ReadRate)
	}
	if got.Nets[0].RXRate.N != 1000 {
		t.Fatalf("net rate = %v, want 1000", got.Nets[0].RXRate)
	}
}

// TestComputeZeroElapsed checks that two stamps at the same instant produce no
// rates.
func TestComputeZeroElapsed(t *testing.T) {
	prev := mk(time.Second, 100, 1000)
	next := mk(time.Second, 200, 1200)

	got := Compute(prev, next)

	if got.CPUPercent.Valid || got.Disks[0].ReadRate.Valid {
		t.Fatal("rates valid with zero elapsed")
	}
}

// TestComputeCounterReset keeps a backwards counter invalid.
func TestComputeCounterReset(t *testing.T) {
	prev := mk(time.Second, 200, 1200)
	next := mk(2*time.Second, 100, 1000)

	got := Compute(prev, next)

	if got.CPUPercent.Valid {
		t.Fatalf("cpu valid after reset: %+v", got.CPUPercent)
	}
}
