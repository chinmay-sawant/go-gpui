package store

import (
	"context"
	"database/sql"
	"strconv"
	"time"
)

func ensureDummyTx(ctx context.Context, db *sql.DB, o DummyOptions) (DummySetup, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return DummySetup{}, err
	}

	defer func() { _ = tx.Rollback() }()

	now := time.Now().UnixNano()

	if _, err := tx.ExecContext(ctx,
		`INSERT INTO sessions(name, kind, created_ns) VALUES(?, 'dummy', ?)
		 ON CONFLICT(name) DO NOTHING`, dummySession, now); err != nil {
		return DummySetup{}, err
	}

	sess, err := sessionByNameTx(ctx, tx, dummySession)
	if err != nil {
		return DummySetup{}, err
	}

	for _, k := range dummySourceKeys {
		if _, err := tx.ExecContext(ctx,
			`INSERT INTO sources
			 (session_id, kind, path, label, state, seed, rate, created_ns, updated_ns)
			 VALUES(?, 'dummy', ?, ?, 'idle', ?, ?, ?, ?)
			 ON CONFLICT(session_id, path) DO NOTHING`,
			int64(sess.ID), k.key, k.label, o.Seed, o.Rate, now, now); err != nil {
			return DummySetup{}, err
		}
	}

	ver, err := metaIntTx(ctx, tx, "dummy_seed_version")
	if err != nil {
		return DummySetup{}, err
	}

	if ver >= dummySeedVer {
		return dummyFinish(ctx, tx, sess, false, ver)
	}

	quota, rem := o.Count/len(dummySourceKeys), o.Count%len(dummySourceKeys)

	for i, k := range dummySourceKeys {
		src, err := sourceByPathTx(ctx, tx, sess.ID, k.key)
		if err != nil {
			return DummySetup{}, err
		}

		n := quota
		if i < rem {
			n++
		}

		if err := seedSource(ctx, tx, src, n, o.Seed, o.Policy); err != nil {
			return DummySetup{}, err
		}
	}

	if _, err := tx.ExecContext(ctx,
		`INSERT INTO meta(key, value) VALUES('dummy_seed_version', ?)
		 ON CONFLICT(key) DO UPDATE SET value = excluded.value`,
		strconv.Itoa(dummySeedVer)); err != nil {
		return DummySetup{}, err
	}

	return dummyFinish(ctx, tx, sess, true, dummySeedVer)
}
