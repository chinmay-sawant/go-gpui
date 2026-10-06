package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

func sessionByName(ctx context.Context, db *sql.DB, name string) (entry.Session, error) {
	var (
		sess entry.Session
		ns   int64
	)

	err := db.QueryRowContext(ctx,
		`SELECT id, name, kind, created_ns FROM sessions WHERE name = ?`, name).
		Scan(&sess.ID, &sess.Name, &sess.Kind, &ns)
	if err != nil {
		return entry.Session{}, err
	}

	sess.Created = time.Unix(0, ns)

	return sess, nil
}
