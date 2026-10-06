package wire

import (
	"context"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
	"github.com/chinmay-sawant/ownframe/examples/download-manager/store"
	"github.com/chinmay-sawant/ownframe/examples/download-manager/ui"
)

// runPage answers one history request.
func (b *Backend) runPage(req ui.PageRequest) {
	resp, err := b.history(b.ctx, req)
	if err != nil {
		b.notice("History: " + err.Error())

		return
	}

	b.push(ui.Update{Kind: ui.UpdateHistory, Page: &resp})
}

// history builds one page. The store pages terminal rows by (UpdatedAt, ID)
// descending; the filtered view scans those pages and keeps the matches,
// so the returned cursor stays a real keyset key.
func (b *Backend) history(ctx context.Context, req ui.PageRequest) (ui.PageResponse, error) {
	limit := req.Limit
	if limit <= 0 {
		limit = ui.HistoryPage
	}

	resp := ui.PageResponse{Gen: req.Gen, Filter: req.Filter}

	counts, err := b.store.Summary(ctx)
	if err != nil {
		return resp, err
	}

	resp.Total = totalFor(counts, req.Filter)
	after := decodeCursor(req.Before)

	if req.Filter == ui.FilterAll {
		page, err := b.store.History(ctx, after, limit)
		if err != nil {
			return resp, err
		}

		resp.Rows = b.rows(page.Jobs)
		if page.HasMore {
			resp.Next = encodeCursor(page.Next)
		}

		return resp, nil
	}

	jobs, next, err := b.scan(ctx, after, limit, req.Filter)
	if err != nil {
		return resp, err
	}

	resp.Rows = b.rows(jobs)
	resp.Next = next

	return resp, nil
}

// totalFor returns the filtered count from the summary.
func totalFor(counts store.Counts, f ui.Filter) int {
	switch f {
	case ui.FilterCompleted:
		return counts.ByState[domain.StateCompleted]
	case ui.FilterFailed:
		return counts.ByState[domain.StateFailed]
	case ui.FilterCancelled:
		return counts.ByState[domain.StateCancelled]
	default:
		return counts.ByState[domain.StateCompleted] +
			counts.ByState[domain.StateFailed] +
			counts.ByState[domain.StateCancelled]
	}
}
