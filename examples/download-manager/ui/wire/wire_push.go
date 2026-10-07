package wire

import "github.com/chinmay-sawant/ownframe/examples/download-manager/ui"

// push queues one result. A full results channel drops it, and every kind
// that travels here is a snapshot the UI re-requests. State transitions
// never take this path; they come from the engine's own outbox, which does
// not drop them.
func (b *Backend) push(u ui.Update) {
	select {
	case b.results <- u:
	default:
		b.mu.Lock()
		b.dropped++
		b.mu.Unlock()
	}
}

// notice queues one message. It waits instead of dropping, so an error is
// never lost; the UI drains every tick.
func (b *Backend) notice(msg string) {
	select {
	case b.results <- ui.Update{Kind: ui.UpdateNotice, Notice: msg}:
	case <-b.ctx.Done():
	}
}
