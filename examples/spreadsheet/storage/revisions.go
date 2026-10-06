package storage

import (
	"context"
	"database/sql"

	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/workbook"
)

// TrimRevisions deletes old revision rows in bounded batches and reports
// how many it removed. keep is the number of newest rows to stay; a call
// does at most 100 batches, so cleanup never runs a long job during an
// interaction.
func (s *Store) TrimRevisions(ctx context.Context, wb workbook.ID, keep, batch int) (int, error) {
	if keep < 0 {
		keep = 0
	}

	if batch <= 0 {
		batch = 100
	}

	deleted := 0

	err := s.do(ctx, func(ctx context.Context, db *sql.DB) error {
		for i := 0; i < 100; i++ {
			var count int

			if err := db.QueryRowContext(ctx,
				`SELECT COUNT(*) FROM revisions WHERE workbook_id = ?`, int64(wb)).Scan(&count); err != nil {
				return err
			}

			if count <= keep {
				return nil
			}

			n := batch
			if n > count-keep {
				n = count - keep
			}

			res, err := db.ExecContext(ctx,
				`DELETE FROM revisions WHERE id IN (
					SELECT id FROM revisions WHERE workbook_id = ? ORDER BY id ASC LIMIT ?)`,
				int64(wb), n)
			if err != nil {
				return err
			}

			removed, err := res.RowsAffected()
			if err != nil {
				return err
			}

			deleted += int(removed)

			if int(removed) < batch {
				return nil
			}
		}

		return nil
	})

	return deleted, err
}
