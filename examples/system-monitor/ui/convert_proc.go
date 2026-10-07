package ui

import (
	"github.com/chinmay-sawant/ownframe/examples/system-monitor/collector"
	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// toProcess converts one table row. A value the source could not read keeps
// its Known flag false, so the table shows n/a instead of a zero.
func toProcess(p domain.Process) Process {
	cpu, cpuOK := p.CPU.Get()
	mem, memOK := p.Memory.Get()
	threads, threadsOK := p.Threads.Get()

	return Process{
		ID:           p.ID.String(),
		PID:          int(p.ID.PID),
		Name:         p.Name,
		User:         p.User,
		State:        p.State,
		CPU:          cpu,
		CPUKnown:     cpuOK,
		Mem:          uint64(max(mem, 0)),
		MemKnown:     memOK,
		Start:        p.StartedAt,
		Threads:      int(max(threads, 0)),
		ThreadsKnown: threadsOK,
	}
}

// toDetail converts one full process read.
func toDetail(d domain.ProcessDetail) ProcDetail {
	out := ProcDetail{
		Process: toProcess(d.Process),
		Command: d.Command,
		Exe:     d.Exe,
		Cwd:     d.Cwd,
	}

	if d.Parent.PID != 0 {
		out.Parent = d.Parent.String()
		if d.ParentName != "" {
			out.Parent += " (" + d.ParentName + ")"
		}
	}

	if v, ok := d.Virtual.Get(); ok {
		out.Virtual, out.VirtualOK = uint64(max(v, 0)), true
	}

	if v, ok := d.Handles.Get(); ok {
		out.Handles, out.HandlesOK = uint64(max(v, 0)), true
	}

	if v, ok := d.Priority.Get(); ok {
		out.Priority, out.PriorityOK = int64(v), true
	}

	if v, ok := d.ReadRate.Get(); ok {
		out.ReadRate, out.ReadRateOK = v, true
	}

	if v, ok := d.WriteRate.Get(); ok {
		out.WriteRate, out.WriteRateOK = v, true
	}

	return out
}

// toTracked converts the manager's tracked state. OK means a result for the
// identity has arrived; Loading or an error leaves it false.
func toTracked(t collector.Tracked) Tracked {
	out := Tracked{
		ID:      t.ID.String(),
		Detail:  toDetail(t.Detail),
		Loading: t.Loading,
		Err:     t.Err,
	}
	out.OK = !t.Loading && t.Err == "" && t.ID != (domain.ProcessIdentity{})

	return out
}
