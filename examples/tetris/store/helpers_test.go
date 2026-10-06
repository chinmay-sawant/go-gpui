package store

import (
	"context"
	"database/sql"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
)

// testResult builds a valid score record.
func testResult(id string, score int) game.Result {
	return game.Result{
		ID:             id,
		Score:          score,
		Lines:          min(score/1000, game.MaxLines),
		Level:          1,
		Pieces:         1,
		DurationMS:     1000,
		Seed:           5,
		Ruleset:        game.Ruleset,
		FixtureVersion: game.FixtureVersion,
	}
}

// scalar reads one text value through the worker.
func (s *Store) scalar(t *testing.T, query string, args ...any) string {
	t.Helper()

	var out string

	err := s.do(context.Background(), func(ctx context.Context, db *sql.DB) error {
		return db.QueryRowContext(ctx, query, args...).Scan(&out)
	})
	if err != nil {
		t.Fatalf("query %q: %v", query, err)
	}

	return out
}

// exec runs one statement through the worker.
func (s *Store) exec(t *testing.T, query string, args ...any) {
	t.Helper()

	err := s.do(context.Background(), func(ctx context.Context, db *sql.DB) error {
		_, err := db.ExecContext(ctx, query, args...)

		return err
	})
	if err != nil {
		t.Fatalf("exec %q: %v", query, err)
	}
}
