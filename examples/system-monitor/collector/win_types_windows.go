//go:build windows

package collector

import "time"

// filetime is one Windows FILETIME: 100 ns units since 1601.
type filetime struct {
	low  uint32
	high uint32
}

func (f filetime) u64() uint64 { return uint64(f.high)<<32 | uint64(f.low) }

// memoryStatusEx matches MEMORYSTATUSEX.
type memoryStatusEx struct {
	length               uint32
	memoryLoad           uint32
	totalPhys            uint64
	availPhys            uint64
	totalPageFile        uint64
	availPageFile        uint64
	totalVirtual         uint64
	availVirtual         uint64
	availExtendedVirtual uint64
}

// processEntry32 matches PROCESSENTRY32W on 64 bit Windows. The explicit pad
// keeps th32DefaultHeapID aligned like the C struct.
type processEntry32 struct {
	size            uint32
	cntUsage        uint32
	processID       uint32
	_               uint32
	defaultHeapID   uintptr
	moduleID        uint32
	threads         uint32
	parentProcessID uint32
	priClassBase    int32
	flags           uint32
	exeFile         [260]uint16
}

// ioCounters matches IO_COUNTERS.
type ioCounters struct {
	readOps, writeOps, otherOps       uint64
	readBytes, writeBytes, otherBytes uint64
}

// processMemoryCounters matches PROCESS_MEMORY_COUNTERS.
type processMemoryCounters struct {
	cb                         uint32
	pageFaultCount             uint32
	peakWorkingSetSize         uintptr
	workingSetSize             uintptr
	quotaPeakPagedPoolUsage    uintptr
	quotaPagedPoolUsage        uintptr
	quotaPeakNonPagedPoolUsage uintptr
	quotaNonPagedPoolUsage     uintptr
	pagefileUsage              uintptr
	peakPagefileUsage          uintptr
}

// filetimeToTime converts a FILETIME to wall clock time.
func filetimeToTime(f filetime) time.Time {
	const epoch = 116444736000000000 // 1601 to 1970 in 100 ns units

	return time.Unix(0, (int64(f.u64())-epoch)*100)
}
