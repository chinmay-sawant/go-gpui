package domain

import (
	"strconv"
	"time"
)

// Record is one metric value in a recording. Device is empty for machine-wide
// metrics; Valid false records an unavailable reading rather than dropping it.
type Record struct {
	Metric Metric
	Device string
	At     time.Time
	Mono   time.Duration
	Value  float64
	Valid  bool
}

// Records flattens a sample into the rows a recording stores. It emits every
// metric the UI can graph, valid or not, so a gap in a recording stays
// visible.
func (s Sample) Records() []Record {
	base := Record{At: s.Stamp.At, Mono: s.Stamp.Mono}
	add := func(metric Metric, device string, v Value) Record {
		r := base
		r.Metric, r.Device = metric, device
		r.Value, r.Valid = v.Get()

		return r
	}

	rows := []Record{
		add(MetricCPU, "", s.CPUPercent),
		add(MetricLoad, "1", s.Load[0]),
		add(MetricLoad, "5", s.Load[1]),
		add(MetricLoad, "15", s.Load[2]),
		add(MetricMemUsed, "", s.Mem.Used),
		add(MetricMemAvail, "", s.Mem.Available),
		add(MetricMemPercent, "", s.Mem.UsedPercent),
		add(MetricSwapUsed, "", s.Mem.SwapUsed),
		add(MetricProcs, "", s.Procs),
		add(MetricThreads, "", s.Threads),
	}

	for i, core := range s.CorePercent {
		rows = append(rows, add(MetricCPUCore, strconv.Itoa(i), core))
	}
	for _, d := range s.Disks {
		rows = append(rows,
			add(MetricDiskUsed, d.Mount, d.UsedPercent),
			add(MetricDiskRead, d.Device, d.ReadRate),
			add(MetricDiskWrite, d.Device, d.WriteRate),
		)
	}
	for _, n := range s.Nets {
		rows = append(rows, add(MetricNetRX, n.Name, n.RXRate), add(MetricNetTX, n.Name, n.TXRate))
	}
	for _, t := range s.Temps {
		rows = append(rows, add(MetricTemp, t.Name, t.Celsius))
	}

	return rows
}
