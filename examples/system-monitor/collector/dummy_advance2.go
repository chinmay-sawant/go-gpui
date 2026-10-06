package collector

// dmDisk is one simulated disk: space plus cumulative IO.
type dmDisk struct {
	name, mount string
	total, free uint64
	read, write uint64
}

// dmNet is one simulated interface with cumulative counters.
type dmNet struct {
	name   string
	up     bool
	rx, tx uint64
	errs   uint64
}

// advanceDisks grows IO counters and walks free space. Every 23rd tick writes
// a burst, so parsing and graphs see a spike.
func (d *Dummy) advanceDisks() {
	for i := range d.disks {
		write := d.rng.Int63n(1 << 22)
		if d.tick%23 == 0 {
			write += 384 << 20
		}
		d.disks[i].read += uint64(d.rng.Int63n(1 << 24))
		d.disks[i].write += uint64(write)

		step := int64(d.rng.Intn(128<<20)) - (64 << 20)
		d.disks[i].free = clampU64(int64(d.disks[i].free)+step, 10<<30, int64(d.disks[i].total))
	}
}

// advanceNets grows interface counters, adds a USB interface once, and raises
// an occasional error.
func (d *Dummy) advanceNets() {
	if d.tick == 30 {
		d.nets = append(d.nets, dmNet{name: "usb0", up: true})
	}

	for i := range d.nets {
		rx := d.rng.Int63n(1 << 20)
		if d.tick%19 == 0 {
			rx += 96 << 20
		}
		d.nets[i].rx += uint64(rx)
		d.nets[i].tx += uint64(d.rng.Int63n(1 << 18))

		if d.rng.Float64() < 0.004 {
			d.nets[i].errs++
		}
	}
}
