package store

import (
	"database/sql"

	"github.com/chinmay-sawant/go-gpui/examples/teams/activity"
	"github.com/chinmay-sawant/go-gpui/examples/teams/calendar"
	"github.com/chinmay-sawant/go-gpui/examples/teams/calls"
	"github.com/chinmay-sawant/go-gpui/examples/teams/channels"
	"github.com/chinmay-sawant/go-gpui/examples/teams/chat"
	"github.com/chinmay-sawant/go-gpui/examples/teams/files"
)

// Data is the whole app state as it lives in the database.
type Data struct {
	Activity activity.Data
	Chat     chat.Data
	Channels channels.Data
	Calendar calendar.Data
	Calls    calls.Data
	Files    files.Data
	Presence string
	Dark     bool
}

// Empty reports whether the database has never been seeded.
func (s *Store) Empty() (bool, error) {
	var n int

	err := s.db.QueryRow(`SELECT COUNT(*) FROM meta WHERE key = 'seeded'`).Scan(&n)

	return n == 0, err
}

// fail rolls the transaction back and returns err.
func fail(tx *sql.Tx, err error) error {
	_ = tx.Rollback()

	return err
}
