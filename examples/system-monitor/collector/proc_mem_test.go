//go:build linux

package collector

import "testing"

// memFixture is a trimmed /proc/meminfo.
const memFixture = `MemTotal:       32768000 kB
MemFree:         1000000 kB
MemAvailable:   20000000 kB
Buffers:          500000 kB
Cached:          8000000 kB
SwapTotal:       8388608 kB
SwapFree:        8000000 kB
`

// TestParseMemInfo checks kB to byte conversion, used from available, and
// swap used.
func TestParseMemInfo(t *testing.T) {
	m := parseMemInfo([]byte(memFixture))

	if m.Total.N != 32768000*1024 {
		t.Fatalf("total = %v", m.Total)
	}
	if m.Available.N != 20000000*1024 {
		t.Fatalf("available = %v", m.Available)
	}
	if m.Used.N != m.Total.N-m.Available.N {
		t.Fatalf("used = %v", m.Used)
	}
	if m.SwapUsed.N != 388608*1024 {
		t.Fatalf("swap used = %v", m.SwapUsed)
	}
	if !m.UsedPercent.Valid {
		t.Fatal("used percent missing")
	}
}

// TestParseMemInfoMissing checks that a file without totals leaves the values
// invalid rather than zero.
func TestParseMemInfoMissing(t *testing.T) {
	m := parseMemInfo([]byte("SomethingElse: 1 kB\n"))

	if m.Total.Valid || m.Used.Valid {
		t.Fatalf("mem = %+v", m)
	}
}
