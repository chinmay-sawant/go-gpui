package domain

import (
	"testing"
	"time"
)

// TestComputeProcesses checks the identity matching: a new identity and a
// backwards counter stay invalid, and a reused PID with a new start value is
// not matched.
func TestComputeProcesses(t *testing.T) {
	id := ProcessIdentity{PID: 10, Start: 1}
	reused := ProcessIdentity{PID: 10, Start: 2}

	prev := map[ProcessIdentity]uint64{id: 1_000_000_000}
	rows := []Process{
		{ID: id, CPUTime: 2_000_000_000},
		{ID: reused, CPUTime: 5},
		{ID: ProcessIdentity{PID: 11, Start: 1}, CPUTime: 0},
	}

	ComputeProcesses(prev, rows, time.Second)

	if !rows[0].CPU.Valid || rows[0].CPU.N != 100 {
		t.Fatalf("row 0 cpu = %+v, want 100", rows[0].CPU)
	}
	if rows[1].CPU.Valid || rows[2].CPU.Valid {
		t.Fatal("new or zero rows got a cpu rate")
	}
}

// TestComputeProcessesBackwards checks that a counter reset is not turned into
// a negative rate.
func TestComputeProcessesBackwards(t *testing.T) {
	id := ProcessIdentity{PID: 10, Start: 1}
	rows := []Process{{ID: id, CPUTime: 5}}

	ComputeProcesses(map[ProcessIdentity]uint64{id: 100}, rows, time.Second)

	if rows[0].CPU.Valid {
		t.Fatal("backwards counter produced a rate")
	}
}

// TestComputeDetail checks IO rates and the identity guard.
func TestComputeDetail(t *testing.T) {
	id := ProcessIdentity{PID: 3, Start: 9}
	prev := ProcessDetail{Process: Process{ID: id}, ReadBytes: Bytes(1000), WriteBytes: Bytes(0)}
	next := ProcessDetail{Process: Process{ID: id}, ReadBytes: Bytes(2000), WriteBytes: Bytes(0)}

	got := ComputeDetail(prev, next, time.Second)
	if got.ReadRate.N != 1000 {
		t.Fatalf("read rate = %v", got.ReadRate)
	}

	next.ID = ProcessIdentity{PID: 3, Start: 10}
	got = ComputeDetail(prev, next, time.Second)
	if got.ReadRate.Valid {
		t.Fatal("rate computed across identities")
	}
}
