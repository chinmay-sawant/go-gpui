package store

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
)

// schemaVersion is the newest schema this build writes.
const schemaVersion = 1

type migration struct {
	version int
	sql     string
}

var migrations = []migration{
	{version: 1, sql: schemaCore + schemaData},
}

// migrate applies every pending migration in one transaction. A failure
// leaves the old version with no partial schema. A newer database is
// rejected before anything is written.
func migrate(ctx context.Context, db *sql.DB) error {
	return migrateList(ctx, db, migrations)
}

func migrateList(ctx context.Context, db *sql.DB, list []migration) error {
	cur, err := schemaOf(ctx, db)
	if err != nil {
		return err
	}

	if cur > len(list) {
		return fmt.Errorf("%w: found %d, supported %d", ErrSchemaNewer, cur, len(list))
	}

	if cur == len(list) {
		return nil
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() { _ = tx.Rollback() }()

	for _, m := range list[cur:] {
		if _, err := tx.ExecContext(ctx, m.sql); err != nil {
			return err
		}
	}

	for _, kv := range [][2]string{
		{"app_id", appID},
		{"schema_version", strconv.Itoa(len(list))},
	} {
		if _, err := tx.ExecContext(ctx,
			`INSERT INTO meta(key, value) VALUES(?, ?)
			 ON CONFLICT(key) DO UPDATE SET value = excluded.value`,
			kv[0], kv[1]); err != nil {
			return err
		}
	}

	return tx.Commit()
}
