package collector

import (
	"context"
)

// collectSummary submits one machine sample. A busy pool or a call still in
// flight skips this tick; the next tick tries again.
func (m *Manager) collectSummary(ctx context.Context, gen uint64) {
	src, name := m.currentSource()
	if !m.summaryBusy.CompareAndSwap(false, true) {
		m.bumpSkipped()

		return
	}

	queued := m.pool.TryDo(ctx, func(cctx context.Context) error {
		defer m.summaryBusy.Store(false)

		if !m.stillCurrent(gen) {
			return nil
		}

		raw, err := src.Sample(cctx)
		if err != nil {
			return err
		}

		m.publishSample(gen, name, raw)

		return nil
	}, func(err error) {
		m.recordError("sample", name, err)
	})

	if !queued {
		m.summaryBusy.Store(false)
		m.bumpSkipped()
	}
}

// collectProcesses submits one process table read.
func (m *Manager) collectProcesses(ctx context.Context, gen uint64) {
	src, name := m.currentSource()
	if !m.procBusy.CompareAndSwap(false, true) {
		m.bumpSkipped()

		return
	}

	queued := m.pool.TryDo(ctx, func(cctx context.Context) error {
		defer m.procBusy.Store(false)

		if !m.stillCurrent(gen) {
			return nil
		}

		rows, err := src.Processes(cctx)
		if err != nil {
			return err
		}

		m.publishProcesses(gen, name, rows)

		return nil
	}, func(err error) {
		m.recordError("processes", name, err)
	})

	if !queued {
		m.procBusy.Store(false)
		m.bumpSkipped()
	}
}
