package collector

import "time"

// evolveProc advances one process: CPU time, resident memory, IO counters,
// and state.
func (d *Dummy) evolveProc(p *dmProc) {
	roll := d.rng.Float64()

	delta := time.Duration(d.rng.Int63n(300_000))
	if roll < 0.15 {
		delta = time.Millisecond + time.Duration(d.rng.Int63n(int64(8*time.Millisecond)))
	}
	if d.spike && roll > 0.9 {
		delta = 20 * time.Millisecond
	}
	p.cpu += uint64(delta)

	step := int64(d.rng.Intn(2<<20)) - (1 << 20)
	p.mem = clampU64(int64(p.mem)+step, 1<<20, 2<<30)

	if d.rng.Float64() < 0.1 {
		p.read += uint64(d.rng.Int63n(4 << 20))
		p.write += uint64(d.rng.Int63n(1 << 20))
	}

	if d.rng.Float64() < 0.4 {
		p.state = "running"
	} else {
		p.state = "sleeping"
	}
}

// advanceProcs evolves every process, then exits some and starts others so
// the table changes between samples.
func (d *Dummy) advanceProcs() {
	for _, p := range d.procs {
		d.evolveProc(p)
	}

	d.churn()
}

// churn exits a few processes and replaces them, keeping the table size.
func (d *Dummy) churn() {
	exits := 1 + d.rng.Intn(3)

	for range exits {
		if len(d.procs) <= 16 {
			break
		}

		idx := d.rng.Intn(len(d.procs))
		d.retired = append(d.retired, d.procs[idx].id)
		if len(d.retired) > 32 {
			d.retired = d.retired[1:]
		}
		d.procs = append(d.procs[:idx], d.procs[idx+1:]...)
	}

	for range exits {
		d.spawn()
	}
}
