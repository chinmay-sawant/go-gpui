//go:build linux

package collector

import (
	"context"
	"os"
	"runtime"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// Capabilities reports what this kernel exposes.
func (s *procSource) Capabilities() domain.Capabilities {
	return domain.Capabilities{
		PerCoreCPU:    true,
		Processes:     true,
		ProcessDetail: true,
		Disk:          s.hasDiskSpace,
		DiskIO:        s.hasDiskIO,
		Net:           s.hasNet,
		Load:          pathExists("/proc/loadavg"),
		Sensors:       s.hasSensors,
		Handles:       false,
		Swap:          true,
		Notes: []string{
			"process user names appear in the detail view only",
			"handle counts are not exposed by /proc",
		},
	}
}

// Sample reads the machine counters.
func (s *procSource) Sample(ctx context.Context) (domain.Sample, error) {
	if err := ctx.Err(); err != nil {
		return domain.Sample{}, err
	}

	stat, err := os.ReadFile("/proc/stat")
	if err != nil {
		return domain.Sample{}, err
	}

	total, cores, err := parseCPUStat(stat)
	if err != nil {
		return domain.Sample{}, err
	}

	smp := domain.Sample{
		Stamp:    domain.Stamp{At: time.Now()},
		Host:     s.host,
		OS:       runtime.GOOS,
		CPUTotal: total,
		CPUCores: cores,
	}

	if data, err := os.ReadFile("/proc/uptime"); err == nil {
		smp.Uptime = parseUptime(data)
	}
	if data, err := os.ReadFile("/proc/loadavg"); err == nil {
		smp.Load = parseLoadAvg(data)
	}
	if data, err := os.ReadFile("/proc/meminfo"); err == nil {
		smp.Mem = parseMemInfo(data)
	}
	if n, ok := countProcs(); ok {
		smp.Procs = domain.Count(float64(n))
	}

	smp.Disks = s.readDisks(ctx)
	smp.Nets = readNets(ctx)
	smp.Temps = readTemps(ctx)

	return smp, nil
}
