package domain

import "time"

// CPUTimes is cumulative scheduler time in nanoseconds. Busy counts time not
// idle. A source that only has total and idle still fills both fields.
type CPUTimes struct {
	Busy  uint64
	Total uint64
}

// Memory is a memory reading at one instant, in bytes.
type Memory struct {
	Total, Used, Available, Free, Cached Value
	UsedPercent                          Value
	SwapTotal, SwapUsed, SwapUsedPercent Value
}

// Disk carries space in bytes and cumulative IO counters.
type Disk struct {
	Device                string
	Mount                 string
	Total, Free           Value
	UsedPercent           Value
	ReadBytes, WriteBytes uint64
	ReadRate, WriteRate   Value
}

// Net carries one interface's cumulative counters and link state.
type Net struct {
	Name               string
	Up                 bool
	RXBytes, TXBytes   uint64
	RXErrors, TXErrors uint64
	RXRate, TXRate     Value
}

// Temp is one temperature sensor. An unreadable sensor keeps its name and an
// invalid value.
type Temp struct {
	Name    string
	Celsius Value
}

// Sample is one read of the machine. A Source fills the raw counters and
// leaves computed fields invalid; Compute fills CPU percent, memory used, and
// rates from two consecutive samples. Gap marks a sample that follows a
// suspend or a clock step, so the UI breaks the graph line instead of drawing
// across the gap.
type Sample struct {
	Stamp       Stamp
	Host        string
	OS          string
	Uptime      time.Duration
	CPUTotal    CPUTimes
	CPUCores    []CPUTimes
	CPUPercent  Value
	CorePercent []Value
	Load        [3]Value
	Mem         Memory
	Disks       []Disk
	Nets        []Net
	Procs       Value
	Threads     Value
	Handles     Value
	Temps       []Temp
	Gap         bool
}
