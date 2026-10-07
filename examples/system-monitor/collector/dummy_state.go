package collector

import (
	"time"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// init builds the fixture's starting state. It runs once per source, so the
// same seed always starts from the same counters.
func (d *Dummy) init() {
	d.cores = 8
	d.cpu = make([]domain.CPUTimes, d.cores)
	for i := range d.cpu {
		d.cpu[i] = domain.CPUTimes{Busy: 100 * uint64(i+3), Total: 100_000}
	}

	d.load = [3]float64{1.1, 1.4, 1.6}
	d.memTotal = 32 << 30
	d.memUsed = 12 << 30
	d.swapUsed = 512 << 20
	d.disks = []dmDisk{
		{name: "sda", mount: "/", total: 512 << 30, free: 220 << 30},
		{name: "nvme0n1", mount: "/home", total: 1 << 40, free: 700 << 30},
	}
	d.nets = []dmNet{{name: "eth0", up: true}, {name: "wlan0", up: true}}

	d.nextPID = 100
	d.nextStart = 1000

	for range d.opts.Processes {
		d.spawn()
	}
}

// advance moves the simulation to stamp and returns the machine sample.
// Callers hold d.mu.
func (d *Dummy) advance(stamp domain.Stamp) domain.Sample {
	elapsed := stamp.Mono - d.lastMono
	if d.tick == 0 || elapsed <= 0 {
		elapsed = time.Second
	}
	d.lastMono = stamp.Mono

	if d.boot.IsZero() {
		d.boot = stamp.At.Add(-72 * time.Hour)
	}

	d.tick++
	d.spike = d.tick%37 < 4

	d.advanceCPU(elapsed)
	d.advanceMem()
	d.advanceDisks()
	d.advanceNets()
	d.advanceProcs()

	return d.machineSample(stamp)
}

// newStart returns the next start value used for process identity.
func (d *Dummy) newStart() uint64 {
	d.nextStart++

	return d.nextStart
}

// startedAt maps a start value to a start time up to three days back.
func (d *Dummy) startedAt(id domain.ProcessIdentity) time.Time {
	const age = uint64((72 * time.Hour) / time.Second)

	return d.boot.Add(-time.Duration(id.Start%age) * time.Second)
}

// clampU64 clamps v into [lo, hi].
func clampU64(v, lo, hi int64) uint64 {
	if v < lo {
		v = lo
	}
	if v > hi {
		v = hi
	}

	return uint64(v)
}
