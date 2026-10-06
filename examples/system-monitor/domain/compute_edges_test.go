package domain

import (
	"testing"
	"time"
)

// TestComputeCoreCountChange checks that a changed core count leaves every
// per-core value invalid for one sample.
func TestComputeCoreCountChange(t *testing.T) {
	prev := mk(time.Second, 100, 1000)
	prev.CPUCores = []CPUTimes{{Busy: 10, Total: 100}, {Busy: 20, Total: 100}}
	next := mk(2*time.Second, 200, 1200)
	next.CPUCores = []CPUTimes{
		{Busy: 10, Total: 100},
		{Busy: 20, Total: 100},
		{Busy: 30, Total: 100},
	}

	got := Compute(prev, next)

	if len(got.CorePercent) != 3 {
		t.Fatalf("cores = %d", len(got.CorePercent))
	}
	for i, c := range got.CorePercent {
		if c.Valid {
			t.Fatalf("core %d valid after count change", i)
		}
	}
}

// TestComputeGap checks a suspend between samples: the wall clock moves much
// further than the monotonic clock.
func TestComputeGap(t *testing.T) {
	prev := mk(time.Second, 100, 1000)
	next := mk(2*time.Second, 200, 1200)
	next.Stamp.At = prev.Stamp.At.Add(time.Hour)

	got := Compute(prev, next)

	if !got.Gap {
		t.Fatal("gap not detected")
	}
	if got.CPUPercent.Valid || got.Disks[0].ReadRate.Valid {
		t.Fatal("rates valid across a gap")
	}
}

// TestComputeHotplug checks that a new disk or interface has an invalid rate
// on its first sample while existing devices keep theirs.
func TestComputeHotplug(t *testing.T) {
	prev := mk(time.Second, 100, 1000)
	next := mk(2*time.Second, 200, 1200)
	next.Disks = append(next.Disks, Disk{Device: "sdb", ReadBytes: 10})
	next.Nets = append(next.Nets, Net{Name: "usb0", RXBytes: 10})

	got := Compute(prev, next)

	if !got.Disks[0].ReadRate.Valid {
		t.Fatal("existing disk rate invalid")
	}
	if got.Disks[1].ReadRate.Valid || got.Nets[1].RXRate.Valid {
		t.Fatal("new device has a rate")
	}
}
