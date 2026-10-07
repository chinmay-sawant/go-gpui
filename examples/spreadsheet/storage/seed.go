package storage

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/workbook"
)

// SeedVersion is the fixture version SeedDummy writes.
const SeedVersion = 1

// SeedDummy writes the shipped fixtures once and reports whether it wrote
// them. A marker row keeps later opens from duplicating fixtures or
// overwriting user edits.
func (s *Store) SeedDummy(ctx context.Context) (bool, error) {
	return s.SeedWith(ctx, SeedVersion, workbook.SeedDummy()...)
}

// SeedWith writes books when the seed marker is older than version. The
// marker and the fixtures commit in one transaction.
func (s *Store) SeedWith(ctx context.Context, version int, books ...*workbook.Workbook) (bool, error) {
	ctx, cancel := bulkCtx(ctx)
	defer cancel()

	wrote := false

	err := s.do(ctx, func(ctx context.Context, db *sql.DB) error {
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			return err
		}

		var have int

		err = tx.QueryRowContext(ctx,
			`SELECT CAST(value AS INTEGER) FROM meta WHERE key = 'seed_version'`).Scan(&have)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return rollback(tx, err)
		}

		if err == nil && have >= version {
			return tx.Rollback()
		}

		for _, w := range books {
			if err := insertWorkbook(ctx, tx, w); err != nil {
				return rollback(tx, err)
			}
		}

		if _, err := tx.ExecContext(ctx,
			`INSERT INTO meta (key, value) VALUES ('seed_version', ?)
			 ON CONFLICT (key) DO UPDATE SET value = excluded.value`,
			strconv.Itoa(version)); err != nil {
			return rollback(tx, err)
		}

		if err := tx.Commit(); err != nil {
			return err
		}

		wrote = true

		return nil
	})

	return wrote, err
}
