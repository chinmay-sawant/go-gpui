//go:build windows

package collector

import (
	"syscall"
	"unsafe"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// windowsDisks lists fixed and removable drives with their free space.
func windowsDisks() []domain.Disk {
	mask, _, _ := procGetLogicalDrives.Call()

	var out []domain.Disk

	for i := range 26 {
		if mask&(1<<uint(i)) == 0 {
			continue
		}

		root := string(rune('A'+i)) + `:\`
		rootPtr, err := syscall.UTF16PtrFromString(root)
		if err != nil {
			continue
		}

		kind, _, _ := procGetDriveTypeW.Call(uintptr(unsafe.Pointer(rootPtr)))
		if kind != driveFixed && kind != driveRemovable {
			continue
		}

		var avail, total, free uint64
		r, _, _ := procGetDiskFreeSpaceExW.Call(
			uintptr(unsafe.Pointer(rootPtr)),
			uintptr(unsafe.Pointer(&avail)),
			uintptr(unsafe.Pointer(&total)),
			uintptr(unsafe.Pointer(&free)),
		)
		if r == 0 {
			continue
		}

		out = append(out, domain.Disk{
			Device: root[:2],
			Mount:  root,
			Total:  domain.Bytes(float64(total)),
			Free:   domain.Bytes(float64(free)),
		})
	}

	return out
}
