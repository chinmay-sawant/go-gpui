package store

import (
	"context"
	"time"
)

// retryDelay spaces out reads and commits that failed, so a locked database
// or a transient IO error does not spin.
const retryDelay = 500 * time.Millisecond

// Run follows the source until the context is cancelled or the source
// finishes. A read or commit error is counted and retried; no entries are
// dropped on the retry path because the store dedupes a replayed batch.
func (in *Ingestor) Run(ctx context.Context) error {
	for {
		if err := ctx.Err(); err != nil {
			in.finish()

			return err
		}

		if in.paused.Load() && in.r.Replayable() {
			if err := in.nap(ctx, in.poll()); err != nil {
				in.finish()

				return err
			}

			continue
		}

		b, err := in.r.Read(ctx)
		if err != nil {
			if ctx.Err() != nil {
				in.finish()

				return ctx.Err()
			}

			in.failed(err)

			if err := in.nap(ctx, retryDelay); err != nil {
				in.finish()

				return err
			}

			continue
		}

		if err := in.commit(ctx, b); err != nil {
			if ctx.Err() != nil {
				in.finish()

				return ctx.Err()
			}

			in.failed(err)

			if err := in.nap(ctx, retryDelay); err != nil {
				in.finish()

				return err
			}

			continue
		}

		if b.Done {
			return nil
		}
	}
}
