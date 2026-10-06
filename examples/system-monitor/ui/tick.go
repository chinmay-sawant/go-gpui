package ui

import "context"

// Tick applies worker results and repaints retained operations. It runs once
// per window frame and never waits on IO: it drains a bounded batch and
// changes the display list in place. It never parses or lays out.
func (a *App) Tick(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	a.drain()

	if !a.valid() {
		a.rebind()
	}

	a.paint()

	return nil
}

// drain applies the mailbox's bounded batch: at most one item per feed. A
// result tagged with an older generation, or a detail for another identity,
// is discarded. A process snapshot is only offered; refresh adopts it.
func (a *App) drain() {
	b := a.mail.take()
	s := a.state
	gen := a.currentGen()

	if b.problems != nil && b.problems.gen == gen {
		s.problems = b.problems.v
		s.dirtyText = true
	}

	if b.summary != nil && b.summary.gen == gen {
		s.at = b.summary.v.At

		for _, r := range b.summary.v.Readings {
			s.apply(r, b.summary.v.Gap)
		}

		s.dirtyText = true
	}

	if b.procs != nil && b.procs.gen == gen {
		s.table.offer(b.procs.v)
		s.dirtyText = true
	}

	if b.tracked != nil && b.tracked.gen == gen {
		if t := b.tracked.v; t.ID == s.sel.id {
			s.sel.trk = t
			s.dirtyText = true
		}
	}
}

// currentGen reads the request generation the pollers tag results with.
func (a *App) currentGen() uint64 {
	return a.state.gen.Load()
}
