package reader

import "context"

// Read returns the next bounded batch. It waits up to one poll interval
// when no records are ready, and returns ctx.Err when cancelled.
func (r *File) Read(ctx context.Context) (Batch, error) {
	if err := ctx.Err(); err != nil {
		return Batch{}, err
	}

	if r.closed {
		return Batch{}, ErrClosed
	}

	b, err := r.readAvailable()
	if err != nil {
		return Batch{}, err
	}

	if len(b.Records) > 0 || b.Rotated || b.Done {
		b.State = r.state

		return b, nil
	}

	if err := r.wait(ctx); err != nil {
		return Batch{}, err
	}

	b.State = r.state

	return b, nil
}
