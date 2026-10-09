package bridge

// SetCommandNotifier wakes Android as soon as a command is accepted.
// Notification runs outside the queue lock and may drain the queue.
func SetCommandNotifier(fn func()) {
	mu.Lock()
	commandReady = fn
	pending := len(q) > 0
	mu.Unlock()
	if pending && fn != nil {
		fn()
	}
}
