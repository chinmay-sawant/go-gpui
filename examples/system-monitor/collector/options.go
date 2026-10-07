package collector

import (
	"time"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// Options configures a Manager. Zero fields take the defaults in defaults.go.
type Options struct {
	Mode            Mode
	Seed            int64
	Stress          bool
	SummaryInterval time.Duration
	ProcessInterval time.Duration
	DetailInterval  time.Duration
	Deadline        time.Duration
	Capacity        int
	Workers         int
	Queue           int
	History         int
	HistoryStep     time.Duration
	// Source overrides the built-in source. Tests use it to drive the
	// manager with a fake; nil chooses dummy or live by Mode.
	Source domain.Source
}

// withDefaults returns a copy of o with zero fields filled in.
func (o Options) withDefaults() Options {
	if o.SummaryInterval <= 0 {
		o.SummaryInterval = DefaultSummaryInterval
	}
	if o.ProcessInterval <= 0 {
		o.ProcessInterval = DefaultProcessInterval
	}
	if o.DetailInterval <= 0 {
		o.DetailInterval = o.ProcessInterval
	}
	if o.Deadline <= 0 {
		o.Deadline = DefaultDeadline
	}
	if o.Capacity <= 0 {
		o.Capacity = DefaultCapacity
	}
	if o.Workers <= 0 {
		o.Workers = DefaultWorkers
	}
	if o.Queue <= 0 {
		o.Queue = DefaultQueue
	}
	if o.History < 0 {
		o.History = 0
	}
	if o.History == 0 {
		o.History = DefaultHistory
	}
	if o.HistoryStep <= 0 {
		o.HistoryStep = DefaultHistoryStep
	}
	if o.Mode == "" {
		o.Mode = ModeDummy
	}

	return o
}

// processCount returns how many processes the dummy fixture builds.
func (o Options) processCount() int {
	if o.Stress {
		return StressProcesses
	}

	return DefaultDummyProcesses
}
