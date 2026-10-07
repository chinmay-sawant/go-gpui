package ui

// run owns the backend: it executes queued jobs one at a time and sends the
// answers back. It exits when stop closes.
func (w *worker) run() {
	defer close(w.done)

	for {
		select {
		case <-w.stop:
			return
		case <-w.wake:
		}

		for {
			j, ok := w.pop()
			if !ok {
				break
			}

			w.emit(w.exec(j))

			w.mu.Lock()
			w.busy = false
			w.mu.Unlock()
		}
	}
}

// pop takes the next job. A fetch older than the newest posted fetch is
// discarded, which cancels a stale range load before it runs.
func (w *worker) pop() (job, bool) {
	w.mu.Lock()
	defer w.mu.Unlock()

	for len(w.queue) > 0 {
		j := w.queue[0]
		w.queue = w.queue[1:]

		if j.kind == jobFetch && j.gen != w.fetch {
			continue
		}

		w.busy = true

		return j, true
	}

	return job{}, false
}

// emit sends one answer. A fetch answer is dropped when the channel is
// full; a command answer waits, because the UI must see it.
func (w *worker) emit(r result) {
	if r.kind == jobFetch {
		select {
		case w.out <- r:
		default:
		}

		return
	}

	select {
	case w.out <- r:
	case <-w.stop:
	}
}
