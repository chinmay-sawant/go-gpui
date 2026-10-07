//go:build windows

package collector

import (
	"context"
	"syscall"
	"unsafe"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// Processes lists processes through a toolhelp snapshot. A protected process
// stays in the table with its name and thread count; its times and memory
// remain unavailable.
func (s *windowsSource) Processes(ctx context.Context) ([]domain.Process, error) {
	snap, _, callErr := procCreateToolhelp32Snapshot.Call(th32csSnapProcess, 0)
	if snap == invalidHandle {
		return nil, callErr
	}
	defer procCloseHandle.Call(snap)

	var entry processEntry32
	entry.size = uint32(unsafe.Sizeof(entry))

	r, _, _ := procProcess32FirstW.Call(snap, uintptr(unsafe.Pointer(&entry)))
	rows := make([]domain.Process, 0, 256)

	for r != 0 {
		if len(rows)%256 == 0 {
			if err := ctx.Err(); err != nil {
				return nil, err
			}
		}

		rows = append(rows, winRow(entry))

		r, _, _ = procProcess32NextW.Call(snap, uintptr(unsafe.Pointer(&entry)))
	}

	return rows, nil
}

// winRow builds a table row from one snapshot entry.
func winRow(entry processEntry32) domain.Process {
	row := domain.Process{
		ID:       domain.ProcessIdentity{PID: int32(entry.processID)},
		Name:     syscall.UTF16ToString(entry.exeFile[:]),
		Threads:  domain.Count(float64(entry.threads)),
		Priority: domain.Count(float64(entry.priClassBase)),
	}

	if t, mem, ok := queryProcess(entry.processID); ok {
		row.ID.Start = t.start
		row.CPUTime = t.cpu
		row.StartedAt = t.started
		if mem > 0 {
			row.Memory = domain.Bytes(float64(mem))
		}
	}

	return row
}
