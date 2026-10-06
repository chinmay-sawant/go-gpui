package store

import (
	"context"
	"database/sql"
)

func overLimit(ctx context.Context, db *sql.DB, r Retention) (bool, error) {
	var count, bytes int64

	if err := db.QueryRowContext(ctx,
		`SELECT count(*), COALESCE(SUM(bytes), 0) FROM entries`).
		Scan(&count, &bytes); err != nil {
		return false, err
	}

	if r.MaxRows > 0 && count > r.MaxRows {
		return true, nil
	}

	return r.MaxBytes > 0 && bytes > r.MaxBytes, nil
}

func deleteWhere(ctx context.Context, db *sql.DB, where string, arg any) (int64, error) {
	res, err := db.ExecContext(ctx,
		`DELETE FROM entries WHERE id IN
		 (SELECT id FROM entries WHERE `+where+` ORDER BY id LIMIT ?)`,
		arg, pruneBatch)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func deleteOldest(ctx context.Context, db *sql.DB) (int64, error) {
	res, err := db.ExecContext(ctx,
		`DELETE FROM entries WHERE id IN
		 (SELECT id FROM entries ORDER BY id LIMIT ?)`, pruneBatch)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
