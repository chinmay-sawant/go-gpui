package wire

import (
	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
	"github.com/chinmay-sawant/ownframe/examples/download-manager/ui"
)

// runSummary answers one summary request. The UI asks at a bounded rate,
// never once per progress event.
func (b *Backend) runSummary(gen uint64) {
	counts, err := b.store.Summary(b.ctx)
	if err != nil {
		b.notice("Summary: " + err.Error())

		return
	}

	summary := ui.Summary{
		Queued:    counts.ByState[domain.StateQueued],
		Running:   counts.ByState[domain.StateRunning],
		Paused:    counts.ByState[domain.StatePaused],
		Completed: counts.ByState[domain.StateCompleted],
		Failed:    counts.ByState[domain.StateFailed],
		Cancelled: counts.ByState[domain.StateCancelled],
	}

	b.push(ui.Update{Kind: ui.UpdateSummary, Gen: gen, Summary: &summary})
}
