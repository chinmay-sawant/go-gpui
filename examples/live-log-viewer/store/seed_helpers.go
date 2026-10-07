package store

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

func sessionByNameTx(ctx context.Context, tx *sql.Tx, name string) (entry.Session, error) {
	var (
		sess entry.Session
		ns   int64
	)

	err := tx.QueryRowContext(ctx,
		`SELECT id, name, kind, created_ns FROM sessions WHERE name = ?`, name).
		Scan(&sess.ID, &sess.Name, &sess.Kind, &ns)
	if err != nil {
		return entry.Session{}, err
	}

	sess.Created = time.Unix(0, ns)

	return sess, nil
}

func sourceByPathTx(ctx context.Context, tx *sql.Tx, session entry.SessionID, path string) (entry.Source, error) {
	row := tx.QueryRowContext(ctx, sourceColumns+
		` WHERE session_id = ? AND path = ?`, int64(session), path)

	return scanSource(row)
}

func metaIntTx(ctx context.Context, tx *sql.Tx, key string) (int, error) {
	var raw string

	err := tx.QueryRowContext(ctx, `SELECT value FROM meta WHERE key = ?`, key).Scan(&raw)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, nil
	}

	if err != nil {
		return 0, err
	}

	return strconv.Atoi(raw)
}
