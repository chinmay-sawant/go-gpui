//go:build windows

package collector

import (
	"time"
	"unsafe"
)

// winTimes is one process's CPU and start times from GetProcessTimes.
type winTimes struct {
	start   uint64
	cpu     uint64
	started time.Time
}

// queryProcess opens one process and reads its times and working set. A
// protected process reports ok false, which leaves the values invalid.
func queryProcess(pid uint32) (winTimes, uintptr, bool) {
	h, _, _ := procOpenProcess.Call(processQueryLimitedInfo, 0, uintptr(pid))
	if h == 0 {
		return winTimes{}, 0, false
	}
	defer procCloseHandle.Call(h)

	var creation, exit, kernel, user filetime
	r, _, _ := procGetProcessTimes.Call(
		h,
		uintptr(unsafe.Pointer(&creation)),
		uintptr(unsafe.Pointer(&exit)),
		uintptr(unsafe.Pointer(&kernel)),
		uintptr(unsafe.Pointer(&user)),
	)
	if r == 0 {
		return winTimes{}, 0, false
	}

	var mem uintptr

	var pmc processMemoryCounters
	pmc.cb = uint32(unsafe.Sizeof(pmc))
	if r, _, _ := procGetProcessMemoryInfo.Call(h, uintptr(unsafe.Pointer(&pmc)), uintptr(pmc.cb)); r != 0 {
		mem = pmc.workingSetSize
	}

	return winTimes{
		start:   creation.u64(),
		cpu:     (kernel.u64() + user.u64()) * 100,
		started: filetimeToTime(creation),
	}, mem, true
}
