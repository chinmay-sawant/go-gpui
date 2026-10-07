package domain

// Metric names are the stable keys storage and the UI agree on. The unit of
// each metric is fixed, and the comment on each constant states it.
type Metric string

const (
	MetricCPU        Metric = "cpu.total"     // percent
	MetricCPUCore    Metric = "cpu.core"      // percent, device = core index
	MetricLoad       Metric = "load"          // count, device = 1, 5, or 15
	MetricMemUsed    Metric = "mem.used"      // bytes
	MetricMemAvail   Metric = "mem.available" // bytes
	MetricMemPercent Metric = "mem.percent"   // percent
	MetricSwapUsed   Metric = "swap.used"     // bytes
	MetricDiskUsed   Metric = "disk.percent"  // percent, device = mount
	MetricDiskRead   Metric = "disk.read"     // bytes/s, device = device name
	MetricDiskWrite  Metric = "disk.write"    // bytes/s, device = device name
	MetricNetRX      Metric = "net.rx"        // bytes/s, device = interface
	MetricNetTX      Metric = "net.tx"        // bytes/s, device = interface
	MetricProcs      Metric = "procs"         // count
	MetricThreads    Metric = "threads"       // count
	MetricTemp       Metric = "temp"          // Celsius, device = sensor
)

// Unit returns the display unit for a metric.
func (m Metric) Unit() string {
	switch m {
	case MetricCPU, MetricCPUCore, MetricMemPercent, MetricDiskUsed:
		return "%"
	case MetricMemUsed, MetricMemAvail, MetricSwapUsed:
		return "B"
	case MetricDiskRead, MetricDiskWrite, MetricNetRX, MetricNetTX:
		return "B/s"
	case MetricTemp:
		return "C"
	default:
		return ""
	}
}
