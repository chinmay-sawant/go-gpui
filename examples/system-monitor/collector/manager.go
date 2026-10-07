package collector

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// Manager owns the sampling loops, the worker pool, and the published
// snapshots. It is safe for concurrent use. A manager must be closed to stop
// its goroutines.
type Manager struct {
	opts Options
	pool *Pool

	mu      sync.Mutex
	started bool
	closed  bool
	start   time.Time
	mode    Mode
	gen     uint64
	source  domain.Source
	live    domain.Source
	cap     domain.Capabilities

	reading  domain.Sample
	haveRead bool
	prev     domain.Sample

	procs     []domain.Process
	procStamp domain.Stamp
	procPrev  map[domain.ProcessIdentity]uint64
	haveProcs bool
	procSrc   string

	tracked     domain.ProcessIdentity
	detail      domain.ProcessDetail
	detailPrev  domain.ProcessDetail
	haveDetail  bool
	detailLoad  bool
	detailErr   string
	detailSrc   string
	detailStamp domain.Stamp

	runtime domain.Runtime
	rt      rtSample

	rings map[domain.SeriesKey]*domain.Ring
	seen  map[domain.SeriesKey]time.Time
	errs  map[string]Error
	sink  Sink

	samples      uint64
	sampleErrs   uint64
	procSamples  uint64
	procErrs     uint64
	detailCount  uint64
	detailErrs   uint64
	skipped      uint64
	lastSample   time.Time
	lastProcess  time.Time
	lastDetailAt time.Time

	summaryBusy atomic.Bool
	procBusy    atomic.Bool
	detailBusy  atomic.Bool
	cancel      context.CancelFunc
	wg          sync.WaitGroup
}

// rtSample keeps the previous self-CPU reading.
type rtSample struct {
	ok   bool
	cpu  float64
	mono time.Duration
}

// New returns a manager with the loops stopped. Call Start to sample.
func New(opts Options) *Manager {
	o := opts.withDefaults()
	m := &Manager{
		opts:  o,
		mode:  o.Mode,
		rings: make(map[domain.SeriesKey]*domain.Ring),
		errs:  make(map[string]Error),
	}
	m.pool = NewPool(o.Workers, o.Queue, o.Deadline)
	m.source = m.buildSource(o.Mode)
	m.cap = m.source.Capabilities()

	return m
}
