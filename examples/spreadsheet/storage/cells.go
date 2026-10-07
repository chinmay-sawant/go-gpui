package storage

import (
	"context"
	"database/sql"

	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/workbook"
)

// SaveCells writes a batch of edits for one sheet in one transaction and
// advances the workbook revision. baseRev is the caller's SavedRev; a
// mismatch returns ErrConflict and changes nothing. The new revision comes
// back only after COMMIT.
func (s *Store) SaveCells(ctx context.Context, wb workbook.ID, sheet workbook.SheetID, edits []workbook.CellEdit, baseRev int64, origin string) (int64, error) {
	if len(edits) == 0 {
		return baseRev, nil
	}

	ctx, cancel := bulkCtx(ctx)
	defer cancel()

	var newRev int64

	err := s.do(ctx, func(ctx context.Context, db *sql.DB) error {
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			return err
		}

		rev, err := checkRev(ctx, tx, wb, sheet, baseRev)
		if err != nil {
			return rollback(tx, err)
		}

		if err := writeCells(ctx, tx, sheet, edits, nil); err != nil {
			return rollback(tx, err)
		}

		newRev = rev + 1

		if err := touchWorkbook(ctx, tx, wb, newRev, origin, len(edits)); err != nil {
			return rollback(tx, err)
		}

		return tx.Commit()
	})
	if err != nil {
		return 0, err
	}

	return newRev, nil
}
