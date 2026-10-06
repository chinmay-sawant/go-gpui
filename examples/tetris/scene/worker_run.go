package scene

// next hands out request ids; only the tick calls it.
func (w *worker) next() uint64 {
	w.seq++

	return w.seq
}

// send queues a request, or reports a busy queue to the tick.
func (w *worker) send(req request) uint64 {
	select {
	case w.reqs <- req:
	case <-w.quit:
		w.push(Result{ID: req.id, Kind: req.kind, Err: ErrWorkerBusy})
	default:
		w.push(Result{ID: req.id, Kind: req.kind, Err: ErrWorkerBusy})
	}

	return req.id
}

// push offers one result without ever blocking the worker.
func (w *worker) push(r Result) {
	select {
	case w.res <- r:
	default:
	}
}

// run serializes the blocking store calls until Close.
func (w *worker) run() {
	defer close(w.done)

	for {
		select {
		case req := <-w.reqs:
			w.exec(req)
		case <-w.quit:
			w.drain()

			return
		}
	}
}

// drain finishes requests already queued at Close; their contexts are
// canceled, so the store returns quickly.
func (w *worker) drain() {
	for {
		select {
		case req := <-w.reqs:
			w.exec(req)
		default:
			return
		}
	}
}
