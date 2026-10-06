package collector

import (
	"context"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// collectDetail submits one read of the tracked process. Nothing is read when
// no process is selected.
func (m *Manager) collectDetail(ctx context.Context, gen uint64) {
	src, name := m.currentSource()

	m.mu.Lock()
	id := m.tracked
	m.mu.Unlock()

	if id == (domain.ProcessIdentity{}) {
		return
	}

	if !m.detailBusy.CompareAndSwap(false, true) {
		m.bumpSkipped()

		return
	}

	queued := m.pool.TryDo(ctx, func(cctx context.Context) error {
		defer m.detailBusy.Store(false)

		if !m.stillCurrent(gen) {
			return nil
		}

		d, err := src.Detail(cctx, id)
		if err != nil {
			return err
		}

		m.publishDetail(gen, name, id, d)

		return nil
	}, func(err error) {
		m.recordError("detail", name, err)
	})

	if !queued {
		m.detailBusy.Store(false)
		m.bumpSkipped()
	}
}
