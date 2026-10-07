package storage

import (
	"context"
	"errors"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// Seed writes a labeled fixture recording once per fixture version. It
// returns true when it wrote and false when the stored version already covers
// f. A later version replaces the fixture session in one transaction.
func (s *Store) Seed(ctx context.Context, f domain.Fixture) (bool, error) {
	if err := s.ready(); err != nil {
		return false, err
	}
	if f.Version <= 0 {
		return false, errors.New("storage: fixture version must be positive")
	}

	ctx, cancel := s.opCtx(ctx)
	defer cancel()

	current, err := s.FixtureVersion(ctx)
	if err != nil {
		return false, err
	}
	if current >= f.Version {
		return false, nil
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return false, err
	}
	defer func() { _ = tx.Rollback() }()

	if _, err := tx.ExecContext(ctx, `DELETE FROM sessions WHERE kind = 'fixture'`); err != nil {
		return false, err
	}

	started, ended := f.Started, f.Started
	if len(f.Samples) > 0 {
		started = f.Samples[0].Stamp.At
		ended = f.Samples[len(f.Samples)-1].Stamp.At
	}

	res, err := tx.ExecContext(ctx, `
INSERT INTO sessions(name, source, mode, state, note, kind, started_ns, ended_ns, sampling_ms)
VALUES(?, ?, 'dummy', ?, ?, 'fixture', ?, ?, ?)`,
		nameOr(f.Name, "dummy history"), f.Source, StateDone, f.Note,
		started.UnixNano(), ended.UnixNano(), f.Step.Milliseconds())
	if err != nil {
		return false, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return false, err
	}

	var rows []Row
	for _, sample := range f.Samples {
		rows = append(rows, sample.Records()...)
	}

	if _, err := appendTx(ctx, tx, id, rows); err != nil {
		return false, err
	}

	if err := markFixtureTx(ctx, tx, f.Version); err != nil {
		return false, err
	}

	if err := tx.Commit(); err != nil {
		return false, err
	}

	return true, nil
}

// nameOr lives in seed_version.go.
