package wire

import (
	"errors"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/ui"
)

// errBusy reports a full command queue. Commands are never dropped
// silently; the caller shows the error and the user retries.
var errBusy = errors.New("download-manager: command queue is full")

// Add queues one job for the worker.
func (b *Backend) Add(req ui.AddRequest) error {
	return b.enqueue(command{kind: cmdAdd, add: req})
}

// Control queues one pause, resume, cancel, retry, or remove command.
func (b *Backend) Control(c ui.Control) error {
	return b.enqueue(command{kind: cmdControl, ctl: c})
}

// Active asks for a fresh active set. A full queue reports an error and
// the caller retries next tick.
func (b *Backend) Active() error {
	return b.enqueue(command{kind: cmdActive})
}

// Page asks for one history page.
func (b *Backend) Page(req ui.PageRequest) error {
	return b.enqueue(command{kind: cmdPage, page: req})
}

// Summary asks for the queue counts.
func (b *Backend) Summary(gen uint64) error {
	return b.enqueue(command{kind: cmdSummary, gen: gen})
}

// SetDark stores the theme for the next launch.
func (b *Backend) SetDark(dark bool) error {
	b.mu.Lock()
	b.dark = dark
	b.mu.Unlock()

	return b.enqueue(command{kind: cmdDark, dark: dark})
}

// Dark returns the theme stored at open time.
func (b *Backend) Dark() (bool, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.dark, nil
}

// enqueue hands one command to the worker without waiting.
func (b *Backend) enqueue(c command) error {
	select {
	case b.cmds <- c:
		return nil
	default:
		return errBusy
	}
}
