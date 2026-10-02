// Package clipboard reads and writes the desktop clipboard.
// The last text is also kept in memory when the OS clipboard is
// missing, or when tests turn memory-only mode on.
package clipboard

import "sync"

var (
	mu     sync.Mutex
	mem    string
	memory bool
)

// UseMemory forces the in-memory clipboard. Tests call this
// before any Write so they never touch the desktop clipboard.
func UseMemory(on bool) {
	mu.Lock()
	memory = on
	mu.Unlock()
}

// Write stores text and tries the OS clipboard.
func Write(text string) {
	mu.Lock()
	mem = text
	only := memory
	mu.Unlock()

	if !only {
		writeOS(text)
	}
}

// Read returns the OS clipboard, or the last Write when that fails.
func Read() string {
	mu.Lock()
	only := memory
	mu.Unlock()

	if !only {
		if text, ok := readOS(); ok {
			return text
		}
	}

	mu.Lock()
	defer mu.Unlock()

	return mem
}
