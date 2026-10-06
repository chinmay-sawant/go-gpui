package store

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
)

// rulesetName and fixtureVersion stamp every stored row, so a score or a
// replay can be checked against the model that produced it.
const (
	rulesetName    = game.Ruleset
	fixtureVersion = game.FixtureVersion
)

// scoreColumns is the select list every score query shares.
const scoreColumns = `id, score, lines, level, pieces, duration_ms, ` +
	`seed, ruleset, fixture_version, dummy`

// validateResult rejects a score record the schema cannot store.
func validateResult(r game.Result) error {
	if r.ID == "" {
		return fmt.Errorf("%w: missing id", ErrInvalid)
	}

	if r.Score < 0 || r.Score > game.MaxScore {
		return fmt.Errorf("%w: score %d", ErrInvalid, r.Score)
	}

	if r.Lines < 0 || r.Lines > game.MaxLines {
		return fmt.Errorf("%w: lines %d", ErrInvalid, r.Lines)
	}

	if r.Level < 1 || r.Level > game.MaxLevel {
		return fmt.Errorf("%w: level %d", ErrInvalid, r.Level)
	}

	if r.Pieces < 0 || r.Pieces > game.MaxPieces {
		return fmt.Errorf("%w: pieces %d", ErrInvalid, r.Pieces)
	}

	if r.DurationMS < 0 || time.Duration(r.DurationMS)*time.Millisecond > game.MaxDuration {
		return fmt.Errorf("%w: duration %d", ErrInvalid, r.DurationMS)
	}

	if r.Ruleset == "" {
		return fmt.Errorf("%w: missing ruleset", ErrInvalid)
	}

	return nil
}

// scanScore reads one score row in scoreColumns order.
func scanScore(rows *sql.Rows) (game.Result, error) {
	var r game.Result

	err := rows.Scan(
		&r.ID, &r.Score, &r.Lines, &r.Level, &r.Pieces, &r.DurationMS,
		&r.Seed, &r.Ruleset, &r.FixtureVersion, &r.Dummy,
	)

	return r, err
}
