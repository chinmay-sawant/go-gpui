package ui

import (
	"fmt"
	"time"
)

// fakeProc builds a process with a stable identity.
func fakeProc(pid int, name string, cpu float64) Process {
	return Process{
		ID:       fmt.Sprintf("%d:%d", pid, pid*10),
		PID:      pid,
		Name:     name,
		CPU:      cpu,
		CPUKnown: true,
	}
}

// snapshotOf builds a snapshot from the given processes.
func snapshotOf(procs ...Process) ProcSnapshot {
	return ProcSnapshot{At: time.Unix(1700000000, 0), Procs: procs}
}

// manyProcs returns n processes with descending CPU.
func manyProcs(n int) []Process {
	out := make([]Process, n)

	for i := range out {
		out[i] = fakeProc(100+i, fmt.Sprintf("proc-%03d", i), float64(n-i))
	}

	return out
}
