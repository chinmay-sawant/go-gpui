package domain

import "time"

// Runtime is this Go process's own metrics, kept apart from the machine
// sample so monitor overhead never hides inside the system's numbers. Rates
// are invalid on the first sample of an era.
type Runtime struct {
	Stamp      Stamp
	Goroutines int
	HeapBytes  uint64
	HeapGoal   uint64
	SysBytes   uint64
	AllocBytes uint64
	GCCount    uint32
	GCPause    time.Duration
	CPUPercent Value
	Uptime     time.Duration
}

// Fixture is a labeled set of samples a fresh database can be seeded from.
// Storage writes one fixture per version and never mixes it with recorded
// sessions.
type Fixture struct {
	Version int
	Name    string
	Source  string
	Note    string
	Started time.Time
	Step    time.Duration
	Samples []Sample
}
