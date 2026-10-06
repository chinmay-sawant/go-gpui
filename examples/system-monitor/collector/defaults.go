package collector

import "time"

// Defaults for Options. The summary runs about once a second; the process
// table is expensive, so it runs less often.
const (
	DefaultSummaryInterval = time.Second
	DefaultProcessInterval = 3 * time.Second
	DefaultDeadline        = 2 * time.Second
	DefaultCapacity        = 600
	DefaultWorkers         = 4
	DefaultQueue           = 8
	DefaultHistory         = 120
	DefaultHistoryStep     = 15 * time.Second
	DefaultDummyProcesses  = 250
	StressProcesses        = 10000
)
