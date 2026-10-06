package store

import (
	"context"
	"database/sql"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

// dummyFinish snapshots the sources and commits the fixture transaction.
func dummyFinish(ctx context.Context, tx *sql.Tx, sess entry.Session, seeded bool, ver int) (DummySetup, error) {
	setup := DummySetup{Session: sess, Seeded: seeded, Version: ver}

	for _, k := range dummySourceKeys {
		src, err := sourceByPathTx(ctx, tx, sess.ID, k.key)
		if err != nil {
			return DummySetup{}, err
		}

		setup.Sources = append(setup.Sources, src)
	}

	return setup, tx.Commit()
}
