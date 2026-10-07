//go:build windows

package collector

import (
	"context"
	"time"
	"unsafe"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// Sample reads CPU, memory, uptime, and disk space. GetSystemTimes reports
// idle, kernel, and user time in 100 ns units; kernel includes idle.
func (s *windowsSource) Sample(ctx context.Context) (domain.Sample, error) {
	if err := ctx.Err(); err != nil {
		return domain.Sample{}, err
	}

	var idle, kernel, user filetime
	r, _, callErr := procGetSystemTimes.Call(
		uintptr(unsafe.Pointer(&idle)),
		uintptr(unsafe.Pointer(&kernel)),
		uintptr(unsafe.Pointer(&user)),
	)
	if r == 0 {
		return domain.Sample{}, callErr
	}

	total := kernel.u64() + user.u64()

	var busy uint64
	if total > idle.u64() {
		busy = total - idle.u64()
	}

	smp := domain.Sample{
		Stamp:    domain.Stamp{At: time.Now()},
		Host:     s.host,
		OS:       "windows",
		CPUTotal: domain.CPUTimes{Busy: busy * 100, Total: total * 100},
	}

	var ms memoryStatusEx
	ms.length = uint32(unsafe.Sizeof(ms))
	if r, _, _ := procGlobalMemoryStatusEx.Call(uintptr(unsafe.Pointer(&ms))); r != 0 {
		smp.Mem = domain.Memory{
			Total:     domain.Bytes(float64(ms.totalPhys)),
			Available: domain.Bytes(float64(ms.availPhys)),
			SwapTotal: domain.Bytes(float64(ms.totalPageFile)),
		}
		if ms.totalPageFile >= ms.availPageFile {
			smp.Mem.SwapUsed = domain.Bytes(float64(ms.totalPageFile - ms.availPageFile))
		}
	}

	if r, _, _ := procGetTickCount64.Call(); r != 0 {
		smp.Uptime = time.Duration(r) * time.Millisecond
	}

	smp.Disks = windowsDisks()

	return smp, nil
}

// windowsDisks is in win_disks_windows.go.
