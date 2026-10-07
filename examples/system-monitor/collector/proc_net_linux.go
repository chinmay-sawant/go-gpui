//go:build linux

package collector

import (
	"context"
	"strings"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// readNets reads interface counters and link state.
func readNets(ctx context.Context) []domain.Net {
	out := parseNetDev(readFile("/proc/net/dev"))
	for i := range out {
		if err := ctx.Err(); err != nil {
			return out
		}
		out[i].Up = linkUp(out[i].Name)
	}

	return out
}

// parseNetDev reads the two-column format of /proc/net/dev. Field 0 is RX
// bytes, field 2 RX errors, field 8 TX bytes, and field 10 TX errors.
func parseNetDev(data []byte) []domain.Net {
	var out []domain.Net

	for _, line := range strings.Split(string(data), "\n") {
		name, rest, ok := strings.Cut(line, ":")
		if !ok {
			continue
		}

		name = strings.TrimSpace(name)
		fields := strings.Fields(rest)
		if name == "" || len(fields) < 11 {
			continue
		}

		out = append(out, domain.Net{
			Name:     name,
			RXBytes:  parseUint(fields[0]),
			RXErrors: parseUint(fields[2]),
			TXBytes:  parseUint(fields[8]),
			TXErrors: parseUint(fields[10]),
		})
	}

	return out
}

// linkUp reads the operational state. A missing file means the interface
// disappeared between reads, which counts as down.
func linkUp(name string) bool {
	data := readFile("/sys/class/net/" + name + "/operstate")
	if data == nil {
		return false
	}

	state := strings.TrimSpace(string(data))

	return state == "up" || state == "unknown"
}
