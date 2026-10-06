package store

import (
	"context"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
)

// Page size rules for history.
const (
	DefaultPageSize = 50
	MaxPageSize     = 200
)

// Cursor points at the last job of a page: its UpdatedAt plus stable ID.
type Cursor struct {
	UpdatedAt time.Time
	ID        string
}

// Page is one chunk of history, newest first.
type Page struct {
	Jobs    []domain.Job
	Next    Cursor
	HasMore bool
}

// History pages terminal history by (UpdatedAt, ID) descending. The zero
// cursor starts at the newest row.
func (s *Store) History(ctx context.Context, after Cursor, limit int) (Page, error) {
	if limit <= 0 {
		limit = DefaultPageSize
	}

	if limit > MaxPageSize {
		limit = MaxPageSize
	}

	query, args := historyQuery(after, limit+1)

	jobs, err := s.queryJobs(ctx, query, args...)
	if err != nil {
		return Page{}, err
	}

	page := Page{Jobs: jobs}
	if len(jobs) > limit {
		page.HasMore = true
		page.Jobs = jobs[:limit]
	}

	if len(page.Jobs) > 0 {
		last := page.Jobs[len(page.Jobs)-1]
		page.Next = Cursor{UpdatedAt: last.UpdatedAt, ID: last.ID}
	}

	return page, nil
}
