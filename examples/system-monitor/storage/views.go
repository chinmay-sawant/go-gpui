package storage

import (
	"context"
	"time"
)

// View is a saved process table filter. Query is the search text, Sort names
// the column, and Desc marks a descending sort.
type View struct {
	ID      string
	Name    string
	Query   string
	Sort    string
	Desc    bool
	Updated time.Time
}

// Views returns every saved view, most recently updated first.
func (s *Store) Views(ctx context.Context) ([]View, error) { return nil, errNotImplemented }

// PutView inserts or replaces one saved view by ID.
func (s *Store) PutView(ctx context.Context, v View) error { return errNotImplemented }

// DeleteView removes one saved view.
func (s *Store) DeleteView(ctx context.Context, id string) error { return errNotImplemented }
