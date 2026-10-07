package workbook

import (
	"errors"
)

// ErrPendingFull reports a full pending queue; the caller keeps the edits
// it tried to push.
var ErrPendingFull = errors.New("workbook: pending edit queue is full")

// Pending is a bounded FIFO of unsaved commands. A caller pushes a command
// and pops it only after storage acknowledges the commit, so a full queue
// refuses new work instead of dropping changes.
type Pending struct {
	cmds []Command
	cap  int
}

// NewPending builds a queue. A capacity below one becomes 64.
func NewPending(capacity int) *Pending {
	if capacity <= 0 {
		capacity = 64
	}

	return &Pending{cap: capacity}
}

// Push appends a command, or reports ErrPendingFull when the queue is full.
func (p *Pending) Push(cmd Command) error {
	if len(p.cmds) >= p.cap {
		return ErrPendingFull
	}

	p.cmds = append(p.cmds, cmd)

	return nil
}

// Pop removes and returns the oldest command. The caller owns it until the
// commit is acknowledged; on failure it must push it back or keep it.
func (p *Pending) Pop() (Command, bool) {
	if len(p.cmds) == 0 {
		return Command{}, false
	}

	cmd := p.cmds[0]
	p.cmds = p.cmds[1:]

	return cmd, true
}

// Len is the number of queued commands.
func (p *Pending) Len() int { return len(p.cmds) }

// Full reports whether the queue refuses new work.
func (p *Pending) Full() bool { return len(p.cmds) >= p.cap }

// Capacity is the queue bound.
func (p *Pending) Capacity() int { return p.cap }
