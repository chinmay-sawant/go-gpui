//go:build windows

package collector

import (
	"syscall"
	"unsafe"
)

// findEntry finds one process in a fresh toolhelp snapshot.
func findEntry(pid uint32) (processEntry32, bool) {
	snap, _, _ := procCreateToolhelp32Snapshot.Call(th32csSnapProcess, 0)
	if snap == invalidHandle {
		return processEntry32{}, false
	}
	defer procCloseHandle.Call(snap)

	var entry processEntry32
	entry.size = uint32(unsafe.Sizeof(entry))

	r, _, _ := procProcess32FirstW.Call(snap, uintptr(unsafe.Pointer(&entry)))
	for r != 0 {
		if entry.processID == pid {
			return entry, true
		}
		r, _, _ = procProcess32NextW.Call(snap, uintptr(unsafe.Pointer(&entry)))
	}

	return processEntry32{}, false
}

// queryImagePath reads the full executable path. It needs only query rights.
func queryImagePath(pid uint32) string {
	h, _, _ := procOpenProcess.Call(processQueryLimitedInfo, 0, uintptr(pid))
	if h == 0 {
		return ""
	}
	defer procCloseHandle.Call(h)

	buf := make([]uint16, 1024)
	size := uint32(len(buf))
	r, _, _ := procQueryFullProcessImageName.Call(
		h,
		0,
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(unsafe.Pointer(&size)),
	)
	if r == 0 {
		return ""
	}

	return syscall.UTF16ToString(buf[:size])
}

// queryIO reads cumulative read and write bytes.
func queryIO(pid uint32) (ioCounters, bool) {
	h, _, _ := procOpenProcess.Call(processQueryLimitedInfo, 0, uintptr(pid))
	if h == 0 {
		return ioCounters{}, false
	}
	defer procCloseHandle.Call(h)

	var io ioCounters
	if r, _, _ := procGetProcessIoCounters.Call(h, uintptr(unsafe.Pointer(&io))); r == 0 {
		return ioCounters{}, false
	}

	return io, true
}
