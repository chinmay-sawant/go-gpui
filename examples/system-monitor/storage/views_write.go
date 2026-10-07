package storage

import (
	"context"
	"errors"
	"time"
)

// PutView inserts or replaces one saved view by ID.
func (s *Store) PutView(ctx context.Context, v View) error {
	if err := s.ready(); err != nil {
		return err
	}
	if v.ID == "" || v.Name == "" {
		return errors.New("storage: view needs an id and a name")
	}

	updated := v.Updated
	if updated.IsZero() {
		updated = time.Now()
	}

	ctx, cancel := s.opCtx(ctx)
	defer cancel()

	_, err := s.db.ExecContext(ctx, `
INSERT INTO saved_views(id, name, query, sort, descending, updated_ns) VALUES(?, ?, ?, ?, ?, ?)
ON CONFLICT(id) DO UPDATE SET
	name = excluded.name, query = excluded.query, sort = excluded.sort,
	descending = excluded.descending, updated_ns = excluded.updated_ns`,
		v.ID, v.Name, v.Query, v.Sort, boolInt(v.Desc), updated.UnixNano())

	return err
}

// DeleteView removes one saved view.
func (s *Store) DeleteView(ctx context.Context, id string) error {
	if err := s.ready(); err != nil {
		return err
	}

	ctx, cancel := s.opCtx(ctx)
	defer cancel()

	_, err := s.db.ExecContext(ctx, `DELETE FROM saved_views WHERE id = ?`, id)

	return err
}
