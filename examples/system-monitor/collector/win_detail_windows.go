//go:build windows

package collector

import (
	"context"
	"syscall"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// Detail reads one process in full. When the creation time is readable and
// differs from the identity's start value, the PID was reused and the call
// reports ErrGone.
func (s *windowsSource) Detail(ctx context.Context, id domain.ProcessIdentity) (domain.ProcessDetail, error) {
	if err := ctx.Err(); err != nil {
		return domain.ProcessDetail{}, err
	}

	entry, ok := findEntry(uint32(id.PID))
	if !ok {
		return domain.ProcessDetail{}, domain.ErrGone
	}

	row := winRow(entry)
	if row.ID.Start != 0 && id.Start != 0 && row.ID.Start != id.Start {
		return domain.ProcessDetail{}, domain.ErrGone
	}

	d := domain.ProcessDetail{Process: row}
	d.Exe = queryImagePath(uint32(id.PID))
	d.Command = d.Exe

	if io, ok := queryIO(uint32(id.PID)); ok {
		d.ReadBytes = domain.Bytes(float64(io.readBytes))
		d.WriteBytes = domain.Bytes(float64(io.writeBytes))
	}

	if parent, ok := findEntry(entry.parentProcessID); ok {
		d.Parent = domain.ProcessIdentity{PID: int32(entry.parentProcessID)}
		if t, _, ok := queryProcess(entry.parentProcessID); ok {
			d.Parent.Start = t.start
		}
		d.ParentName = syscall.UTF16ToString(parent.exeFile[:])
	}

	return d, nil
}

// findEntry, queryImagePath, and queryIO are in win_query_windows.go.
