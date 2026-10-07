package storage

import (
	"database/sql"
	"time"
)

// scanSession reads one session row.
func scanSession(scanner interface{ Scan(...any) error }) (Session, error) {
	var (
		sess    Session
		started int64
		ended   sql.NullInt64
		ms      int64
	)

	if err := scanner.Scan(&sess.ID, &sess.Name, &sess.Source, &sess.Mode, &sess.State,
		&sess.Note, &started, &ended, &ms, &sess.Rows); err != nil {
		return Session{}, err
	}

	sess.Started = time.Unix(0, started)
	sess.Sampling = time.Duration(ms) * time.Millisecond
	if ended.Valid {
		sess.Ended = time.Unix(0, ended.Int64)
	}

	return sess, nil
}
