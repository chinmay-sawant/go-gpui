package scheduler

import (
	"context"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/transfer"
)

// burstTransport reports many progress steps and then completes.
type burstTransport struct{ steps int }

// Download reports steps progress events before finishing.
func (b *burstTransport) Download(ctx context.Context, req transfer.Request, report transfer.Reporter) (transfer.Outcome, error) {
	for i := 1; i <= b.steps; i++ {
		if err := ctx.Err(); err != nil {
			return transfer.Outcome{}, err
		}

		if report != nil {
			report(transfer.Progress{Done: int64(i), Total: int64(b.steps)})
		}
	}

	return transfer.Outcome{Path: req.Dest, Bytes: int64(b.steps), Total: int64(b.steps)}, nil
}

// instantTransport completes every download at once.
type instantTransport struct{}

// Download returns a fixed outcome.
func (instantTransport) Download(ctx context.Context, req transfer.Request, report transfer.Reporter) (transfer.Outcome, error) {
	return transfer.Outcome{Path: req.Dest, Bytes: 7, Total: 7}, nil
}
