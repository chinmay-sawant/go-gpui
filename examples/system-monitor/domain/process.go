package domain

import (
	"fmt"
	"time"
)

// ProcessIdentity names one run of a process. A PID alone is not stable,
// because operating systems reuse PIDs. Start carries the source's process
// start value (Linux start ticks, a Windows creation FILETIME, or a dummy
// counter) and the pair identifies the run. Two rows with the same PID and a
// different Start are different processes, so cached state keyed by identity
// never leaks between runs.
type ProcessIdentity struct {
	PID   int32
	Start uint64
}

// String renders the identity for logs and map keys.
func (id ProcessIdentity) String() string {
	return fmt.Sprintf("%d:%d", id.PID, id.Start)
}

// Process is one row of the process table. CPUTime is cumulative scheduler
// nanoseconds, and the manager turns consecutive reads into CPU percent. A
// source fills what it can and leaves the rest invalid.
type Process struct {
	ID        ProcessIdentity
	Name      string
	User      string
	State     string
	StartedAt time.Time
	CPUTime   uint64
	CPU       Value
	Memory    Value
	Threads   Value
	Handles   Value
	Priority  Value
}

// ProcessDetail is one process read in full for the detail view.
type ProcessDetail struct {
	Process
	Command    string
	Exe        string
	Cwd        string
	Parent     ProcessIdentity
	ParentName string
	Virtual    Value
	ReadBytes  Value
	WriteBytes Value
	ReadRate   Value
	WriteRate  Value
}
