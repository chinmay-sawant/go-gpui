package ui

import "time"

// Process is one row of the process table. ID is the stable identity (PID
// plus start value), so PID reuse never merges two runs. The Known flags
// separate "0" from "denied".
type Process struct {
	ID           string
	PID          int
	Name         string
	State        string
	CPU          float64
	CPUKnown     bool
	Mem          uint64
	MemKnown     bool
	Start        time.Time
	Threads      int
	ThreadsKnown bool
}

// ProcSnapshot is one immutable view of every process at a timestamp.
type ProcSnapshot struct {
	At    time.Time
	Procs []Process
}

// ProcDetail is one process read in full. The Known flags mark values the
// source denied or could not read.
type ProcDetail struct {
	Process
	Command     string
	Exe         string
	Cwd         string
	Parent      string
	Virtual     uint64
	VirtualOK   bool
	Handles     uint64
	HandlesOK   bool
	Priority    int64
	PriorityOK  bool
	ReadRate    float64
	ReadRateOK  bool
	WriteRate   float64
	WriteRateOK bool
}

// Tracked is the state of the detailed process. ID names the identity the
// state belongs to; a result for a different identity is stale.
type Tracked struct {
	ID      string
	Detail  ProcDetail
	OK      bool
	Loading bool
	Err     string
}
