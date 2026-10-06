package wire

import (
	"context"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
	"github.com/chinmay-sawant/ownframe/examples/download-manager/store"
	"github.com/chinmay-sawant/ownframe/examples/download-manager/ui"
)

// scan walks keyset pages until limit matching rows or the end of the
// history. It returns the cursor for the next page, or "" at the end.
func (b *Backend) scan(ctx context.Context, after store.Cursor, limit int, f ui.Filter) ([]domain.Job, string, error) {
	var out []domain.Job

	cur := after
	for {
		page, err := b.store.History(ctx, cur, store.MaxPageSize)
		if err != nil {
			return nil, "", err
		}

		for i, job := range page.Jobs {
			if !matches(f, job.State) {
				continue
			}

			out = append(out, job)
			if len(out) < limit {
				continue
			}

			if i == len(page.Jobs)-1 && !page.HasMore {
				return out, "", nil
			}

			last := store.Cursor{UpdatedAt: job.UpdatedAt, ID: job.ID}

			return out, encodeCursor(last), nil
		}

		if !page.HasMore {
			return out, "", nil
		}

		cur = page.Next
	}
}

// matches reports whether one state belongs to a filter.
func matches(f ui.Filter, s domain.State) bool {
	switch f {
	case ui.FilterCompleted:
		return s == domain.StateCompleted
	case ui.FilterFailed:
		return s == domain.StateFailed
	case ui.FilterCancelled:
		return s == domain.StateCancelled
	default:
		return true
	}
}
