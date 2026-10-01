package ipc

// Send calls each Listen callback for channel.
// A callback that panics does not stop the others.
// An empty channel does nothing.
func Send(channel, payload string) {
	if channel == "" {
		return
	}

	mu.Lock()
	snap := append([]ear(nil), ears[channel]...)
	mu.Unlock()

	for _, item := range snap {
		call(item.fn, payload)
	}
}

func call(fn func(string), payload string) {
	defer func() { recover() }()

	fn(payload)
}

// Listen registers fn for channel.
// Cancel removes that registration only.
// An empty channel does nothing.
func Listen(channel string, fn func(string)) (cancel func()) {
	if channel == "" || fn == nil {
		return noop
	}

	mu.Lock()
	defer mu.Unlock()

	seq++
	id := seq
	ears[channel] = append(ears[channel], ear{id: id, fn: fn})

	return func() { removeEar(channel, id) }
}

func removeEar(channel string, id int64) {
	mu.Lock()
	defer mu.Unlock()

	list := ears[channel]
	for i, item := range list {
		if item.id != id {
			continue
		}

		ears[channel] = append(list[:i], list[i+1:]...)
		if len(ears[channel]) == 0 {
			delete(ears, channel)
		}

		return
	}
}
