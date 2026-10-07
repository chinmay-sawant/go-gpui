package storage

import (
	"context"
	"database/sql"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/workbook"
)

// touchWorkbook advances the revision and records the change set.
func touchWorkbook(ctx context.Context, tx *sql.Tx, wb workbook.ID, rev int64, origin string, changed int) error {
	if origin == "" {
		origin = "edit"
	}

	now := time.Now().UTC().Format(time.RFC3339Nano)

	if _, err := tx.ExecContext(ctx,
		`UPDATE workbooks SET rev = ?, updated_at = ? WHERE id = ?`,
		rev, now, int64(wb)); err != nil {
		return err
	}

	_, err := tx.ExecContext(ctx,
		`INSERT INTO revisions (workbook_id, rev, origin, changed, created_at) VALUES (?, ?, ?, ?, ?)`,
		int64(wb), rev, origin, changed, now)

	return err
}
