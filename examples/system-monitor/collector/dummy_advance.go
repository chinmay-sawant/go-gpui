package collector

import "time"

// advanceCPU adds one interval of scheduler time to every core. A rotating
// core spikes every 37 ticks, which shows up in the overview graph.
func (d *Dummy) advanceCPU(elapsed time.Duration) {
	for i := range d.cpu {
		fraction := 0.06 + d.rng.Float64()*0.12
		if d.spike && (i+int(d.tick/4))%2 == 0 {
			fraction = 0.82 + d.rng.Float64()*0.1
		}

		d.cpu[i].Total += uint64(elapsed)
		d.cpu[i].Busy += uint64(float64(elapsed) * fraction)
	}

	drift := (d.rng.Float64() - 0.5) * 0.3
	d.load[0] = clampF(d.load[0]+drift, 0.05, 16)
	d.load[1] = clampF(d.load[1]*0.97+d.load[0]*0.03, 0.05, 16)
	d.load[2] = clampF(d.load[2]*0.99+d.load[0]*0.01, 0.05, 16)
}

// advanceMem walks used memory and swap within sane bounds.
func (d *Dummy) advanceMem() {
	step := int64(d.rng.Intn(2<<27)) - (1 << 27)
	d.memUsed = clampU64(int64(d.memUsed)+step, 3<<30, 30<<30)

	if d.rng.Float64() < 0.3 {
		d.swapUsed = clampU64(int64(d.swapUsed)+step/4, 0, 6<<30)
	}
}

// clampF clamps a float into [lo, hi].
func clampF(v, lo, hi float64) float64 {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}

	return v
}
