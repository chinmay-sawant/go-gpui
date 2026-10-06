package ui

// The recording methods of the fake backend.

func (f *fakeBackend) Control(c Control) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.controls = append(f.controls, c)

	return nil
}

func (f *fakeBackend) Active() error {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.busy {
		return errFakeBusy
	}

	f.actives++

	return nil
}

func (f *fakeBackend) Page(req PageRequest) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.busy {
		return errFakeBusy
	}

	f.pages = append(f.pages, req)

	return nil
}

func (f *fakeBackend) Summary(gen uint64) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.busy {
		return errFakeBusy
	}

	f.summary = append(f.summary, gen)

	return nil
}

// setBusy makes the request methods fail until it is cleared.
func (f *fakeBackend) setBusy(busy bool) {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.busy = busy
}

// Poll returns up to budget queued updates, oldest first.
func (f *fakeBackend) Poll(budget int) []Update {
	f.mu.Lock()
	defer f.mu.Unlock()

	if budget <= 0 || len(f.queue) == 0 {
		return nil
	}

	if budget > len(f.queue) {
		budget = len(f.queue)
	}

	out := make([]Update, budget)
	copy(out, f.queue[:budget])
	f.queue = f.queue[budget:]

	return out
}

func (f *fakeBackend) Dark() (bool, error) { return f.dark, nil }

func (f *fakeBackend) SetDark(dark bool) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.dark = dark

	return nil
}

func (f *fakeBackend) Close() error {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.closed = true

	return nil
}

// send queues one worker update for the next tick.
func (f *fakeBackend) send(u Update) {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.queue = append(f.queue, u)
}

// queued returns the number of updates not yet polled.
func (f *fakeBackend) queued() int {
	f.mu.Lock()
	defer f.mu.Unlock()

	return len(f.queue)
}
