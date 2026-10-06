package storage

import (
	"context"
	"time"
)

// Recording session states.
const (
	StateRecording = "recording"
	StateDone      = "done"
	StateFailed    = "failed"
	StateTemp      = "temporary"
)

// Session is one recording. Source labels the collector ("dummy", "live"), so
// fixture rows never pass as live history. Rows counts the raw history rows
// when a read fills it.
type Session struct {
	ID       int64
	Name     string
	Source   string
	Mode     string
	State    string
	Note     string
	Started  time.Time
	Ended    time.Time
	Sampling time.Duration
	Rows     int64
}

// StartSession inserts a recording session and returns its ID.
func (s *Store) StartSession(ctx context.Context, sess Session) (int64, error) {
	return 0, errNotImplemented
}

// EndSession marks a session done or failed and stores a closing note. It
// leaves Started in place and never deletes rows.
func (s *Store) EndSession(ctx context.Context, id int64, state, note string) error {
	return errNotImplemented
}

// Sessions returns the most recent sessions first. limit <= 0 returns all.
func (s *Store) Sessions(ctx context.Context, limit int) ([]Session, error) {
	return nil, errNotImplemented
}

// Session returns one session by ID.
func (s *Store) Session(ctx context.Context, id int64) (Session, bool, error) {
	return Session{}, false, errNotImplemented
}
