//go:build linux

package collector

import (
	"testing"
	"time"
)

// cpuFixture is a three line /proc/stat.
const cpuFixture = `cpu  100 0 50 1000 20 0 5 0 0 0
cpu0 50 0 25 500 10 0 2 0 0 0
cpu1 50 0 25 500 10 0 3 0 0 0
intr 0 0 0
`

// TestParseCPUStat checks tick to nanosecond conversion, the idle plus iowait
// definition of busy, and per-core extraction.
func TestParseCPUStat(t *testing.T) {
	total, cores, err := parseCPUStat([]byte(cpuFixture))
	if err != nil {
		t.Fatal(err)
	}

	if len(cores) != 2 {
		t.Fatalf("cores = %d", len(cores))
	}

	wantTotal := uint64(1175) * tickNS
	if total.Total != wantTotal {
		t.Fatalf("total = %d, want %d", total.Total, wantTotal)
	}

	wantBusy := uint64(155) * tickNS
	if total.Busy != wantBusy {
		t.Fatalf("busy = %d, want %d", total.Busy, wantBusy)
	}
	if cores[0].Total != uint64(587)*tickNS {
		t.Fatalf("core0 total = %d", cores[0].Total)
	}
}

// TestParseCPUStatBad checks that a file without a cpu line is an error, not
// a zeroed reading.
func TestParseCPUStatBad(t *testing.T) {
	if _, _, err := parseCPUStat([]byte("intr 1 2 3\n")); err == nil {
		t.Fatal("missing cpu line accepted")
	}
}

// TestParseUptime checks the seconds field.
func TestParseUptime(t *testing.T) {
	if got := parseUptime([]byte("12345.67 98765.43\n")); got != 12345670*time.Millisecond {
		t.Fatalf("uptime = %v", got)
	}
}

// TestParseLoadAvg checks the three values and a short line.
func TestParseLoadAvg(t *testing.T) {
	got := parseLoadAvg([]byte("0.52 0.58 0.59 1/234 5678\n"))
	if !got[0].Valid || got[0].N != 0.52 || got[2].N != 0.59 {
		t.Fatalf("load = %+v", got)
	}

	short := parseLoadAvg([]byte("1.0\n"))
	if !short[0].Valid || short[1].Valid {
		t.Fatalf("short load = %+v", short)
	}
}
