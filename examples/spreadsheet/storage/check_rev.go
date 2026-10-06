package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/workbook"
)

// checkRev loads the workbook revision, compares it with baseRev, and
// verifies that sheet belongs to the workbook.
func checkRev(ctx context.Context, tx *sql.Tx, wb workbook.ID, sheet workbook.SheetID, baseRev int64) (int64, error) {
	var rev int64

	err := tx.QueryRowContext(ctx, `SELECT rev FROM workbooks WHERE id = ?`, int64(wb)).Scan(&rev)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, fmt.Errorf("%w: workbook %d", ErrNotFound, wb)
	}

	if err != nil {
		return 0, err
	}

	if rev != baseRev {
		return 0, fmt.Errorf("%w: workbook %d at %d, caller has %d", ErrConflict, wb, rev, baseRev)
	}

	var owner int64

	err = tx.QueryRowContext(ctx, `SELECT workbook_id FROM sheets WHERE id = ?`, int64(sheet)).Scan(&owner)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, fmt.Errorf("%w: sheet %d", ErrNotFound, sheet)
	}

	if err != nil {
		return 0, err
	}

	if owner != int64(wb) {
		return 0, fmt.Errorf("%w: sheet %d", ErrNotFound, sheet)
	}

	return rev, nil
}
