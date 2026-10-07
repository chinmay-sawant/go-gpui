package storage

import (
	"context"
	"time"
)

// bulkTimeout bounds migrations, seeds, full loads, imports, and backups
// when the caller supplies no deadline. Interactive calls keep the shorter
// default in Store.timeout.
const bulkTimeout = 5 * time.Minute

// bulkCtx adds the bulk deadline only when the caller set none, so a
// caller's cancellation or deadline still wins.
func bulkCtx(ctx context.Context) (context.Context, context.CancelFunc) {
	if ctx == nil {
		ctx = context.Background()
	}

	if _, ok := ctx.Deadline(); ok {
		return ctx, func() {}
	}

	return context.WithTimeout(ctx, bulkTimeout)
}
