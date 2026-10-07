package storage

import (
	"context"
	"fmt"
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
// fixture rows never pass as live history. Rows counts the raw history rows.
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
	if err := s.ready(); err != nil {
		return 0, err
	}
	if sess.Name == "" {
		sess.Name = "recording"
	}
	if sess.State == "" {
		sess.State = StateRecording
	}

	started := sess.Started
	if started.IsZero() {
		started = time.Now()
	}

	ctx, cancel := s.opCtx(ctx)
	defer cancel()

	res, err := s.db.ExecContext(ctx, `
INSERT INTO sessions(name, source, mode, state, note, started_ns, sampling_ms)
VALUES(?, ?, ?, ?, ?, ?, ?)`,
		sess.Name, sess.Source, sess.Mode, sess.State, sess.Note,
		started.UnixNano(), sess.Sampling.Milliseconds())
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// EndSession marks a session done or failed and stores a closing note. It
// never deletes rows.
func (s *Store) EndSession(ctx context.Context, id int64, state, note string) error {
	if err := s.ready(); err != nil {
		return err
	}
	if state == "" {
		state = StateDone
	}

	ctx, cancel := s.opCtx(ctx)
	defer cancel()

	res, err := s.db.ExecContext(ctx, `
UPDATE sessions SET state = ?, note = ?, ended_ns = ? WHERE id = ?`,
		state, note, time.Now().UnixNano(), id)
	if err != nil {
		return err
	}

	if n, _ := res.RowsAffected(); n == 0 {
		return fmt.Errorf("storage: session %d not found", id)
	}

	return nil
}
