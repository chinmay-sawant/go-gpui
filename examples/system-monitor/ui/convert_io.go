package ui

import (
	"fmt"
	"strings"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// diskReading graphs the combined read and write rate.
func diskReading(s domain.Sample) Reading {
	var rate, free float64

	rateOK, freeOK := false, false

	for _, d := range s.Disks {
		if v, ok := d.ReadRate.Get(); ok {
			rate += v
			rateOK = true
		}

		if v, ok := d.WriteRate.Get(); ok {
			rate += v
			rateOK = true
		}

		if v, ok := d.Free.Get(); ok {
			free += v
			freeOK = true
		}
	}

	r := Reading{ID: "disk", Value: rate, OK: rateOK, Text: formatRate(rate, rateOK)}
	if freeOK {
		r.Sub = fmt.Sprintf("%d disks \u00b7 %s free", len(s.Disks), formatBytes(free, true))
	}

	return r
}

// netReading graphs the combined receive and transmit rate.
func netReading(s domain.Sample) Reading {
	var rate float64

	rateOK := false
	ups := 0
	names := make([]string, 0, len(s.Nets))

	for _, n := range s.Nets {
		if v, ok := n.RXRate.Get(); ok {
			rate += v
			rateOK = true
		}

		if v, ok := n.TXRate.Get(); ok {
			rate += v
			rateOK = true
		}

		if n.Up {
			ups++
		}

		names = append(names, n.Name)
	}

	r := Reading{ID: "net", Value: rate, OK: rateOK, Text: formatRate(rate, rateOK)}
	if len(names) > 0 {
		r.Sub = fmt.Sprintf("%d/%d up \u00b7 %s", ups, len(names), strings.Join(names, ", "))
	}

	return r
}
