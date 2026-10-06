package ui

// The recording methods of the fake backend.

func (f *fakeBackend) Control(c Control) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.controls = append(f.controls, c)

	return nil
}

func (f *fakeBackend) Active() {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.actives++
}

func (f *fakeBackend) Page(req PageRequest) {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.pages = append(f.pages, req)
}

func (f *fakeBackend) Summary(gen uint64) {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.summary = append(f.summary, gen)
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
