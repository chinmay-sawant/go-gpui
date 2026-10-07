package collector

import (
	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// rowOf builds the process table row for one simulated process.
func (d *Dummy) rowOf(p *dmProc) domain.Process {
	return domain.Process{
		ID:        p.id,
		Name:      p.name,
		User:      p.user,
		State:     p.state,
		StartedAt: d.startedAt(p.id),
		CPUTime:   p.cpu,
		Memory:    domain.Bytes(float64(p.mem)),
		Threads:   domain.Count(float64(p.threads)),
		Priority:  domain.Count(float64(p.nice)),
	}
}

// procRows copies the process table.
func (d *Dummy) procRows() []domain.Process {
	rows := make([]domain.Process, 0, len(d.procs))
	for _, p := range d.procs {
		rows = append(rows, d.rowOf(p))
	}

	return rows
}

// findProc returns the process with this identity, or nil when it exited.
func (d *Dummy) findProc(id domain.ProcessIdentity) *dmProc {
	for _, p := range d.procs {
		if p.id == id {
			return p
		}
	}

	return nil
}
