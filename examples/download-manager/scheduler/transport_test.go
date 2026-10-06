package scheduler

import (
	"context"
	"sync"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/transfer"
)

// gateTransport blocks every download until open runs, so tests can pause,
// cancel, or shut down mid-flight.
type gateTransport struct {
	started chan string
	release chan struct{}
	once    sync.Once

	mu  sync.Mutex
	err error
}

// newGate builds a gated transport.
func newGate() *gateTransport {
	return &gateTransport{
		started: make(chan string, 64),
		release: make(chan struct{}),
	}
}

// Download reports a first step, waits for open or the context, then
// finishes or fails.
func (g *gateTransport) Download(ctx context.Context, req transfer.Request, report transfer.Reporter) (transfer.Outcome, error) {
	select {
	case g.started <- req.JobID:
	default:
	}

	if report != nil {
		report(transfer.Progress{Done: 5, Total: 10})
	}

	select {
	case <-ctx.Done():
		return transfer.Outcome{}, ctx.Err()
	case <-g.release:
	}

	g.mu.Lock()
	err := g.err
	g.mu.Unlock()

	if err != nil {
		return transfer.Outcome{}, err
	}

	if report != nil {
		report(transfer.Progress{Done: 10, Total: 10})
	}

	return transfer.Outcome{Path: req.Dest, Bytes: 10, Total: 10}, nil
}

// open lets every waiting download finish.
func (g *gateTransport) open() { g.once.Do(func() { close(g.release) }) }

// setError makes later downloads fail.
func (g *gateTransport) setError(err error) {
	g.mu.Lock()
	g.err = err
	g.mu.Unlock()
}
