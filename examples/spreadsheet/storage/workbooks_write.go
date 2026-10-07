package storage

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/workbook"
)

// CreateWorkbook inserts a new workbook with one sheet named Sheet1.
func (s *Store) CreateWorkbook(ctx context.Context, name string) (*workbook.Workbook, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		name = "Untitled"
	}

	w := workbook.New(0, name)
	sheet := w.AddSheet("Sheet1")
	now := time.Now().UTC().Format(time.RFC3339Nano)

	err := s.do(ctx, func(ctx context.Context, db *sql.DB) error {
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			return err
		}

		res, err := tx.ExecContext(ctx,
			`INSERT INTO workbooks (name, rev, created_at, updated_at) VALUES (?, 0, ?, ?)`,
			name, now, now)
		if err != nil {
			return rollback(tx, err)
		}

		id, err := res.LastInsertId()
		if err != nil {
			return rollback(tx, err)
		}

		w.SetID(workbook.ID(id))

		res, err = tx.ExecContext(ctx,
			`INSERT INTO sheets (workbook_id, name, position) VALUES (?, ?, 0)`,
			id, sheet.Name())
		if err != nil {
			return rollback(tx, err)
		}

		sid, err := res.LastInsertId()
		if err != nil {
			return rollback(tx, err)
		}

		sheet.SetID(workbook.SheetID(sid))

		return tx.Commit()
	})
	if err != nil {
		return nil, err
	}

	return w, nil
}

// DeleteWorkbook removes a workbook. Foreign keys cascade to sheets, cells,
// and revisions.
func (s *Store) DeleteWorkbook(ctx context.Context, id workbook.ID) error {
	return s.do(ctx, func(ctx context.Context, db *sql.DB) error {
		res, err := db.ExecContext(ctx, `DELETE FROM workbooks WHERE id = ?`, int64(id))
		if err != nil {
			return err
		}

		n, err := res.RowsAffected()
		if err != nil {
			return err
		}

		if n == 0 {
			return fmt.Errorf("%w: workbook %d", ErrNotFound, id)
		}

		return nil
	})
}
