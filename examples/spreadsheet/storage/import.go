package storage

import (
	"context"
	"database/sql"

	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/workbook"
)

// ImportOptions controls an import transaction.
type ImportOptions struct {
	Replace  bool                        // clear the sheet's cells first
	Origin   string                      // revision label; "import" when empty
	Progress func(done, total int) error // optional; an error aborts
}

// ImportCells writes a whole import in one transaction. A cancelled or
// failed import leaves the existing workbook intact, because nothing is
// committed until every cell is written.
func (s *Store) ImportCells(ctx context.Context, wb workbook.ID, sheet workbook.SheetID, edits []workbook.CellEdit, baseRev int64, opt ImportOptions) (int64, error) {
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

		if opt.Replace {
			if _, err := tx.ExecContext(ctx, `DELETE FROM cells WHERE sheet_id = ?`, int64(sheet)); err != nil {
				return rollback(tx, err)
			}
		}

		if err := writeCells(ctx, tx, sheet, edits, opt.Progress); err != nil {
			return rollback(tx, err)
		}

		newRev = rev + 1

		origin := opt.Origin
		if origin == "" {
			origin = "import"
		}

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
