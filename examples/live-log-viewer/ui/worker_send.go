package ui

// send queues a request without ever blocking the UI loop. A full queue
// keeps the newest disposable snapshot for the next tick and counts the
// coalesce.
func (a *App) send(r req) {
	if a.feed == nil {
		return
	}

	select {
	case a.reqs <- r:
	default:
		a.dropped.Add(1)

		a.pendingMu.Lock()
		a.pendingReq = &r
		a.pendingMu.Unlock()
	}
}

// retry sends the coalesced snapshot, if any.
func (a *App) retry() {
	a.pendingMu.Lock()
	r := a.pendingReq
	a.pendingReq = nil
	a.pendingMu.Unlock()

	if r == nil {
		return
	}

	select {
	case a.reqs <- *r:
	default:
		a.pendingMu.Lock()
		a.pendingReq = r
		a.pendingMu.Unlock()
	}
}
