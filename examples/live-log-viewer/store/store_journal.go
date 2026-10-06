package store

import (
	"context"
	"database/sql"
	"errors"
	"io/fs"
	"strings"
)

func applyJournal(ctx context.Context, db *sql.DB) (string, error) {
	var mode string

	if err := db.QueryRowContext(ctx, "PRAGMA journal_mode=WAL").Scan(&mode); err != nil {
		if isReadOnly(err) {
			return "", wrapOpen(err)
		}

		return "", err
	}

	if strings.EqualFold(mode, "wal") {
		return "wal", nil
	}

	// A filesystem that cannot host WAL keeps a rollback journal instead.
	var fallback string

	if err := db.QueryRowContext(ctx, "PRAGMA journal_mode=DELETE").Scan(&fallback); err == nil {
		return strings.ToLower(fallback), nil
	}

	return strings.ToLower(mode), nil
}

func isReadOnly(err error) bool {
	msg := strings.ToLower(err.Error())

	return errors.Is(err, fs.ErrPermission) ||
		strings.Contains(msg, "readonly") ||
		strings.Contains(msg, "read-only")
}

func wrapOpen(err error) error {
	if isReadOnly(err) {
		return ErrReadOnly
	}

	return err
}
